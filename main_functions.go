package hangmanpackage

import (
	"log"
	"math/rand"
)

// HangManData is a structure that stores :
// the Word we are currently finding,
// the word we have ToFind,
// the Number of Attempts left,
// and an array of 10 strings with the display of the HangMan

type Data struct {
	Word       string
	ToFind     string
	Attempts   int
	UsedInputs []string
	Difficulty string
	Words      []string
	UserName   string
}

// Return the number of Uderscores in a string, if there is none it mean the word was found
func UderNbr(s string) int {
	word := []rune(s)
	c := 0
	for i := 0; i < len(word); i++ {
		if word[i] == '_' {
			c++
		}
	}
	return c
}

// Reveal the letter every time it is in the word and return the number of time it found the letter
func RevealLetters(Data *Data, input string) int {
	found := 0
	WordRune := []rune(Data.Word)
	for i := 0; i < len(Data.ToFind); i++ {
		if input == string(Data.ToFind[i]) {
			WordRune[i] = rune(Data.ToFind[i])
			found++
		}
	}
	if found > 0 {
		Data.Word = string(WordRune)
	}
	return found
}

// Reveal n letters of the word as starting letters
func RevealStartingLetters(pData *Data) {
	n := len([]rune(pData.ToFind))/2 - 1
	letter := 0
	for letter < n && n != 0 {
		randpos := rand.Intn(len(pData.ToFind))
		if pData.Word[randpos] == '_' {
			found := RevealLetters(pData, string(pData.ToFind[randpos]))
			letter += found
		}
	}
}

// Function that test the error and exit the program if needed
func TestErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Test if the letter is already used
func Used(Data *Data, letter string) bool {
	for i := 0; i < len(Data.UsedInputs); i++ {
		if letter == Data.UsedInputs[i] {
			return true
		}
	}
	return false
}
