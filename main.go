package main

import (
	"amine-ameur/go_kafka/kafka"
	"fmt"
	"time"
)

func main() {
	fmt.Println("starting connexion with kafka ...")
	go kafka.StartKafka()
	time.Sleep(10 * time.Minute)
}
