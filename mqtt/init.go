package mqtt

import (
	"log"
	"os"
	"os/signal"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var Client *MQTT.Client

func InitMQTT() MQTT.Client {
	// Create an MQTT client options
	opts := MQTT.NewClientOptions()

	hostMqtt := os.Getenv("MQTT_HOST")
	portMqtt := os.Getenv("MQTT_PORT")

	connStr := "ws://" + hostMqtt + ":" + portMqtt + "/mqtt"

	opts.AddBroker(connStr) // Replace with your broker's address

	// Set client ID
	opts.SetClientID("mqttx_24d9a248")

	// Set optional username and password
	opts.SetUsername("newsecurity")
	opts.SetPassword("Testing1234")

	// Create a new MQTT client
	client := MQTT.NewClient(opts)

	// Set up signal handler to gracefully close the client
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		client.Disconnect(250)
		os.Exit(0)
	}()

	// Connect to the MQTT broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	return client
}
