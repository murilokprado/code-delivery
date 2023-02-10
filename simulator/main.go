package main

import (
	"fmt"
	"log"

	"github.com/codeedu/imersaofsfc2-simulator/infra/kafka"
	"github.com/joho/godotenv"

	kafka2 "github.com/codeedu/imersaofsfc2-simulator/application/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	messageChannel := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(messageChannel)

	go consumer.Consume()

	for msg := range messageChannel {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}
