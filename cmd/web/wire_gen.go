// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/tnborg/panel/internal/app"
	"github.com/tnborg/panel/internal/apps/codeserver"
	"github.com/tnborg/panel/internal/apps/docker"
	"github.com/tnborg/panel/internal/apps/fail2ban"
	"github.com/tnborg/panel/internal/apps/frp"
	"github.com/tnborg/panel/internal/apps/gitea"
	"github.com/tnborg/panel/internal/apps/memcached"
	"github.com/tnborg/panel/internal/apps/minio"
	"github.com/tnborg/panel/internal/apps/mysql"
	"github.com/tnborg/panel/internal/apps/nginx"
	"github.com/tnborg/panel/internal/apps/php74"
	"github.com/tnborg/panel/internal/apps/php80"
	"github.com/tnborg/panel/internal/apps/php81"
	"github.com/tnborg/panel/internal/apps/php82"
	"github.com/tnborg/panel/internal/apps/php83"
	"github.com/tnborg/panel/internal/apps/php84"
	"github.com/tnborg/panel/internal/apps/phpmyadmin"
	"github.com/tnborg/panel/internal/apps/podman"
	"github.com/tnborg/panel/internal/apps/postgresql"
	"github.com/tnborg/panel/internal/apps/pureftpd"
	"github.com/tnborg/panel/internal/apps/redis"
	"github.com/tnborg/panel/internal/apps/rsync"
	"github.com/tnborg/panel/internal/apps/s3fs"
	"github.com/tnborg/panel/internal/apps/supervisor"
	"github.com/tnborg/panel/internal/bootstrap"
	"github.com/tnborg/panel/internal/data"
	"github.com/tnborg/panel/internal/http/middleware"
	"github.com/tnborg/panel/internal/job"
	"github.com/tnborg/panel/internal/route"
	"github.com/tnborg/panel/internal/service"
)

import (
	_ "time/tzdata"
)

// Injectors from wire.go:

