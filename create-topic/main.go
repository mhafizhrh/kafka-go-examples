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

	// err = clusterAdmin.CreateTopic("topic-name", &sarama.TopicDetail{NumPartitions: 1, ReplicationFactor: 3}, false)
	err = clusterAdmin.CreateTopic("topic-name", &sarama.TopicDetail{NumPartitions: 1, ReplicationFactor: 1}, false)
	if err != nil {
		log.Panic(err)
	}

	log.Println("create topic success")
}
