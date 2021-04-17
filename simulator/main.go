package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	kafka2 "github.com/gdeggau/imersaofsfc2-simulator/application/kafka"
	"github.com/gdeggau/imersaofsfc2-simulator/infra/kafka"
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
		go kafka2.Produce(msg)
	}
	// route := route.Route{
	// 	ID: "1",
	// 	ClientID: "1",
	// }
	// route.LoadPositions()
	// stringjson,_ := route.ExportJsonPositions()
	// fmt.Println(stringjson)
}