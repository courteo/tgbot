package main

// https://api.telegram.org/bot5244227470:AAEModcsPOS8TxZehTmFoTwH5Kr3mctcMv0/getUpdates

import (
	"MYtgbot/pkg/ForBot"
	"context"
)

func main() {
	err := ForBot.StartTaskBot(context.Background())
	if err != nil {
		panic(err)
	}
}
