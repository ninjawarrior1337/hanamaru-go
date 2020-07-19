package main

import (
	"flag"
	"fmt"
	_ "github.com/markbates/pkger"
	"github.com/ninjawarrior1337/hanamaru-go/events"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"syscall"
)

var config *viper.Viper

var commands []*framework.Command
var optionalEvents []*framework.EventListener

func init() {
	config = viper.New()
	config.AddConfigPath(".")
	config.AllSettings()
	config.AutomaticEnv()

	config.SetDefault("owner", "")
	config.SetDefault("prefix", "!")
	config.SetDefault("token", "")
	config.SetDefault("listening", "Aqours' Songs")
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

	if playing := config.GetString("playing"); playing != "" {
		bot.Session.UpdateListeningStatus(playing)
	} else if listening := config.GetString("listening"); listening != "" {
		bot.Session.UpdateListeningStatus(listening)
	} else {
		bot.Session.UpdateListeningStatus("Aqours' Songs")
	}

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
