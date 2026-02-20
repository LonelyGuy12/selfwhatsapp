package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	bot "selfwhatsapp/src"
)

func main() {
	ctx := context.Background()

	store, err := bot.NewStore(ctx)
	if err != nil {
		panic(err)
	}

	client, err := bot.NewClient(ctx, store)
	if err != nil {
		panic(err)
	}

	if err := bot.Connect(ctx, client); err != nil {
		panic(err)
	}

	fmt.Println("Connected and running")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	fmt.Println("Shutting down...")
	client.Disconnect()
}
