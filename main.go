package main

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

const NUM_PLAYS = 6

var wordsList = []string{"about", "above", "admit", "adult", "after", "again", "agent", "agree", "ahead", "allow", "alone",
	"along", "among", "apply", "argue", "avoid", "begin", "black", "blood", "board", "break", "bring",
	"build", "carry", "catch", "cause", "chair", "check", "child", "civil", "claim", "class", "clear",
	"close", "coach", "color", "could", "court", "cover", "crime", "death", "dream", "drive", "early",
	"eight", "enjoy", "enter", "event", "every", "exist", "field", "fight", "final", "first", "floor",
	"focus", "force", "front", "glass", "great", "green", "group", "guess", "happy", "heart", "heavy",
	"hotel", "house", "human", "image", "issue", "large", "later", "laugh", "learn", "least", "leave",
	"legal", "level", "light", "local", "major", "maybe", "media", "might", "model", "money", "month",
	"mouth", "movie", "music", "never", "night", "north", "occur", "offer", "often", "order", "other",
	"owner", "paper", "party", "peace", "phone", "piece", "place", "plant", "point", "power", "price",
	"prove", "quilt", "quite", "radio", "raise", "range", "reach", "ready", "right", "scene", "score",
	"sense", "serve", "seven", "shake", "share", "shoot", "short", "since", "skill", "small", "smile",
	"sound", "south", "space", "speak", "spend", "sport", "staff", "stage", "stand", "start", "state",
	"still", "stock", "store", "story", "study", "stuff", "style", "table", "teach", "thank", "their",
	"there", "these", "thing", "think", "third", "those", "three", "throw", "today", "total", "tough",
	"trade", "treat", "trial", "truth", "under", "until", "value", "visit", "voice", "watch", "water",
	"where", "which", "while", "white", "whole", "whose", "woman", "world", "worry", "would", "write",
	"wrong", "young"}

const RESET = "\033[0m"

type Colors struct {
	code string
}

func (c *Colors) Println(message string) (int, error) {
	return fmt.Printf("%s%s%s\n", c.code, message, RESET)
}

func (c *Colors) Fprintln(w io.Writer, message string) (int, error) {
	return fmt.Fprintf(w, "%s%s%s", c.code, message, RESET)
}

func (c *Colors) Sprintf(message string) string {
	return fmt.Sprintf("%s%s%s", c.code, message, RESET)
}

var Red = Colors{code: "\033[;31m"}
var Green = Colors{code: "\033[;32m"}
var Yellow = Colors{code: "\033[;33m"}
var White = Colors{code: "\033[;37m"}

var RedBackground = Colors{code: "\033[7;31m"}
var GreenBackground = Colors{code: "\033[7;32m"}
var YellowBackground = Colors{code: "\033[7;33m"}
var WhiteBackground = Colors{code: "\033[7;37m"}

func isGuessCorrect(theGuess, theWord string) bool {
	if len(theGuess) != 5 {
		Red.Println("That wasn't five characters, try again")
		return false
	}
	correct := true
	var response string = ""
	for index, chr := range theGuess {
		if chr == int32(theWord[index]) {
			response += Green.Sprintf(string(chr))
		} else {
			found := false
			correct = false
			for _, chr2 := range theWord {
				if chr == chr2 {
					response += Yellow.Sprintf(string(chr))
					found = true
					break
				}
			}
			if !found {
				response += string(chr)
			}
		}
	}
	fmt.Println(response)
	return correct
}

func main() {
	fmt.Print("\033[H\033[2J")
	rand.Seed(time.Now().Unix())
	randIndex := rand.Intn(len(wordsList))
	theWord := strings.ToLower(wordsList[randIndex])
	fmt.Printf("Picked a word from list of %d words.\n", len(wordsList))

	won := false
	for play := 0; play < NUM_PLAYS; play++ {
		var guess string
		fmt.Scanln(&guess)
		if isGuessCorrect(guess, theWord) {
			won = true
			break
		}
	}

	if won {
		Green.Println("\nYou won!!")
	} else {
		Red.Println(fmt.Sprintf("\nThe word was %s\nBetter luck next time.", theWord))
	}
}
