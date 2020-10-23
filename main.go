package main

import (
	"log"

	efogIotEdgeHubServer "github.com/efog/efog-iotedge-hub"
	zap "go.uber.org/zap"
)

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
	server.Run()
}
