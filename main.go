package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	models "github.com/ColeJSmith19/hotsPicker/models"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	text := strings.Split(m.Content, " ")

	if text[0] != "/hots" {
		return
	}

	if isHeroName(text[1]) {
		s.ChannelMessageSend(m.ChannelID, stateHeroInfo(text[1]))
	} else {

		switch text[1] {
		case "help":
			s.ChannelMessageSend(m.ChannelID, getHelpDialogue())
		case "us":
			if !isHeroName(text[2]) {
				s.ChannelMessageSend(m.ChannelID, "Please use a valid hero name")
			}
			heroName := text[2]
			s.ChannelMessageSend(m.ChannelID, stateHeroInfo(heroName))
		default:
			s.ChannelMessageSend(m.ChannelID, "I didn't catch that. Please try '/hots help' for more info")
		}
	}
}

func stateHeroInfo(candidate string) string {
	hero, e := findHeroFromCandidate(candidate)
	if e != nil {
		return "hero not found"
	}
	comps := generateCompList(hero.Comp)
	return fmt.Sprintf("%s is a %s that is good in %+v comps", strings.Title(hero.Name), hero.Role, comps)
}

func generateCompList(comp []models.Comp) string {
	var list string
	for _, x := range comp {
		switch x {
		case 0:
			list += "team fight/"
		case 1:
			list += "pick/"
		case 2:
			list += "split/"
		case 3:
			list += "dive/"
		case 4:
			list += "safe/"
		case 5:
			list += "poke/"
		}
	}

	return strings.Trim(list, "/")
}

func getHelpDialogue() string {
	return "This is the help menu\nUse the command '/hots us *hero-name*' to query a hero"
}

func stateHeroRolls(candidate string) string {
	hero, _ := findHeroFromCandidate(candidate)
	b, _ := json.Marshal(hero.Role)
	return string(b)
}

func isHeroName(candidate string) bool {
	_, e := findHeroFromCandidate(candidate)
	if e != nil {
		return false
	}
	return true
}

func findHeroFromCandidate(candidate string) (models.Hero, error) {
	for _, hero := range models.AllHeroes {
		if hero.Name == strings.ToLower(candidate) {
			return hero, nil
		}
	}

	return models.Hero{}, errors.New("No hero found")
}
