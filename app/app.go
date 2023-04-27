package app

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var appCron *cron.Cron
var appClient mqtt.Client //interface
var appLogger *zap.Logger
var appMssql *gorm.DB
var appSqlite *gorm.DB
var settings *settingsStruct

func Run(isDebug bool) {

	// 日志服务
	LogService(isDebug)
	// 设置读取
	SettingsService()
	// mssql服务
	MssqlService()
	// sqlite
	//SqliteService()
	//appMssql = appSqlite
	//getInData(true)
	// 定时任务
	CreateCronJob()
	// mqtt 服务
	MqttInit()
}
