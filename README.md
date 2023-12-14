# Hangman

This project is based on the famous HangMan Game.
Hangman is a word-guessing game usually played with two people. One person thinks of a word but keeps it a secret. They then draw a set of dashes on a piece of paper, each dash representing a letter in the word. The other person tries to guess the word by suggesting letters.

For each correct letter guessed, the person reveals its position in the word by filling in the corresponding dash. However, if a guessed letter is not in the word, a part of a stick figure (representing a "hangman") is drawn. The goal for the person guessing the word is to figure out the complete word before the hangman is fully drawn.

The game typically has a limit on the number of incorrect guesses allowed. If the person guessing the word reaches that limit before figuring out the word, the hangman is fully drawn, and they lose. If they guess the word before that happens, they win!

It's a fun and simple game that involves a bit of guessing and word knowledge.

This program play the role of the other person to play with you, you start with 10 attemps and can display the word using ascii art from the asset folder or create your own using standard.txt as a model (every character must use 10 lines) you can also create your own dictionary (no spaces in words, one word per line, no special characters), you can suggest either a word or a letter, if the letter is not in the word you lose an attempt and if a word is false you lose 2 attemps.

## Index

 - [Index](#index)
 - [Prerequisites](#prerequisites)
 - [Installation](#installation)
 - [Authors](#authors)

## Prerequisites

To install and run the project, you will need :

- [Go](https://go.dev/doc/install) installed on your local machine
- [Git](https://git-scm.com/downloads) installed on your local machine

## Installation

0. Ask for access from us to be added as a collaborator

1. Clone the Repository using git :
```bash
  git clone https://ytrack.learn.ynov.com/git/omary/Hangman.git
  cd Hangman/
```

2. Download the necessary packages with : 
```bash
  go mod tidy
```

## Play the game

1. To Start a game you have to start the game with a dictionary file to randomly choose a word or start a stopped game :  
    >No option Game : The game start using `go run main.go path/to/words/file`  
    >Ascii Art : It write the word in Ascii Art run using `go run main.go path/to/words/file --letterFile path/to/ascii/file`  
    >Resume a previously stopped game : `go run main.go --startWith save.txt`  

2. To Stop in the middle of a game Write "STOP" or "stop", to resume it later start a game with the command `go run main.go --startWith save.txt`


## Authors

- [@yomar](https://ytrack.learn.ynov.com/git/yomar)# hangman
