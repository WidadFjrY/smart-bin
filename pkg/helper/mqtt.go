package helper

import (
	"os"
	"smart-trash-bin/domain/web"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

func NewMQTT(mqttConf web.MQTTRequest) bool {
	godotenv.Load()

	broker := os.Getenv("MQTT_BROKER")
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(mqttConf.ClientId)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token := client.Publish(mqttConf.Topic, 1, false, mqttConf.Payload)
	token.Wait()

	done := make(chan bool)

	token = client.Subscribe(mqttConf.Topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		if string(msg.Payload()) == mqttConf.MsgResp {
			client.Disconnect(250)
			done <- true
		}
		done <- false
	})

	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	timeOut, _ := strconv.Atoi(os.Getenv("TIMEOUT_MQTT_SUB"))

	select {
	case <-done:
		return true
	case <-time.After(time.Duration(timeOut) * time.Second):
		client.Disconnect(250)
		return false
	}
}
