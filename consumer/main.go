package main

import (
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	brokersList       = "localhost:9092"
	topicName         = "topic"
	partition         = 0
	offset            = -1
	messageCountStart = 0
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	master, err := sarama.NewConsumer(strings.Split(brokersList, ","), config)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			log.Panic(err)
		}
	}()

	consumer, err := master.ConsumePartition(topicName, int32(partition), int64(offset))
	if err != nil {
		log.Panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Println(err)
			case msg := <-consumer.Messages():
				messageCountStart++
				log.Println("Received message: ", string(msg.Key), string(msg.Value))
			case <-signals:
				log.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh

	log.Println("Processed", messageCountStart, "messages")
}