// initWeb init application.
func initWeb() (*app.Web, error) {
	koanf, err := bootstrap.NewConf()
	if err != nil {
		return nil, err
	}
	locale, err := bootstrap.NewT(koanf)
	if err != nil {
		return nil, err
	}
	logger := bootstrap.NewLog(koanf)
	db, err := bootstrap.NewDB(koanf, logger)
	if err != nil {
		return nil, err
	}
	manager, err := bootstrap.NewSession(koanf, db)
	if err != nil {
		return nil, err
	}
	cacheRepo := data.NewCacheRepo(db)
	queue := bootstrap.NewQueue()
	taskRepo := data.NewTaskRepo(locale, db, logger, queue)
	appRepo := data.NewAppRepo(locale, koanf, db, cacheRepo, taskRepo)
	userTokenRepo := data.NewUserTokenRepo(locale, db)
	middlewares := middleware.NewMiddlewares(koanf, logger, manager, appRepo, userTokenRepo)
	userRepo := data.NewUserRepo(locale, db)
	userService := service.NewUserService(locale, koanf, manager, userRepo)
	userTokenService := service.NewUserTokenService(locale, userTokenRepo)
	databaseServerRepo := data.NewDatabaseServerRepo(locale, db, logger)
	databaseUserRepo := data.NewDatabaseUserRepo(locale, db, databaseServerRepo)
	databaseRepo := data.NewDatabaseRepo(locale, db, databaseServerRepo, databaseUserRepo)
	certRepo := data.NewCertRepo(locale, db, logger)
	certAccountRepo := data.NewCertAccountRepo(locale, db, userRepo, logger)
	websiteRepo := data.NewWebsiteRepo(locale, db, cacheRepo, databaseRepo, databaseServerRepo, databaseUserRepo, certRepo, certAccountRepo)
	settingRepo := data.NewSettingRepo(locale, db, koanf, taskRepo)
	cronRepo := data.NewCronRepo(locale, db)
	backupRepo := data.NewBackupRepo(locale, db, settingRepo, websiteRepo)
	dashboardService := service.NewDashboardService(locale, koanf, taskRepo, websiteRepo, appRepo, settingRepo, cronRepo, backupRepo)
	taskService := service.NewTaskService(taskRepo)
	websiteService := service.NewWebsiteService(websiteRepo, settingRepo)
	databaseService := service.NewDatabaseService(databaseRepo)
	databaseServerService := service.NewDatabaseServerService(databaseServerRepo)
	databaseUserService := service.NewDatabaseUserService(databaseUserRepo)
	backupService := service.NewBackupService(locale, backupRepo)
	certService := service.NewCertService(locale, certRepo)
	certDNSRepo := data.NewCertDNSRepo(db)
	certDNSService := service.NewCertDNSService(certDNSRepo)
	certAccountService := service.NewCertAccountService(certAccountRepo)
	appService := service.NewAppService(locale, appRepo, cacheRepo, settingRepo)
	cronService := service.NewCronService(cronRepo)
	processService := service.NewProcessService()
	safeRepo := data.NewSafeRepo()
	safeService := service.NewSafeService(safeRepo)
	firewallService := service.NewFirewallService()
	sshRepo := data.NewSSHRepo(locale, db)
	sshService := service.NewSSHService(sshRepo)
	containerRepo := data.NewContainerRepo()
	containerService := service.NewContainerService(containerRepo)
	containerComposeRepo := data.NewContainerComposeRepo()
	containerComposeService := service.NewContainerComposeService(containerComposeRepo)
	containerNetworkRepo := data.NewContainerNetworkRepo()
	containerNetworkService := service.NewContainerNetworkService(containerNetworkRepo)
	containerImageRepo := data.NewContainerImageRepo()
	containerImageService := service.NewContainerImageService(containerImageRepo)
	containerVolumeRepo := data.NewContainerVolumeRepo()
	containerVolumeService := service.NewContainerVolumeService(containerVolumeRepo)
	fileService := service.NewFileService(locale, taskRepo)
	monitorRepo := data.NewMonitorRepo(db, settingRepo)
	monitorService := service.NewMonitorService(settingRepo, monitorRepo)
	settingService := service.NewSettingService(settingRepo)
	systemctlService := service.NewSystemctlService(locale)
	toolboxSystemService := service.NewToolboxSystemService(locale)
	toolboxBenchmarkService := service.NewToolboxBenchmarkService(locale)
	codeserverApp := codeserver.NewApp()
	dockerApp := docker.NewApp()
	fail2banApp := fail2ban.NewApp(locale, websiteRepo)
	frpApp := frp.NewApp()
	giteaApp := gitea.NewApp()
	memcachedApp := memcached.NewApp(locale)
	minioApp := minio.NewApp()
	mysqlApp := mysql.NewApp(locale, settingRepo)
	nginxApp := nginx.NewApp(locale)
	php74App := php74.NewApp(locale, taskRepo)
	php80App := php80.NewApp(locale, taskRepo)
	php81App := php81.NewApp(locale, taskRepo)
	php82App := php82.NewApp(locale, taskRepo)
	php83App := php83.NewApp(locale, taskRepo)
	php84App := php84.NewApp(locale, taskRepo)
	phpmyadminApp := phpmyadmin.NewApp(locale)
	podmanApp := podman.NewApp()
	postgresqlApp := postgresql.NewApp(locale)
	pureftpdApp := pureftpd.NewApp(locale)
	redisApp := redis.NewApp(locale)
	rsyncApp := rsync.NewApp(locale)
	s3fsApp := s3fs.NewApp(locale)
	supervisorApp := supervisor.NewApp(locale)
	loader := bootstrap.NewLoader(codeserverApp, dockerApp, fail2banApp, frpApp, giteaApp, memcachedApp, minioApp, mysqlApp, nginxApp, php74App, php80App, php81App, php82App, php83App, php84App, phpmyadminApp, podmanApp, postgresqlApp, pureftpdApp, redisApp, rsyncApp, s3fsApp, supervisorApp)
	http := route.NewHttp(userService, userTokenService, dashboardService, taskService, websiteService, databaseService, databaseServerService, databaseUserService, backupService, certService, certDNSService, certAccountService, appService, cronService, processService, safeService, firewallService, sshService, containerService, containerComposeService, containerNetworkService, containerImageService, containerVolumeService, fileService, monitorService, settingService, systemctlService, toolboxSystemService, toolboxBenchmarkService, loader)
	wsService := service.NewWsService(locale, koanf, sshRepo)
	ws := route.NewWs(wsService)
	mux, err := bootstrap.NewRouter(locale, middlewares, http, ws)
	if err != nil {
		return nil, err
	}
	server, err := bootstrap.NewHttp(koanf, mux)
	if err != nil {
		return nil, err
	}
	gormigrate := bootstrap.NewMigrate(db)
	jobs := job.NewJobs(db, logger, settingRepo, certRepo, backupRepo, cacheRepo, taskRepo)
	cron, err := bootstrap.NewCron(koanf, logger, jobs)
	if err != nil {
		return nil, err
	}
	validation := bootstrap.NewValidator(koanf, db)
	web := app.NewWeb(koanf, mux, server, gormigrate, cron, queue, validation)
	return web, nil
}
