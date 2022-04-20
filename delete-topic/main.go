package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()

	brokerAddress := []string{"localhost:9092"}

	clusterAdmin, err := sarama.NewClusterAdmin(brokerAddress, config)
	if err != nil {
		log.Panic(err)
	}

	err = clusterAdmin.DeleteTopic("topic-name")
	if err != nil {
		log.Panic(err)
	}

	log.Println("delete topic success")
}
