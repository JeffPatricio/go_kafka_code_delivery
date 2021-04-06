package main

import (
	"fmt"
	"log"

	"github.com/codeedu/imersaofsfc-simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {

	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)

	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
	}

	// route := route.Route{ID: "1", ClientID: "1"}
	// route.LoadPositions()
	// stringjson, _ := route.ExportsJsonPositions()
	// fmt.Println(stringjson[0])
}
