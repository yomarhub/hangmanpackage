package hangmanpackage

import (
	"log"
)

// HangManData is a structure that stores :
// the Word we are currently finding,
// the word we have ToFind,
// the Number of Attempts left,
// and an array of 10 strings with the display of the HangMan
type HangManData struct {
	Word             []rune       // Word composed of '_', ex: H_ll_
	ToFind           string       // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int          // Number of attempts left
	UsedLetters      []rune       // Letters that were already used
	HangmanPositions [10]string   // It can be the array where the positions parsed in "hangman.txt" are stored
	Font             [95][]string // Font Used for Ascii Art
}

// Return the number of Uderscores in an rune array, if there is none it mean the word was found
func UderNbr(word []rune) int {
	c := 0
	for i := 0; i < len(word); i++ {
		if word[i] == '_' {
			c++
		}
	}
	return c
}

// Reveal the letter every time it is in the word and return the number of time it found the letter
func RevealLetters(HMD HangManData, l string) int {
	found := 0
	for i := 0; i < len(HMD.ToFind); i++ {
		if l == string(HMD.ToFind[i]) {
			HMD.Word[i] = rune(HMD.ToFind[i])
			found++
		}
	}
	return found
}

// Function that test the error and exit the program if needed
func TestErr(err error) {
	if err != nil {
		log.Fatal(err)
	}

}

// Test if the letter is already used
func Used(HMD HangManData, letter rune) bool {
	for i := 0; i < len(HMD.UsedLetters); i++ {
		if letter == HMD.UsedLetters[i] {
			return true
		}
	}
	return false
}
