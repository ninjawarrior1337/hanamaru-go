package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/ninjawarrior1337/hanamaru-go/events"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/spf13/viper"
)

var config *viper.Viper

var commands []*framework.Command
var optionalEvents []*framework.EventListener

func init() {
	config = viper.New()
	if os.Getenv("IN_DOCKER") != "" {
		config.AddConfigPath("/data")
	} else {
		config.AddConfigPath(".")
	}
	config.AllSettings()
	config.AutomaticEnv()
	config.WatchConfig()

	config.SetDefault("owner", "")
	config.SetDefault("prefix", "!")
	config.SetDefault("token", "")
	config.SetDefault("listening", "地元愛 Dash")
	config.SetDefault("playing", "")

	err := config.ReadInConfig()
	if err != nil {
		_ = config.WriteConfigAs("config.yml")
		log.Printf("Failed to read config: %v\n", err)
		log.Println("A default config file has been created for you at config.yml")
		log.Fatalln("Or you may set environment variables...Exiting now.")
	}

	if !config.IsSet("token") || len(config.GetString("token")) == 0 {
		log.Fatalln("A token must be set in the config.")
	}
}

var memprofile = flag.String("memprofile", "", "write memory profile to a file")

func main() {
	flag.Parse()
	var syscallChan = make(chan os.Signal)

	bot := framework.New("Bot "+config.GetString("token"), config.GetString("prefix"), config.GetString("owner"))
	defer bot.Close()

	bot.EnableHelpCommand()

	err := bot.SetupDB()
	if err != nil {
		fmt.Println("Failed to setup db: " + err.Error())
	}

	for _, command := range commands {
		bot.AddCommand(command)
	}

	for _, event := range optionalEvents {
		bot.AddEventListener(event)
	}

	bot.AddEventListener(events.Nhentai)
	bot.AddEventListener(events.ReactionExpansion)
	bot.AddEventListener(events.Roboragi)
	bot.AddEventListener(events.BigEmoji)
	bot.AddEventListener(events.AwardsAddHandler)
	bot.AddEventListener(events.AwardsRemoveHandler)

	go func() {
		setStatus := func() {
			if playing := config.GetString("playing"); playing != "" {
				bot.Session.UpdateListeningStatus(playing)
			} else if listening := config.GetString("listening"); listening != "" {
				bot.Session.UpdateListeningStatus(listening)
			} else {
				bot.Session.UpdateListeningStatus("Aqours' Songs")
			}
		}
		setStatus()
		config.OnConfigChange(func(in fsnotify.Event) {
			setStatus()
		})
		for range time.Tick(time.Hour * 2) {
			setStatus()
		}
	}()

	signal.Notify(syscallChan, syscall.SIGTERM, syscall.SIGINT)
	<-syscallChan
	if *memprofile != "" {
		doMemprofile()
	}
}

func doMemprofile() {
	f, err := os.Create(*memprofile)
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()
	runtime.GC()
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
}
