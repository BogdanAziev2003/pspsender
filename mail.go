package pkg

import (
	"gopkg.in/gomail.v2"
)

func (bot *Bot) SendMailMessage(subject, text string) {
	go func() {
		bot.StartedFunc++
		for _, mail := range bot.Emails_to_send {
			message := gomail.NewMessage()
			message.SetHeader("From", bot.Email_address)
			message.SetHeader("To", mail)
			message.SetHeader("Subject", subject)
			message.SetBody("text/plain", text)

			a := gomail.NewDialer("smtp.gmail.com", 587, bot.Email_address, bot.Email_pass)
			//log.Printf("\nto_send: %s\nadress: %s\npass: %s\n", mail, bot.Email_address, bot.Email_pass)
			err := a.DialAndSend(message)
			if err != nil {
				break
			}
		}
		bot.Done <- true
	}()
}
