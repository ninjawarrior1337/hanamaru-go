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

	"github.com/bwmarrin/discordgo"
	"github.com/fsnotify/fsnotify"

	"github.com/ninjawarrior1337/hanamaru-go/events"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/spf13/viper"
)

var config *viper.Viper

//go:generate go run tools/cmd/gen_command_imports.go
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

	err := config.ReadInConfig()
	if err != nil && os.Getenv("IN_DOCKER") == "" {
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
	var syscallChan = make(chan os.Signal, 32)

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

	// bot.AddEventListener(events.Nhentai): Scraping no longer viable
	// bot.AddEventListener(events.BigEmoji): Annoying and also ruins the humor
	bot.AddEventListener(events.ReactionExpansion)
	bot.AddEventListener(events.Roboragi)
	bot.AddEventListener(events.AwardsAddHandler)
	bot.AddEventListener(events.AwardsRemoveHandler)

	go func() {
		usd := discordgo.UpdateStatusData{
			Status:     "online",
			AFK:        false,
			Activities: []*discordgo.Activity{{}},
		}
		setStatus := func() {
			if playing := config.GetString("playing"); playing != "" {
				usd.Activities[0].Name = playing
				usd.Activities[0].Type = discordgo.ActivityTypeGame
			} else if listening := config.GetString("listening"); listening != "" {
				usd.Activities[0].Name = listening
				usd.Activities[0].Type = discordgo.ActivityTypeListening
			} else if watching := config.GetString("watching"); watching != "" {
				usd.Activities[0].Name = watching
				usd.Activities[0].Type = discordgo.ActivityTypeWatching
			} else if streaming := config.GetString("streaming"); streaming != "" {
				usd.Activities[0].Details = streaming
				usd.Activities[0].Type = discordgo.ActivityTypeStreaming
				usd.Activities[0].URL = "https://twitch.tv/btreelar"
			} else {
				usd.Activities[0].Name = "Aqours"
				usd.Activities[0].Type = discordgo.ActivityTypeListening
			}
			bot.Session.UpdateStatusComplex(usd)
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
