package main

import (
	"log"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	brokersList = "localhost:9092"
	topicName   = "topic"
	maxRetry    = 5
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = maxRetry
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(strings.Split(brokersList, ","), config)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Panic(err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: topicName,
		Value: sarama.StringEncoder("Test message. Hello :)"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Message is stored in topic: %s, partition: %d, offset: %d\n", topicName, partition, offset)
}
