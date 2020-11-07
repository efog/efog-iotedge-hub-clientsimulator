package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func publish(publisher *zmq.Socket) {
	for {
		topic := rand.Intn(3) + 'A'
		s := fmt.Sprintf("%c-%05d", topic, rand.Intn(100000))
		log.Printf("Sending %s to topic %c", s, topic)
		_, err := publisher.SendMessage(fmt.Sprintf("%c", topic), s)
		if err != nil {
			break //  Interrupted
		}
		time.Sleep(1000 * time.Millisecond) //  Wait for 1/10th second
	}
}

func main() {
	publisherEndpoint := os.Getenv("PUBLISHER_ENDPOINT")
	subscriberEndpoint := os.Getenv("SUBSCRIBER_ENDPOINT")
	if &publisherEndpoint == nil {
		publisherEndpoint = "tcp://localhost:7000"
	}
	if &subscriberEndpoint == nil {
		subscriberEndpoint = "tcp://localhost:7001"
	}

	log.Print("Starting pub server")
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Connect(publisherEndpoint)
	log.Printf("Connected publisher to endpoint %s", publisherEndpoint)

	log.Print("Starting sub server")
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect(subscriberEndpoint)
	subscriber.SetSubscribe("A")
	subscriber.SetSubscribe("B")
	subscriber.SetSubscribe("C")
	log.Printf("Connected subscriber to endpoint %s", subscriberEndpoint)

	go publish(publisher)

	for {
		msg, _ := subscriber.RecvMessage(0)
		log.Printf("Received %s", msg)
	}

}
