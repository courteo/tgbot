package ForBot

import (
	"MYtgbot/pkg/users"
	"context"
	"fmt"
	tgbotapi "github.com/skinass/telegram-bot-api/v5"
	"log"
	"net/http"
	"os"
)

var (
	// @BotFather в телеграме даст вам это
	BotToken = "5618658874:AAE5m7hdV4zwahkY-FbrCXtt8CahFOAA8QU"

	// урл выдаст вам игрок или хероку
	WebhookURL = "https://2a07-194-186-53-138.eu.ngrok.io"
)

func StartTaskBot(ctx context.Context) error {
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Fatalf("NewBotAPI failed: %s", err)
	}

	bot.Debug = true
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	wh, err := tgbotapi.NewWebhook(WebhookURL)
	if err != nil {
		log.Fatalf("NewWebhook failed: %s", err)
	}

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatalf("SetWebhook failed: %s", err)
	}

	updates := bot.ListenForWebhook("/")

	http.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		_, err1 := w.Write([]byte("all is working"))
		if err1 != nil {
			fmt.Println("не работает")
			return
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	go func() {
		log.Fatalln("http err:", http.ListenAndServe(":"+port, nil))
	}()
	fmt.Println("start listen :" + port)

	// получаем все обновления из канала updates
	for update := range updates {
		if update.Message == nil {
			continue
		}

		currUser, err := users.GetUser(update.Message.From.UserName)
		if err != nil {
			currUser = users.User{UserName: update.Message.From.UserName, ChatId: update.Message.Chat.ID}
			users.AllUsers = append(users.AllUsers, currUser)
		}

		if update.Message.IsCommand() {
			ForCommand(*bot, currUser, update)
		} else {
			_, err1 := bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"Привет, напиши /help для команд",
			))
			if err1 != nil {
				fmt.Println("ошибка при отправке")
				return err1
			}
		}
	}
	return nil
}
