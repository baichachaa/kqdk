package app

import (
	"github.com/kardianos/service"
	"os"
)

type Program struct{}

func (p *Program) Start(s service.Service) error {
	appLogger.Info("开始服务")
	appCron.Start()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	appLogger.Info("停止服务")
	_ = appLogger.Sync()
	appCron.Stop()
	return nil
}

func SystemService() {

	var serviceConfig = &service.Config{
		Name:        "HyS6000",
		DisplayName: "HyS6000",
		Description: "S6000自动反馈，拒绝超期\nv230425.3",
	}

	prg := &Program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		appLogger.Error(err.Error())
		return
	}

	if err != nil {
		appLogger.Error(err.Error())
		return
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			err := s.Install()
			if err != nil {
				appLogger.Error("请以管理员权限运行")
				return
			}
			appLogger.Info("服务安装成功")
			return
		}

		if os.Args[1] == "remove" {
			err := s.Uninstall()
			if err != nil {
				appLogger.Error("请以管理员权限运行")
				return
			}
			appLogger.Info("服务卸载成功")
			return
		}
	}

	err = s.Run()
	if err != nil {
		appLogger.Error(err.Error())
		_ = appLogger.Sync()
	}

}
