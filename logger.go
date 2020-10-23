package main

import (
	"log"

	"go.uber.org/zap"
)

// ScaffoldLogger Scaffolds zap logger
func ScaffoldLogger() {
	logger := zap.NewExample()
	defer logger.Sync()

	undo := zap.RedirectStdLog(logger)
	defer undo()
	log.Print("redirected standard library")
}
