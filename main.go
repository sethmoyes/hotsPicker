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

	compCountName[0].Name = models.TeamFight
	compCountName[0].Count = compCount.TeamFightCount
	compCountName[1].Name = models.Pick
	compCountName[1].Count = compCount.PickCount
	compCountName[2].Name = models.SplitPush
	compCountName[2].Count = compCount.SplitPushCount
	compCountName[3].Name = models.Dive
	compCountName[3].Count = compCount.DiveCount
	compCountName[4].Name = models.Safe
	compCountName[4].Count = compCount.SafeCount
	compCountName[5].Name = models.Poke
	compCountName[5].Count = compCount.PokeCount

	var highestPointComps []models.Comp
	var n int
	//DOES NOT CONSIDER TIED VALUES
	for _, v := range compCountName {
		if v.Count > n {
			n = v.Count
			// greatest = v.Name
		}
	}
	for _, v := range compCountName {
		if v.Count == n {
			highestPointComps = append(highestPointComps, v.Name)
		}
	}

	var suggestions string
	if len(allHeroes) > 2 && len(allHeroes) < 5 {
		suggestions = suggestHeroesForComp(allHeroes, highestPointComps)
	}

	comps := generateCompList(highestPointComps)

	var stringToReturn string

	if len(highestPointComps) == 1 {
		stringToReturn = fmt.Sprintf("***Your most viable comp is*** __%s__\n", comps)
	} else {
		stringToReturn = fmt.Sprintf("***Your most viable comps are*** __%s__\n", comps)
	}

	if len(suggestions) > 0 {
		stringToReturn = fmt.Sprintf("%sHere are some suggestions...\n%s", stringToReturn, suggestions)
	}
	return stringToReturn
	// return fmt.Sprintf("Comp Type - Points\n%s - %d\n%s - %d\n%s - %d\n%s - %d\n%s - %d\n%s - %d\n", compCountName[0].Name, compCountName[0].Count, compCountName[1].Name, compCountName[1].Count, compCountName[2].Name, compCountName[2].Count, compCountName[3].Name, compCountName[3].Count, compCountName[4].Name, compCountName[4].Count, compCountName[5].Name, compCountName[5].Count)
}

func suggestHeroesForComp(heroes []models.Hero, highestCompType []models.Comp) string {
	var heroSuggestions string
	var hasHealer, hasTank, hasBrusier, hasMeleeAssassin, hasRangedAssassin bool

	for _, hero := range heroes {
		switch hero.Role {
		case models.Healer:
			hasHealer = true
		case models.Tank:
			hasTank = true
		case models.Bruiser:
			hasBrusier = true
		case models.MeleeAssassin:
			hasMeleeAssassin = true
		case models.RangedAssassin:
			hasRangedAssassin = true
		}
	}

	for _, comp := range highestCompType {
		var suggestionsForComp string
		if !hasHealer {
			suggestionsForComp += concatSuggestions(models.Healer, comp)
		}
		if !hasTank {
			suggestionsForComp += concatSuggestions(models.Tank, comp)
		}
		if !hasBrusier {
			suggestionsForComp += concatSuggestions(models.Bruiser, comp)
		}
		if !hasRangedAssassin {
			suggestionsForComp += concatSuggestions(models.RangedAssassin, comp)
		}
		if !hasMeleeAssassin {
			suggestionsForComp += concatSuggestions(models.MeleeAssassin, comp)
		}
		heroSuggestions += suggestionsForComp + "\n"
	}

	return heroSuggestions
}

func concatSuggestions(role models.Role, comp models.Comp) string {
	var suggestionsForComp string
	suggestions := filterForHeroByRoleAndComp(role, comp)
	var sugg string
	for _, suggestion := range suggestions {
		sugg += suggestion.Name + "/"
	}
	suggestionsForComp += fmt.Sprintf("You need a **%s**. For **%s** try %s\n", string(role), getCompName(comp), strings.Trim(sugg, "/"))
	return suggestionsForComp
}

func filterForHeroByRoleAndComp(role models.Role, comp models.Comp) []models.Hero {
	var suitableHeroes []models.Hero

	for _, hero := range models.AllHeroes {
		if hero.Role == role && comp.IsInList(hero.Comp) {
			suitableHeroes = append(suitableHeroes, hero)
		}
	}

	return suitableHeroes
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
	return "This is the help menu\nUse the command '/hots *hero-name*' to query a hero\nUse '/hots team *hero-name*' with up to five hero names to see what type of comp you are going and who would make a good addition to your team\nNeed suggestions? Use '/hots team *hero-name* (include three or four heroes this way) too see what kind of comp you are running and who would be a nice addition to your team.\n"
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
