package service

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/go-rat/chix"
	"github.com/go-rat/sessions"
	"github.com/knadh/koanf/v2"
	"github.com/leonelquinteros/gotext"
	"github.com/spf13/cast"

	"github.com/tnb-labs/panel/internal/biz"
	"github.com/tnb-labs/panel/internal/http/request"
	"github.com/tnb-labs/panel/pkg/rsacrypto"
)

type UserService struct {
	t        *gotext.Locale
	conf     *koanf.Koanf
	session  *sessions.Manager
	userRepo biz.UserRepo
}

func NewUserService(t *gotext.Locale, conf *koanf.Koanf, session *sessions.Manager, user biz.UserRepo) *UserService {
	gob.Register(rsa.PrivateKey{}) // 必须注册 rsa.PrivateKey 类型否则无法反序列化 session 中的 key
	return &UserService{
		t:        t,
		conf:     conf,
		session:  session,
		userRepo: user,
	}
}

func (s *UserService) GetKey(w http.ResponseWriter, r *http.Request) {
	key, err := rsacrypto.GenerateKey()
	if err != nil {
		Error(w, http.StatusInternalServerError, "%v", err)
		return
	}

	sess, err := s.session.GetSession(r)
	if err != nil {
		Error(w, http.StatusInternalServerError, "%v", err)
		return
	}
	sess.Put("key", *key)

	pk, err := rsacrypto.PublicKeyToString(&key.PublicKey)
	if err != nil {
		Error(w, http.StatusInternalServerError, "%v", err)
		return
	}

	Success(w, pk)
}

func (s *UserService) Login(w http.ResponseWriter, r *http.Request) {
	sess, err := s.session.GetSession(r)
	if err != nil {
		Error(w, http.StatusInternalServerError, "%v", err)
		return
	}

	req, err := Bind[request.UserLogin](r)
	if err != nil {
		Error(w, http.StatusUnprocessableEntity, "%v", err)
		return
	}

	key, ok := sess.Get("key").(rsa.PrivateKey)
	if !ok {
		Error(w, http.StatusForbidden, s.t.Get("invalid key, please refresh the page"))
		return
	}

	decryptedUsername, _ := rsacrypto.DecryptData(&key, req.Username)
	decryptedPassword, _ := rsacrypto.DecryptData(&key, req.Password)
	user, err := s.userRepo.CheckPassword(string(decryptedUsername), string(decryptedPassword))
	if err != nil {
		Error(w, http.StatusForbidden, "%v", err)
		return
	}

	// 安全登录下，将当前客户端与会话绑定
	// 安全登录只在未启用面板 HTTPS 时生效
	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		Error(w, http.StatusInternalServerError, "%v", err)
		return
	}
	if req.SafeLogin && !s.conf.Bool("http.tls") {
		sess.Put("safe_login", true)
		sess.Put("safe_client", fmt.Sprintf("%x", sha256.Sum256([]byte(ip))))
	} else {
		sess.Forget("safe_login")
		sess.Forget("safe_client")
	}

	sess.Put("user_id", user.ID)
	sess.Forget("key")
	Success(w, nil)
}

func (s *UserService) Logout(w http.ResponseWriter, r *http.Request) {
	sess, err := s.session.GetSession(r)
	if err == nil {
		if err = sess.Invalidate(); err != nil {
			Error(w, http.StatusInternalServerError, "%v", err)
			return
		}
	}

	Success(w, nil)
}

func (s *UserService) IsLogin(w http.ResponseWriter, r *http.Request) {
	sess, err := s.session.GetSession(r)
	if err != nil {
		Success(w, false)
		return
	}
	Success(w, sess.Has("user_id"))
}

func (s *UserService) Info(w http.ResponseWriter, r *http.Request) {
	userID := cast.ToUint(r.Context().Value("user_id"))
	if userID == 0 {
		ErrorSystem(w)
		return
	}

	user, err := s.userRepo.Get(userID)
	if err != nil {
		ErrorSystem(w)
		return
	}

	Success(w, chix.M{
		"id":       user.ID,
		"role":     []string{"admin"},
		"username": user.Username,
		"email":    user.Email,
	})
}
