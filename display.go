package hangman

import "fmt"

func DisplayWord(HMD HangManData) {
	wordarr := HMD.Word
	for i := 0; i < len(wordarr); i++ {
		fmt.Printf("%c", wordarr[i])
		if i < len(wordarr)-1 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}
}

func PrintResult(HMD HangManData) {
	if HMD.Attempts == 0 && UderNbr(HMD.Word) > 0 {
		fmt.Println("=====================")
		fmt.Println("    You Have Lost")
		fmt.Printf("=====================\n\n")
		fmt.Println("The Word Was :", HMD.ToFind)
	} else {
		DisplayWord(HMD)
		fmt.Println("\nCongrats !")
	}
}

func FontCreate(fontFile string) [95][]string {
	// font := make([][]string, 95, 95)
	font := [95][]string{}

	lines, err := FileToLines(fontFile)
	TestErr(err)

	start, stop := 0, 9
	for letter := range font {
		font[letter] = lines[start:stop]
		start += 9
		stop += 9
	}

	return font
}
