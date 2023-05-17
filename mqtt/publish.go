package mqtt

import (
	"fmt"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// PublishMessage publishes a message to a specified topic
func PublishMessage(client MQTT.Client, topic string, message string) {
	token := client.Publish(topic, 0, false, message)
	token.Wait()

	if token.Error() != nil {
		log.Fatal(token.Error())
	}

	fmt.Printf("Message published to topic: %s\n", topic)
}
