package pkg

import (
	"net/http"
	"strconv"
)

func (bot *Bot) SendTgMessage(text string) {
	go func() {
		bot.StartedFunc++
		for _, chat_id := range bot.Chats_Id {
			_, _ = http.Post(bot.BotUrl+"/sendMessage?chat_id="+strconv.Itoa(chat_id)+"&text="+text, "text/html", nil)
		}
		bot.Done <- true
	}()
}
