package request

type PanelSetting struct {
	Name        string `json:"name" validate:"required"`
	Locale      string `json:"locale" validate:"required"`
	Entrance    string `json:"entrance" validate:"required"`
	OfflineMode bool   `json:"offline_mode"`
	AutoUpdate  bool   `json:"auto_update"`
	WebsitePath string `json:"website_path" validate:"required"`
	BackupPath  string `json:"backup_path" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"password"`
	Email       string `json:"email" validate:"required"`
	Port        uint   `json:"port" validate:"required|min:1|max:65535"`
	HTTPS       bool   `json:"https"`
	Cert        string `json:"cert" validate:"required"`
	Key         string `json:"key" validate:"required"`
}
