package main

import (
	"flag"
	"log"

	tgClient "Sec/clients/telegram"
	eventconsumer "Sec/consumer/event-consumer"
	"Sec/events/telegram"
	"Sec/storage/files"
)

// 6900584168:AAH33akFmLTjv3GOg61OGjpTXcwL_IcmYhk

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(detcher, processor)
}

func mustToken() string {

	token := flag.String(
		"tg-bot-token",
		"",
		"token for acces to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token

}
