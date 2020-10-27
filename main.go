package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	efogIotEdgeHubServer "github.com/efog/efog-iotedge-hub"
	zmq "github.com/pebbe/zmq4"
	zap "go.uber.org/zap"
)

func subscriber_thread(endpoint *string) {
	//  Subscribe to "A" and "B"
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	subscriber.Connect(*endpoint)
	subscriber.SetSubscribe("A")
	subscriber.SetSubscribe("B")
	subscriber.SetSubscribe("C")
	defer subscriber.Close() // cancel subscribe
	for {
		msg, err := subscriber.RecvMessage(0)
		if err != nil {
			break //  Interrupted
		}
		log.Printf("Received %q", msg)
	}
}

func publisher_thread(endpoint *string) {
	publisher, _ := zmq.NewSocket(zmq.PUB)
	publisher.Bind(*endpoint)
	for {
		s := fmt.Sprintf("%c-%05d", rand.Intn(3)+'A', rand.Intn(100000))
		log.Printf("Sending %q", s)
		_, err := publisher.SendMessage(s)
		if err != nil {
			break //  Interrupted
		}
		time.Sleep(10 * time.Millisecond) //  Wait for 1/10th second
	}
}

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	undo := zap.RedirectStdLog(logger)
	defer undo()
	log.Print("redirected standard library")
	log.Print("Starting client simulator")
	wantFrontEndBind := "tcp://*:12345"
	wantFrontEndConnect := "tcp://localhost:12345"
	wantBackEndBind := "tcp://*:56789"
	wantBackEndConnect := "tcp://localhost:56789"
	server := efogIotEdgeHubServer.NewServer(&wantBackEndBind, &wantBackEndConnect, &wantFrontEndBind, &wantFrontEndConnect)

	go publisher_thread(&wantFrontEndBind)
	go subscriber_thread(&wantBackEndConnect)
	
	server.Run()
}
