package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/voloshink/dggchat"
)

func handleCommand(m dggchat.Message, s *dggchat.Session) {

	if isYikesCommand(m.Message) {
		handleYikesCommand(m, s)
		return
	}

	if strings.HasPrefix(m.Message, "!reset") {
		if isAdmin(m.Sender.Nick) {
			resetYikes()
			_ = s.SendMessage("YIKER-COUNTER RESET FeelsBadMan")
		}
	}

	if strings.HasPrefix(m.Message, "!topyikes") {
		if isAdmin(m.Sender.Nick) {
			_ = s.SendMessage(fmt.Sprintf("Highest recorded yikes level (since starting the bot): %s", strconv.Itoa(yikesTop)))
		}
	}

	if strings.HasPrefix(m.Message, "!sleep") {

		//if strings.EqualFold("theneedyguy", m.Sender.Nick) {
		if isAdmin(m.Sender.Nick) {
			isSleeping := toggleSleep()
			if isSleeping == true {
				_ = s.SendMessage("I am going to sleep...")
			} else {
				_ = s.SendMessage("Ready to count yikes again.")
			}

		}
	}

	if strings.HasPrefix(m.Message, "!ipban") {
		if strings.EqualFold("Destiny", m.Sender.Nick) {

			switch ipbanMessage {
			case 0:
				_ = s.SendMessage("Yikes increased by 100 due to ban.")
				raiseYikesLevel(100)
				ipbanMessage = 1
			case 1:
				_ = s.SendMessage("Ban! 100 added to yikes-o-meter.")
				ipbanMessage = 0
			default:
				_ = s.SendMessage("Ban! 100 added to yikes-o-meter.")
				ipbanMessage = 0
			}

		}
	}

	if strings.HasPrefix(m.Message, "!ver") {
		if isAdmin(m.Sender.Nick) {
			_ = s.SendMessage(yikesVersion)
		}
	}
	if strings.HasPrefix(m.Message, "!ping") {
		if isAdmin(m.Sender.Nick) {
			_ = s.SendMessage("YIKESPONG")
		}
	}
}

func isYikesCommand(s string) bool {
	for _, c := range yikesCommands {
		if strings.HasPrefix(s, c) {
			return true
		}
	}
	return false
}

func setTopYikes(yikesLevelInput int) {
	if yikesLevelInput > yikesTop {
		yikesTop = yikesLevelInput
	}
}

func resetYikes() {
	yikesLevel = 0
}

func toggleSleep() bool {
	if yikesSleep == false {
		yikesSleep = true
		return true
	} else {
		yikesSleep = false
		return false
	}
}

func getYikesLevel() (string, error) {
	switch yikesMessage {
	case 0:
		message := []string{"AMOUNT OF YIKES:", strconv.Itoa(yikesLevel)}
		yikesMessage = yikesMessage + 1
		return strings.Join(message, " "), nil
	case 1:
		message := []string{"LEVELS OF YIKES:", strconv.Itoa(yikesLevel)}
		yikesMessage = yikesMessage + 1
		return strings.Join(message, " "), nil
	case 2:
		message := []string{"TOTAL YIKES:", strconv.Itoa(yikesLevel)}
		yikesMessage = yikesMessage - 2
		return strings.Join(message, " "), nil
	default:
		message := []string{"CHAT YIKES:", strconv.Itoa(yikesLevel)}
		yikesMessage = 0
		return strings.Join(message, " "), nil
	}
}

func raiseYikesLevel(amount int) {
	yikesLevel = yikesLevel + amount
	setTopYikes(yikesLevel)
}

func omegaYikes() {
	omegaYikesElapsed := time.Since(lastOmegaYikes)
	if omegaYikesElapsed < omegaYikesinterval {
		return
	}
	raiseYikesLevel(50)
	lastOmegaYikes = time.Now()
}

func decreaseYikesOverTime() {
	for {
		//goroutine that runs indefinitely.
		if yikesLevel != 0 {
			yikesLevel = yikesLevel - 1
		}
		time.Sleep(time.Second * 2)
	}
}

func isAdmin(s string) bool {
	return isInList(s, admins)
}

func isInList(s string, list []string) bool {
	for _, v := range list {
		if strings.EqualFold(s, v) {
			return true
		}
	}
	return false
}

func handleYikesCommand(m dggchat.Message, s *dggchat.Session) {
	timeElapsed := time.Since(lastSent)
	if timeElapsed < messageInterval {
		return
	}

	ylevel, err := getYikesLevel()
	if err != nil {
		return
	}
	/*
		if ylevel == lastMessage {
			return
		}
	*/
	if yikesSleep == false {
		err = s.SendMessage(fmt.Sprintf("%s WhoahDude", ylevel))
		if err != nil {
			return
		}
	} else {
		err = s.SendMessage(fmt.Sprintf("%s I am sleeping until awakened by my master.", "SLEEPSTINY"))
		if err != nil {
			return
		}
	}
	lastSent = time.Now()
	lastMessage = ylevel
}
