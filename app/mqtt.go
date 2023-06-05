package app

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	appLogger.Info(fmt.Sprintf("MQTT Received message: %s from topic: %s\n", msg.Payload(), msg.Topic()))
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	appLogger.Info("MQTT Connected")
}

var reConnectHandler mqtt.ReconnectHandler = func(client mqtt.Client, options *mqtt.ClientOptions) {
	appLogger.Info("MQTT ReConnection")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	appLogger.Error(fmt.Sprintf("MQTT Connect lost: %v\n", err))
}

func MqttInit() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", settings.Mqtt.Broker, settings.Mqtt.Port))
	opts.SetClientID(settings.Mqtt.ClientID)
	opts.SetUsername(settings.Mqtt.Username)
	opts.SetPassword(settings.Mqtt.Password)

	opts.DefaultPublishHandler = messagePubHandler
	opts.OnConnect = connectHandler
	opts.OnReconnecting = reConnectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)

	if *isTest {
		mqttInitConnect(client)
	} else {
		go mqttInitConnect(client)
	}

	appClient = client
}

func mqttInitConnect(client mqtt.Client) {
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		if *isTest {
			appLogger.Error("MQTT 登录失败")
			_ = appLogger.Sync()
			fmt.Println(token.Error())
			os.Exit(0)
		} else {
			appLogger.Error("MQTT 登录失败，正在重试")
			mqttInitConnect(client)
		}
	} else {
		appLogger.Info("MQTT 登录成功")
	}
}
