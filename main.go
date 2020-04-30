package main

import (
	"github.com/spf13/viper"
	"hanamaru/commands/debug"
	"hanamaru/commands/fun"
	"hanamaru/commands/image"
	"hanamaru/commands/info"
	"hanamaru/commands/music"
	"hanamaru/events"
	"hanamaru/hanamaru"
	"log"
	"os"
	"os/signal"
	"syscall"
)

//go:generate pkger

var config *viper.Viper

var optionals []*hanamaru.Command

func init() {
	config = viper.New()
	config.AddConfigPath(".")
	config.AllSettings()
	config.AutomaticEnv()

	config.SetDefault("owner", "")
	config.SetDefault("prefix", "!")
	config.SetDefault("token", "")

	err := config.ReadInConfig()
	if err != nil {
		_ = config.WriteConfigAs("config.yml")
		log.Printf("Failed to read config: %v\n", err)
		log.Fatalln("A default config file has been created for you at config.yml")
	}

	if !config.IsSet("token") || len(config.GetString("token")) == 0 {
		log.Fatalln("A token must be set in the config.")
	}
}

func main() {
	var syscallChan = make(chan os.Signal)

	bot := hanamaru.New("Bot "+config.GetString("token"), config.GetString("prefix"))
	defer bot.Close()

	bot.SetOwner(config.GetString("owner"))
	bot.EnableHelpCommand()

	bot.AddCommand(debug.ListArgs)

	bot.AddCommand(info.About)
	bot.AddCommand(info.Avatar)

	bot.AddCommand(image.Rumble)
	bot.AddCommand(image.CAS)
	bot.AddCommand(image.Jpg)
	bot.AddCommand(image.Latex)
	bot.AddCommand(image.Stretch)
	bot.AddCommand(image.Bishop)

	bot.AddCommand(music.Leave)
	bot.AddCommand(music.Join)
	bot.AddCommand(music.Play)

	bot.AddCommand(fun.Dance)
	bot.AddCommand(fun.Suntsu)

	for _, command := range optionals {
		bot.AddCommand(command)
	}

	bot.AddHandler(events.Nhentai)
	bot.AddHandler(events.Boomer)
	bot.AddHandler(events.RepeatMessage)

	signal.Notify(syscallChan, syscall.SIGTERM, syscall.SIGINT)
	<-syscallChan
}
