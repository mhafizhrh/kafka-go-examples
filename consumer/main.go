package main

import (
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

var wg sync.WaitGroup

func main() {
	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Panic(err)
	}

	topic := "topic-name"

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		log.Println(err)
	}
	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Panic(err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		log.Println("Start consume topic:", topic)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				log.Println(string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}
