package app

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type settingsStruct struct {
	Debug bool `yaml:"Debug"`
	Mssql struct {
		IP       string `yaml:"IP"`
		Port     int    `yaml:"Port"`
		Database string `yaml:"Database"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
	} `yaml:"Mssql"`
	Mqtt struct {
		Broker   string `yaml:"Broker"`
		Port     int    `yaml:"Port"`
		ClientID string `yaml:"ClientID"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
	} `yaml:"Mqtt"`
	In struct {
		Index int    `yaml:"Index"`
		Start string `yaml:"Start"`
		End   string `yaml:"End"`
		Cron  string `yaml:"Cron"`
	} `yaml:"In"`
	Out struct {
		Index int    `yaml:"Index"`
		Start string `yaml:"Start"`
		End   string `yaml:"End"`
		Cron  string `yaml:"Cron"`
	} `yaml:"Out"`
	Devices struct {
		Sn      string `yaml:"Sn"`
		Message string `yaml:"Message"`
	} `yaml:"Devices"`
}

func SettingsService() {
	config, err := os.ReadFile("settings.yaml")
	if err != nil {
		appLogger.Error(err.Error())
		panic(err)
	}
	err = yaml.Unmarshal(config, &settings)
	if err != nil {
		appLogger.Error(err.Error())
		panic(err)
	}
	appLogger.Info(fmt.Sprintf("%+v\n", settings))
	//fmt.Println(string(utils.StructToJsonByte(settings)))
	saveSettings()
}

func saveSettings() {
	bt, err := yaml.Marshal(&settings)
	if err != nil {
		appLogger.Error(err.Error())
		return
	}
	f, err := os.Create("settings.yaml")
	defer f.Close()
	if err != nil {
		appLogger.Error(err.Error())
		return
	}
	_, err = f.Write(bt)
	if err != nil {
		appLogger.Error(err.Error())
		return
	}

}
