package main

import (
	"flag"
	"github.com/kr/pretty"
	"github.com/tucnak/telebot"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var (
	token_file = flag.String("token", "", "File contains telegram bot token")
)

func main() {
	var msg, country string
	flag.Parse()
	if *token_file == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	token, err := ioutil.ReadFile(*token_file)
	if err != nil {
		log.Fatal(err)
	}
	bot, err := telebot.NewBot(string(token))
	if err != nil {
		return
	}

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		// pretty.Println(message.Sender)
		msg = message.Text
		if msg == "/hi" {
			count := 1
			for {
				pretty.Println(count)
				count++
				bot.SendMessage(message.Chat,
					"Hello, "+message.Sender.FirstName+"!", nil)
				time.Sleep(1000 * time.Millisecond)
			}
		} else if strings.HasPrefix(msg, "/flag") {
			//check if flag is empty
			country = "ASEAN" //msg[6:]
			pretty.Print(country)
			photo := "./resources/flags/" + country + ".png"
			boom, err := telebot.NewFile(photo)
			if err != nil {
				pretty.Print(err)
			}
			pretty.Print(&bot)
			pretty.Print(&boom)

			// SendPhoto
			// telebot.File{}ASEAN&telebot.File{FileID:"", FileSize:0, filename:"./resources/flags/ASEAN.png"}
			// pretty.Print(reflect.TypeOf((*bot).SendMessage))
			// // get from directory
			// err = bot.SendAudio(message.Chat, &boom, nil)
			// err = bot.SendMessage(message.Chat, &boom, nil)
			if err != nil {
				pretty.Print(err)
			}
		}
	}
}
