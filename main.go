package main

import (
	"fmt"
	"strings"

	"github.com/gempir/go-twitch-irc"
)

func main() {
	client := twitch.NewClient("tarot_bot", "oauth:s9xlwc9re3o761oh505glku44c5y6q")

	client.OnNewMessage(func(channel string, user twitch.User, message twitch.Message) {
		if strings.HasPrefix(message.Text, "!tarot") {
			reply := tarotHandler(user, message)
			client.Say(channel, reply)
		}
	})

	client.OnNewWhisper(func(user twitch.User, message twitch.Message) {
		if strings.HasPrefix(message.Text, "!tarot") {
			reply := tarotHandler(user, message)
			fmt.Println(user.Username, reply)
			client.Whisper(user.Username, reply)
		}
	})

	client.Join("jellygator")

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func tarotHandler(user twitch.User, message twitch.Message) string {
	//Users to tag when replying
	usernames := []string{user.DisplayName}
	//Slice of cards to return information on
	var cards = make(map[card]bool)
	//Is the user requesting the Biddy Tarot link?
	isLinkRequest := false
	//Is the user requesting help with the syntax?
	isHelpRequest := false
	//Is the user requesting info on the creator?
	isCreditsRequest := false
	//Slice of terms that tarot_bot doesn't recognize
	var wrongTerms []string

	//If no terms supplied, return help info
	if message.Text == "!tarot" {
		isHelpRequest = true
	}

	//Trim off the command prefix "!tarot"
	msg := strings.TrimPrefix(message.Text, "!tarot")
	if strings.HasPrefix(msg, " ") {
		msg = strings.TrimPrefix(msg, " ")
	}
	//Create a slice of individual terms from the command message
	terms := strings.Split(msg, " ")

	//Loop through all terms
	for _, term := range terms {
		//Test if user is tagged
		if strings.HasPrefix(term, "@") {
			username := strings.TrimPrefix(term, "@")
			//If usernames contains the message sender, overwrite it, otherwise append
			if contains(usernames, user.DisplayName) {
				usernames = []string{username}
			} else {
				usernames = append(usernames, username)
			}
			continue
		}
		//Test if requesting a link
		if term == "link" {
			isLinkRequest = true
			continue
		}
		//Test if requesting help info
		if term == "help" {
			isHelpRequest = true
			continue
		}
		//Test if requesting credits
		if term == "credits" {
			isCreditsRequest = true
			continue
		}
		//Attempt to get a card by its abbreviation
		card, inverted, err := getCardByAbbreviation(term)
		//If an error is returned, add term to array for use in response
		if err != nil {
			wrongTerms = append(wrongTerms, term)
			continue
		}
		if len(cards) < 3 {
			cards[card] = inverted
		}
	}

	//Start building the message
	var reply string

	//Tag users, e.g. "@Rhythmatic_FM @Jellygator --> "
	for i, name := range usernames {
		if i > 0 {
			reply += " "
		}
		reply += "@" + name
	}
	reply += " --> "

	if isHelpRequest {
		reply += "Usage Example: For Ace of Pentacles, The Stars, and Queen of Cups Reversed, type '!tarot AP XVII QCr'. To get a link: '!tarot AP link'. Mentioning @users works. '!tarot credits' for info on the bot's creator."
	}

	if isCreditsRequest && !isHelpRequest {
		reply += "Created by Nathan Young aka @Rhythmatic_FM. Projects and contact info can be found at http://funnykeyboardsymbols.com"
	}

	if !isHelpRequest && !isCreditsRequest {
		loops := 0
		for card, inverted := range cards {
			loops++

			if isLinkRequest {
				reply += card.name + ": " + card.link
				break
			} else {
				switch loops {
				case 1:
					reply += "1️⃣ "
				case 2:
					reply += "2️⃣ "
				case 3:
					reply += "3️⃣ "
				default:
					break
				}

				if inverted {
					reply += card.name + " (inverted): " + card.inverted
				} else {
					reply += card.name + ": " + card.upright
				}
				reply += " "
			}
		}

		if len(wrongTerms) > 0 {
			reply += " ❓ I didn't understand these terms: "
			for i, wrongTerm := range wrongTerms {
				if i > 0 {
					reply += " "
				}
				reply += wrongTerm
			}
		}
	}
	return reply
}
