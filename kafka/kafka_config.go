package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func StartKafka() {
	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "topic1",
		GroupID:  "g1",
		MaxBytes: 10,
	}
	reader := kafka.NewReader(conf)
	fmt.Println("listing on topic : ", conf.Topic)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error Occured", err)
			continue
		}
		fmt.Println("Message is : ", string(m.Value))
	}

}
