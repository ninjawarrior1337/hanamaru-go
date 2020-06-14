package main

import (
	"fmt"
	"github.com/ninjawarrior1337/hanamaru-go/events"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var config *viper.Viper

var commands []*framework.Command
var optionalEvents []interface{}

func init() {
	config = viper.New()
	config.AddConfigPath(".")
	config.AllSettings()
	config.AutomaticEnv()

	config.SetDefault("owner", "")
	config.SetDefault("prefix", "!")
	config.SetDefault("token", "")
	config.SetDefault("listening", "")
	config.SetDefault("playing", "")

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

	bot := framework.New("Bot "+config.GetString("token"), config.GetString("prefix"))
	defer bot.Close()

	bot.SetOwner(config.GetString("owner"))
	bot.EnableHelpCommand()

	err := bot.SetupDB()
	if err != nil {
		fmt.Println("Failed to setup db: " + err.Error())
	}

	for _, command := range commands {
		bot.AddCommand(command)
	}

	for _, event := range optionalEvents {
		bot.AddHandler(event)
	}

	bot.AddHandler(events.Nhentai)
	bot.AddHandler(events.ReactionExpansion)
	bot.AddHandler(events.Roboragi)

	if playing := config.GetString("playing"); playing != "" {
		bot.Session.UpdateListeningStatus(playing)
	} else if listening := config.GetString("listening"); listening != "" {
		bot.Session.UpdateListeningStatus(listening)
	} else {
		bot.Session.UpdateListeningStatus("Aqours' Songs")
	}

	signal.Notify(syscallChan, syscall.SIGTERM, syscall.SIGINT)
	<-syscallChan
}
