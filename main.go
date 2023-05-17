package main

import (
	"database/sql"
	"log"
	"os"
	"os/signal"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"s
)

var DB *sql.DB

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error laoding .env file")
	}
}

func InitDB() {

	loadEnv()

	connStr := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	DB = db
}

var MiQ *MQTT.Client

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

func main() {

	InitDB()

	token := InitMQTT()

	topics := []string{
		"submitted",
		"review",
		"approval",
	}

	tokens := subscribeToTopics(token, topics)
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
