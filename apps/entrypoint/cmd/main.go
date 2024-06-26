package main

import (
	"context"
	"log"
	"os"

	"github.com/Mad-Pixels/golang-etc-playground/apps"
	"github.com/Mad-Pixels/golang-etc-playground/apps/entrypoint"
)

func main() {
	logLevel := os.Getenv(apps.LogLevelEnv)
	if logLevel == "" {
		logLevel = apps.LogLevelValue
	}
	listenPort := os.Getenv(apps.ListenPortEnv)
	if listenPort == "" {
		listenPort = apps.ListenPortValue
	}

	app, err := entrypoint.New(listenPort, logLevel)
	if err != nil {
		log.Fatal(err)
	}
	app.Run(context.Background())
}
