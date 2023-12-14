package hangmanpackage

import (
	"bufio"
	"log"
	"os"
)

// Take the file and convert it to a string array
func FileToArr(file string) ([]string, error) {
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

// Take a file and return an array of each line
func FileToLines(fontFile string) ([]string, error) {
	file, err := os.Open(fontFile)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines, nil
}

// Convert the file to an array of 10 strings that have the display of the HangMan
func HangManDisplay(filename string) ([10]string, error) {
	positions := [10]string{}
	file, err := os.ReadFile(filename)
	if err != nil {
		return positions, err
	}
	start, finish := 0, 71
	for i := 0; i < 10; i++ {
		positions[i] = string(file[start:finish])
		start += 71
		finish += 71
	}
	return positions, err
}
