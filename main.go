package main

import (
	tgClient "articles-tbot/clients/telegram"
	event_consumer "articles-tbot/consumer/event-consumer"
	"articles-tbot/events/show_schedule_telegram"
	"articles-tbot/schedules_storage"
	"articles-tbot/storage/schedules"
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func main() {
	eventsProcessor := show_schedule_telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		schedules.New(schedules_storage.StoragePath),
	)

	log.Print("service has been started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"t",
		"",
		"token for access to save_articles_telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
