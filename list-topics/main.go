package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()

	brokerAddress := []string{"localhost:9092"}

	consumer, err := sarama.NewConsumer(brokerAddress, config)
	if err != nil {
		log.Panic(err)
	}

	topics, err := consumer.Topics()
	if err != nil {
		log.Panic(err)
	}
	for _, topic := range topics {
		fmt.Println(topic)
	}
}
