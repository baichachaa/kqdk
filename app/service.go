package app

import (
	"fmt"
	"github.com/kardianos/service"
	"os"
)

type Program struct{}

func (p *Program) Start(s service.Service) error {
	fmt.Println("开始服务")
	appLogger.Info("开始服务")
	appCron.Start()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	fmt.Println("停止服务")
	appLogger.Info("停止服务")
	_ = appLogger.Sync()
	appCron.Stop()
	return nil
}

func InitService() {
	var serviceConfig = &service.Config{
		Name:        "zhrz-kqdk",
		DisplayName: "智慧人资-考勤打卡",
		Description: "门禁考勤数据自动推送至智慧人资\nv230428.3",
	}

	prg := &Program{}
	s, err := service.New(prg, serviceConfig)
	appServ = s

	if err != nil {
		fmt.Println(err.Error())
		appLogger.Error(err.Error())
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		appLogger.Error(err.Error())
		return
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			err := s.Install()
			if err != nil {
				fmt.Println(err.Error())
				appLogger.Error(err.Error())
				os.Exit(0)
			}
			fmt.Println("服务安装成功")
			appLogger.Info("服务安装成功")
			os.Exit(0)
		}

		if os.Args[1] == "remove" {
			err := s.Uninstall()
			if err != nil {
				fmt.Println(err.Error())
				appLogger.Error(err.Error())
				os.Exit(0)
			}
			fmt.Println("服务卸载成功")
			appLogger.Info("服务卸载成功")
			os.Exit(0)
		}
	}

}

func ServiceRun() {

	err := appServ.Run()
	if err != nil {
		fmt.Println("系统服务启动异常")
		fmt.Println(err.Error())
		appLogger.Error(err.Error())
		_ = appLogger.Sync()
	}

}
