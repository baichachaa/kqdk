package app

import (
	"flag"
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

var (
	// 初始化为 unknown，如果编译时没有传入这些值，则为 unknown
	Version         = "unknown"
	Revision        = "unknown"
	Branch          = "unknown"
	BuildDate       = "unknown"
	BuildTag        = "调度大楼"
	ServName        = "HrCheckIn"
	ServDisplayName = "智慧人资考勤打卡"
)

func InitService() {

	isInstall := flag.Bool("i", false, "system service install")
	isUnInstall := flag.Bool("u", false, "system service uninstall")
	isVersion := flag.Bool("v", false, "show bin version")
	isTest = flag.Bool("t", false, "connect test")

	flag.Parse()

	if *isVersion {
		fmt.Printf("  Version:          %s\n", Version)
		fmt.Printf("  Revision:         %s\n", Revision)
		fmt.Printf("  Branch:           %s\n", Branch)
		fmt.Printf("  BuildDate:        %s\n", BuildDate)
		fmt.Printf("  BuildTag:         %s\n", BuildTag)
		fmt.Printf("  ServName:         %s\n", ServName)
		fmt.Printf("  ServDisplayName:  %s\n", ServDisplayName)
		os.Exit(0)
	}

	var serviceConfig = &service.Config{
		Name:        ServName,
		DisplayName: ServDisplayName,
		Description: "门禁考勤数据自动推送至智慧人资\n" + Version + "\n" + BuildDate,
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

	if *isInstall == true {
		err := s.Install()
		if err != nil {
			fmt.Println(err.Error())
			appLogger.Error(err.Error())
			os.Exit(0)
		}
		fmt.Println("system service install success")
		appLogger.Info("system service install success")
		os.Exit(0)
	}

	if *isUnInstall == true {
		err := s.Uninstall()
		if err != nil {
			fmt.Println(err.Error())
			appLogger.Error(err.Error())
			os.Exit(0)
		}
		fmt.Println("system service uninstall success")
		appLogger.Info("system service uninstall success")
		os.Exit(0)
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
