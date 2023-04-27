package app

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
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

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		appLogger.Error("MQTT 首次连接失败")
		_ = appLogger.Sync()
		panic(token.Error())
	}

	appClient = client
}

func publish(client mqtt.Client, text []byte) {

	//text := fmt.Sprintf("Message %d", i)
	token := client.Publish("/v1/devices/SNs-HT-XT-BenBuDaLou-RLSB/datas", 2, false, text)

	// token是阻塞，需要启动多线程
	go func() {
		_ = token.WaitTimeout(30 * time.Second) // Can also use '<-t.Done()' in releases > 1.2.0
		if token.Error() != nil {
			appLogger.Error(token.Error().Error()) // Use your preferred logging technique (or just fmt.Printf)
		}
	}()

}
