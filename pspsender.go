package pkg

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Sender interface {
	SendTgMessage(string) error
	SendMailMessage(string, string) error
}

type Bot struct {
	BotToken       string
	BotApi         string
	BotUrl         string
	Chats_Id       []int
	Email_pass     string
	Email_address  string
	Emails_to_send []string
	Done           chan bool
	StartedFunc    int
}

func NewBot() *Bot {
	Done := make(chan bool)
	api := "https://api.telegram.org/bot"
	token := os.Getenv("TELEGRAM_TOKEN")
	email_pass := "zjpzqctdasfsehom"
	email_address := "sender.psp2022@gmail.com"

	chats_id := strings.Split(os.Getenv("CHATS_ID"), " ")
	slice_of_chat_id := make([]int, 0)
	for _, val := range chats_id {
		intVal, _ := strconv.Atoi(val)
		slice_of_chat_id = append(slice_of_chat_id, intVal)
	}

	emails_to_send := strings.Split(os.Getenv("TO_MAILS"), " ")
	slice_of_emails_to_send := make([]string, 0)
	for _, val := range emails_to_send {
		slice_of_emails_to_send = append(slice_of_emails_to_send, val)
	}

	url := api + token
	return &Bot{
		BotToken:       token,
		BotApi:         api,
		BotUrl:         url,
		Chats_Id:       slice_of_chat_id,
		Email_pass:     email_pass,
		Email_address:  email_address,
		Emails_to_send: slice_of_emails_to_send,
		Done:           Done,
		StartedFunc:    0,
	}
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func (bot *Bot) Close() {
	for i := 0; i < bot.StartedFunc; i++ {
		<-bot.Done
	}
}
