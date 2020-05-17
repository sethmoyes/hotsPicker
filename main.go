package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"regexp"
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

	if len(text) == 1 {
		s.ChannelMessageSend(m.ChannelID, "Be sure to include some parameters. Start with '/hots help' if you are unsure what to do.")
		return
	}

	if isHeroName(text[1]) {
		s.ChannelMessageSend(m.ChannelID, stateHeroInfo(text[1]))
		return
	}

	switch text[1] {
	case "help":
		s.ChannelMessageSend(m.ChannelID, getHelpDialogue())
	case "us":
		if !isHeroName(text[2]) {
			s.ChannelMessageSend(m.ChannelID, "Please use a valid hero name")
		}
		heroName := text[2]
		s.ChannelMessageSend(m.ChannelID, stateHeroInfo(heroName))
	case "team":
		var heroes []string
		for i := 2; i < len(text); i++ {
			heroes = append(heroes, text[i])
		}
		s.ChannelMessageSend(m.ChannelID, analyzeTeamComp(heroes))
	default:
		s.ChannelMessageSend(m.ChannelID, "I didn't catch that. Please try '/hots help' for more info")
	}

}

func analyzeTeamComp(candidates []string) string {
	var allHeroes []models.Hero

	for _, candidate := range candidates {
		hero, e := findHeroFromCandidate(candidate)
		if e != nil {
			return fmt.Sprintf("I don't know who %s is...", candidate)
		}
		allHeroes = append(allHeroes, hero)
	}

	var compCount models.CompCount
	fmt.Println(allHeroes)
	for _, hero := range allHeroes {
		for _, heroComp := range hero.Comp {
			switch heroComp {
			case 0:
				compCount.TeamFightCount++
			case 1:
				compCount.PickCount++
			case 2:
				compCount.SplitPushCount++
			case 3:
				compCount.DiveCount++
			case 4:
				compCount.SafeCount++
			case 5:
				compCount.PokeCount++
			}
		}
	}

	var compCountName = make(models.CompCountName, 6)

	compCountName[0].Name = "Team Fight"
	compCountName[0].Count = compCount.TeamFightCount
	compCountName[1].Name = "Pick"
	compCountName[1].Count = compCount.PickCount
	compCountName[2].Name = "Split Push"
	compCountName[2].Count = compCount.SplitPushCount
	compCountName[3].Name = "Dive"
	compCountName[3].Count = compCount.DiveCount
	compCountName[4].Name = "Safe"
	compCountName[4].Count = compCount.SafeCount
	compCountName[5].Name = "Poke"
	compCountName[5].Count = compCount.PokeCount

	fmt.Println(compCountName)

	var greatest string
	var n int
	//DOES NOT CONSIDER TIED VALUES
	for _, v := range compCountName {
		if v.Count > n {
			n = v.Count
			greatest = v.Name
		}
	}

	return greatest
}

func stateHeroInfo(candidate string) string {
	hero, e := findHeroFromCandidate(candidate)
	if e != nil {
		return "hero not found"
	}
	comps := generateCompList(hero.Comp)
	return fmt.Sprintf("%s is a %s that is good in %+v comps", strings.Title(hero.Name), hero.Role, comps)
}

func getCompName(candidate models.Comp) string {

	switch candidate {
	case 0:
		return "team fight"
	case 1:
		return "pick"
	case 2:
		return "split push"
	case 3:
		return "dive"
	case 4:
		return "safe"
	case 5:
		return "poke"
	}

	return ""
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
			list += "split push"
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
	return "This is the help menu\nUse the command '/hots *hero-name*' to query a hero\nUse '/hots team *hero-name*' with up to five hero names to see what type of comp you are going and who would make a good addition to your team"
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
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		return models.Hero{}, err
	}
	processedString := reg.ReplaceAllString(candidate, "")
	for _, hero := range models.AllHeroes {
		processedHeroName := reg.ReplaceAllString(hero.Name, "")
		if strings.ToLower(processedHeroName) == strings.ToLower(processedString) {
			return hero, nil
		}
	}

	return models.Hero{}, errors.New("No hero found")
}
