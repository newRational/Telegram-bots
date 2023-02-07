package main

import (
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	// tgClient := telegram.New(tgBotHost, mustToken())

	// fetcher = fethcer.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)
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
