package main

import (
	tgClient "articles-tbot/clients/telegram"
	event_consumer "articles-tbot/consumer/event-consumer"
	"articles-tbot/events/telegram"
	"articles-tbot/files_storage"
	"articles-tbot/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(files_storage.StoragePath),
	)

	log.Print("service has been started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
