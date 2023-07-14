package app

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type settingsStruct struct {
	Debug    bool `yaml:"Debug"`
	DataBase struct {
		IP       string `yaml:"IP"`
		Port     int    `yaml:"Port"`
		Database string `yaml:"Database"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
	} `yaml:"DataBase"`
	Mqtt struct {
		Broker   string `yaml:"Broker"`
		Port     int    `yaml:"Port"`
		ClientID string `yaml:"ClientID"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
		Qos      uint8  `yaml:"Qos"`
	} `yaml:"Mqtt"`
	In struct {
		Index string `yaml:"Index"`
		Start string `yaml:"Start"`
		End   string `yaml:"End"`
		Cron  string `yaml:"Cron"`
	} `yaml:"In"`
	Out struct {
		Index string `yaml:"Index"`
		Start string `yaml:"Start"`
		End   string `yaml:"End"`
		Cron  string `yaml:"Cron"`
	} `yaml:"Out"`
	Devices struct {
		DevicesSn string `yaml:"DevicesSn"`
		ServiceId string `yaml:"ServiceId"`
	} `yaml:"Devices"`
}

func SettingsService() {
	config, err := os.ReadFile("settings.yaml")
	if err != nil {
		appLogger.Error(err.Error())
		fmt.Println(err)
		os.Exit(0)
	}
	settings = &settingsStruct{}
	err = yaml.Unmarshal(config, &settings)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
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
