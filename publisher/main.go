package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	// Allow auto create topic, default true
	// config.Metadata.AllowAutoTopicCreation = false

	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		return
	}
	defer client.Close()

	messages := []*sarama.ProducerMessage{
		{
			Topic: "topic-name",
			Value: sarama.StringEncoder("{“status”:”OK”, “data”: “hello world”}"),
		},
		{
			Topic: "topic-name",
			Value: sarama.StringEncoder("{“status”:”OK”, “data”: “create table failed”}"),
		},
	}

	for _, msg := range messages {
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			log.Panic(err)
		}

		log.Printf("Published. pid:%v offset:%v\n", pid, offset)
	}
}
