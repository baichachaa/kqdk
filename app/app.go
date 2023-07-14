package app

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/kardianos/service"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var appServ service.Service // interface
var appCron *cron.Cron
var appClient mqtt.Client //interface
var appLogger *zap.Logger
var appMysql *gorm.DB
var appSqlite *gorm.DB
var settings *settingsStruct
var isTest *bool

func Run(isDebug bool) {

	// 日志服务
	LogService(isDebug)
	// 系统服务安装
	InitService()
	// 设置读取
	SettingsService()
	// mssql服务
	MssqlService()
	// sqlite
	//SqliteService()
	//appMysql = appSqlite
	getInData(true)
	// 定时任务
	CreateCronJob()
	// mqtt 服务
	MqttInit()
	// 系统服务运行
	ServiceRun()
}
