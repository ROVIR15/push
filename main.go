package main

import (
	"log"
	"time"

	MQTT "github.com/ROVIR15/push-notification-service/mqtt"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error laoding .env file")
	}
}

func main() {

	token := MQTT.InitMQTT()

	topics := []string{
		"submitted",
		"review",
		"approval",
	}

	tokens := MQTT.SubscribeToTopics(token, topics)
	for _, token := range tokens {
		if token.Error() != nil {
			log.Fatal(token.Error())
		}
	}

	// Keep the main program running to receive messages
	for {
		select {
		case <-time.After(1 * time.Second):
			// Perform any other tasks or checks here
		}
	}

	// // Execute the SQL query
	// rows, err := DB.Query("SELECT * FROM notification")

	// // Error
	// if err != nil {
	// 	log.Fatal("Error executing query:", err)
	// }
	// defer rows.Close()

	// // Get the column names
	// columns, err := rows.Columns()
	// if err != nil {
	// 	log.Fatal("Error getting column names:", err)
	// }

	// // Create a slice to hold the row values
	// values := make([]interface{}, len(columns))
	// row := make([]interface{}, len(columns))
	// for i := range row {
	// 	values[i] = &row[i]
	// }

	// // Print column names
	// for _, col := range columns {
	// 	fmt.Printf("%s\t", col)
	// }
	// fmt.Println()

	// // Print row values
	// for rows.Next() {
	// 	err := rows.Scan(values...)
	// 	if err != nil {
	// 		log.Fatal("Error scanning row:", err)
	// 	}

	// 	for _, value := range row {
	// 		fmt.Printf("%v\t", value)
	// 	}
	// 	fmt.Println()
	// }

	// if err := rows.Err(); err != nil {
	// 	log.Fatal("Error iterating rows:", err)
	// }

}
