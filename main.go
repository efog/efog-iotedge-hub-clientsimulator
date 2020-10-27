package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	zmq "github.com/pebbe/zmq4"
	zap "go.uber.org/zap"
)

func subscriberThread(endpoint *string) {
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
		log.Printf("Received %s", msg)
	}
}

func publisherThread(endpoint *string) {
	publisher, _ := zmq.NewSocket(zmq.PUB)
	publisher.Bind(*endpoint)
	for {
		s := fmt.Sprintf("%c-%05d", rand.Intn(3)+'A', rand.Intn(100000))
		log.Printf("Sending %s", s)
		_, err := publisher.SendMessage(s)
		if err != nil {
			break //  Interrupted
		}
		time.Sleep(1000 * time.Millisecond) //  Wait for 1/10th second
	}
}

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	undo := zap.RedirectStdLog(logger)
	defer undo()

	var backendHost string
	var backendPort string
	var frontendPort string

	backendHost = os.Getenv("BACKEND_HOST")
	backendPort = os.Getenv("BACKEND_PORT")
	frontendPort = os.Getenv("FRONTEND_PORT")
	if backendHost == "" {
		backendHost = "localhost"
	}
	if backendPort == "" {
		backendPort = "56789"
	}
	if frontendPort == "" {
		frontendPort = "12345"
	}

	log.Print("Redirected standard library")
	log.Print("Starting client simulator")
	wantFrontEndBind := fmt.Sprintf("tcp://*:%s", frontendPort)
	wantBackEndConnect := fmt.Sprintf("tcp://%s:%s", backendHost, backendPort)
	log.Printf("Frontend endpoint %s", wantFrontEndBind)
	log.Printf("Backend endpoint %s", wantBackEndConnect)

	go subscriberThread(&wantBackEndConnect)
	publisherThread(&wantFrontEndBind)

}
