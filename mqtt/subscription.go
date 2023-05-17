package mqtt

import (
	"fmt"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func SubscribeToTopics(client MQTT.Client, topics []string) []MQTT.Token {
	var tokens []MQTT.Token

	// Define the callback function to handl incoming message
	callback := func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("Received message on topic: %s - Message: %s\n", msg.Topic(), string(msg.Payload()))
	}

	// Subscribe to the specified topics with the defined callback function
	for _, topic := range topics {
		token := client.Subscribe(topic, 0, callback)
		if token.Wait() && token.Error() != nil {
			log.Fatal(token.Error())
		}
		fmt.Printf("Subscribed to topic: %s\n", topic)
		tokens = append(tokens, token)
	}

	return tokens
}
