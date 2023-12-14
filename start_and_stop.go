package hangmanpackage

import (
	"encoding/json"
	"os"
)

// Stop and Save the game
func StopGame(HMD HangManData) error {
	savedData, err := json.Marshal(HMD)
	TestErr(err)

	file, err := os.Create("save.txt")
	TestErr(err)

	defer file.Close()

	_, err = file.Write(savedData)
	TestErr(err)

	return nil
}

// Overwrite an HangManData variable with the saved data in the file
func ContinueGame(HMD *HangManData, savedDataFile string) {
	encriptedData, err := os.ReadFile(savedDataFile)
	TestErr(err)
	json.Unmarshal(encriptedData, HMD)
}
