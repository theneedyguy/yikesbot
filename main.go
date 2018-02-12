package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/voloshink/dggchat"
)

const (
	defaultConfigFile = "config.json"
	pingInterval      = time.Minute
)

var (
	yikesCommands      = []string{"!yikes", "!YIKES", "! yikes", "! YIKES", "!yikers", "!YIKERS"}
	lastMessage        = ""
	yikesVersion       = "1.5.1"
	yikesLevel         = 0
	yikesMessage       = 0
	ipbanMessage       = 0
	yikesSleep         = false
	yikesTop           = 0
	lastOmegaYikes     = time.Now()
	omegaYikesinterval = time.Second * 60
	lastSent           = time.Now()
	lastPM             = time.Now()
	startTime          = time.Now()
	lastPing           = timeToUnix(time.Now())
	lastPong           = timeToUnix(time.Now())
	messageInterval    = time.Second * 10
	pmInterval         = time.Second * 30
	configFile         string
	admins             []string
	key                string
)

type (
	config struct {
		Key    string   `json:"login_key"`
		Admins []string `json:"admins"`
	}

	apiResp struct {
		URL string `json:"url"`
	}
)

func main() {
	var file string
	if len(os.Args) < 2 {
		file = defaultConfigFile
	} else {
		file = os.Args[1]
	}

	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}

	var c config
	err = json.Unmarshal(f, &c)
	if err != nil {
		log.Fatalln(err)
	}

	configFile = file

	if c.Key == "" {
		log.Fatalln("No login key provided")
	}

	key = c.Key

	if c.Admins == nil {
		c.Admins = make([]string, 0)
	}

	admins = c.Admins

	go startBot(c.Key)
	go decreaseYikesOverTime()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT)
	<-sc
}

func startBot(key string) {
	dgg, err := dggchat.New(key)
	if err != nil {
		log.Fatalln(err)
	}

	err = dgg.Open()
	if err != nil {
		log.Fatalln(err)
	}

	defer dgg.Close()

	messages := make(chan dggchat.Message)
	errors := make(chan string)
	pings := make(chan dggchat.Ping)

	dgg.AddMessageHandler(func(m dggchat.Message, s *dggchat.Session) {
		messages <- m
	})

	dgg.AddErrorHandler(func(e string, s *dggchat.Session) {
		errors <- e
	})

	dgg.AddPingHandler(func(p dggchat.Ping, s *dggchat.Session) {
		pings <- p
	})

	go checkConnection(dgg)

	for {
		select {
		case m := <-messages:
			if strings.Contains(m.Message, "omegayikes") ||
				strings.Contains(m.Message, "OMEGAYIKES") {
				omegaYikes()
			}
			if strings.Contains(m.Message, "YIKES") ||
				strings.Contains(m.Message, "Y I K E S") ||
				strings.Contains(m.Message, "yikes") ||
				strings.Contains(m.Message, "yikers") ||
				strings.Contains(m.Message, "YIKERS") {
				raiseYikesLevel(10)
			}
			if strings.HasPrefix(m.Message, "!") {
				handleCommand(m, dgg)
			}
		case e := <-errors:
			log.Printf("Error %s\n", e)
		case p := <-pings:
			lastPong = p.Timestamp
		}
	}

}

func checkConnection(s *dggchat.Session) {
	ticker := time.NewTicker(pingInterval)
	for {
		<-ticker.C
		if lastPing != lastPong {
			log.Println("Ping mismatch, attempting to reconnect")
			err := s.Close()
			if err != nil {
				log.Fatalln(err)
			}

			err = s.Open()
			if err != nil {
				log.Fatalln(err)
			}

			continue
		}
		s.SendPing()
		lastPing = timeToUnix(time.Now())
	}
}

func timeToUnix(t time.Time) int64 {
	return t.Unix() * 1000
}
