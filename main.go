package main

import (
	"hanamaru/commands/debug"
	"hanamaru/commands/fun"
	"hanamaru/commands/image"
	"hanamaru/commands/info"
	"hanamaru/commands/music"
	"hanamaru/events"
	"hanamaru/hanamaru"
	"os"
	"os/signal"
	"syscall"
)

var TOKEN string
var bot *hanamaru.Hanamaru

var optionals []*hanamaru.Command

func main() {
	var syscallChan = make(chan os.Signal)
	defer bot.Close()

	bot = hanamaru.New("Bot "+TOKEN, "!")

	bot.AddCommand(info.About)
	bot.AddCommand(debug.ListArgs)
	bot.AddCommand(image.Rumble)
	bot.AddCommand(image.CAA)
	bot.AddCommand(info.Avatar)
	bot.AddCommand(image.Jpg)

	bot.AddCommand(music.Leave)
	bot.AddCommand(music.Join)
	bot.AddCommand(music.Play)

	bot.AddCommand(fun.Dance)

	for _, command := range optionals {
		bot.AddCommand(command)
	}

	bot.AddHandler(events.Nhentai)

	signal.Notify(syscallChan, syscall.SIGTERM, syscall.SIGINT)
	<-syscallChan
}
