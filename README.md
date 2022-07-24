# pspsender
 ### Для корректной работы, в файле <strong>.env</strong> нужно указать такие переменные:
  - TELEGRAM_TOKEN ( string )- токен телеграм бота от которого вы хотите получать сообщения
  - CHATS_ID ( []int ) - id чатов в телеграм, куда бот будет присылать сообщения ( указываются через пробел )
  - TO_MAILS ( []string ) - адреса электронных почт, куда бот будет присылать письма (также указываются через пробел)
  
  
  
## Getting started
  
  

  ```go
 bot := psp.NewBot() //создание бота
 defer bot.Close()
  
bot.SendTgMessage("Текст сообщения") //отправка сообщений в Телеграм чаты

bot.SendMailMessage("Заголовок", "Тело сообщения") //отправка email письм 
  ```
