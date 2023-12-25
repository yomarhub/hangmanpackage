package hangmanpackage

import (
	"bufio"
	"os"
)

// Take the file and convert it to a string array
func FileToLines(file string) ([]string, error) {
	readFile, err := os.Open(file)
	if err != nil {
		return []string{}, err
	}
	defer readFile.Close()
	var wordsarr []string

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		wordsarr = append(wordsarr, fileScanner.Text())
	}
	return wordsarr, err
}
