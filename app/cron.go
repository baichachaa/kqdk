package app

import (
	"github.com/robfig/cron/v3"
)

func CreateCronJob() {

	appCron = cron.New()

	_, err := appCron.AddFunc(settings.In.Cron, cronIn)
	if err != nil {
		appLogger.Error(err.Error())
	}

	_, err = appCron.AddFunc(settings.Out.Cron, cronOut)
	if err != nil {
		appLogger.Error(err.Error())
	}
}
