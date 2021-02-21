package main

import (
	"log"
	"os"
	"time"

	_ "github.com/CodyGuo/godaemon"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var token string
	if len(os.Args) == 3 {
		token = os.Args[2]
	}
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}
	bu := tb.InlineButton{Unique: "bu", Text: "2"}
	b.Handle(&bu, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			Text:      c.Data,
			ShowAlert: true,
		})

	})
	b.Start()
}
