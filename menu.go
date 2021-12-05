package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	// "math/rand"
	"os"
	"time"
)

var deathCount int = 10
var count int = 0
var wordhidden string
var word string

func Clear() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}

func SlowPrint(str ...string) {
	for _, strpart := range str {
		for _, char := range strpart {
			fmt.Print(string(char))
			time.Sleep(40_000_000 * time.Nanosecond)
		}
	}
}

func start() {
	SlowPrint("Bonjour \n")
	fmt.Println("1 = Démarrer l'éxécution")
	fmt.Println("2 = Non, je ne souhaite tuer personne")
	// créer une var scanner qui va lire ce que l'utilisateur va écrire
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan() // l'utilisateur input dans la console

	// lis ce que l'utilisateur a écrit
	o := scanner.Text()
	switch o {
	case "1":
		Clear()
		debut()
	case "2":
		os.Exit(2)
	}
}

func debut() {
	SlowPrint("Vous ")
	fmt.Println("1 = Choisir la premiere version")
	fmt.Println("2 = Choisir la deuxieme version")
	fmt.Println("3 = Choisir la troisieme version")
	// créer une var scanner qui va lire ce que l'utilisateur va écrire
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan() // l'utilisateur input dans la console

	// lis ce que l'utilisateur a écrit
	o := scanner.Text()
	switch o {
	case "1":
		startGame("words.txt", 37)
	case "2":
		startGame("words2.txt", 23)
	case "3":
		startGame("words3.txt", 24)
	}
}

func startGame(filename string, nbword int) {
	word = Readword(filename, nbword)
	wordhidden = wordToUnderscore()
	fmt.Println(wordhidden)
	for {
		if testmot() {
			fmt.Println("vous avez gagné")
			break
		}
	}
	// trouve le mot et transforme le mot choisi en underscore
}
func Readword(filename string, nbword int) string {
	rand.Seed(time.Now().UnixNano())
	randnumber := (rand.Intn(nbword))
	// Open the file.
	f, _ := os.Open(filename)
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	index := 1
	for scanner.Scan() {
		line := scanner.Text()
		index++
		if index == randnumber {
			fmt.Println(line)
			return line
		}
	}
	return ""
}

func wordToUnderscore() string {
	sampleRegexp := regexp.MustCompile("[a-z,A-Z]")

	input := word

	result := sampleRegexp.ReplaceAllString(input, "_")
	return (string(result))
}

func findAndReplace(letterToReplace string) string {
	isFound := strings.Index(word, letterToReplace)
	if isFound == -1 {
		if deathCount > 1 {
			deathCount--
			deathCountStage()
			fmt.Println("raté")
			fmt.Println("ils vous restent", deathCount, "essaies")
			return wordhidden
			// mettre à jour le score
		}
		if deathCount == 1 {
			deathCount--
			deathCountStage()
			Retry()
		}
	} else {
		// var res string
		// str1 := hiddenword[i]
		// str2 := motChoisi[i]
		str3 := []rune(wordhidden)
		for i, lettre := range word {
			if string(lettre) == letterToReplace {
				str3[i] = lettre
				wordhidden = string(str3)
				fmt.Println(wordhidden)
			}
		}
	}
	return wordhidden
}

func testmot() bool {
	count++
	countPrint()
	fmt.Println("Veuillez saisir une lettre ou un mot")
	// créer une var scanner qui va lire ce que l'utilisateur va écrire
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan() // l'utilisateur input dans la console

	// lis ce que l'utilisateur a écrit
	println(wordhidden)
	lettreoumot := scanner.Text()
	if len(lettreoumot) == 1 {
		findAndReplace(lettreoumot)
	} else {
		if lettreoumot == word {
			return true
		}
	}
	return false
}

func deathCountStage() {
	if deathCount == 9 {
		fmt.Printf("      \n")
		fmt.Printf("       \n")
		fmt.Printf("       \n")
		fmt.Printf("       \n")
		fmt.Printf("       \n")
		fmt.Printf("       \n")
		fmt.Printf("       \n")
		fmt.Printf("       \n")
		fmt.Printf("========\n")
	}
	if deathCount == 8 {
		fmt.Printf("      \n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	}
	if deathCount == 7 {
		fmt.Printf("  +---+\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	}
	if deathCount == 6 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	}
	if deathCount == 5 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	}
	if deathCount == 4 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	}
	if deathCount == 3 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	}
	if deathCount == 2 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|/  |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	}
	if deathCount == 1 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|/ |\n")
		fmt.Printf(" /   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	}
	if deathCount == 0 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|/  |\n")
		fmt.Printf(" / /  |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	}
}

func countPrint() {
	if count == 1 {
		fmt.Println("------------", count, "er tour", "-------------")
	}
	if count > 1 {
		fmt.Println("-------------", count, "ème tour", "-------------")
	}
}

func Retry() {
	count = 0
	deathCount = 10
	SlowPrint("Vous avez perdu \n")
	SlowPrint("Voulez vous recommencez? \n")
	fmt.Println("1 = Oui")
	fmt.Println("2 = Non")
	// créer une var scanner qui va lire ce que l'utilisateur va écrire
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan() // l'utilisateur input dans la console

	// lis ce que l'utilisateur a écrit
	o := scanner.Text()
	switch o {
	case "1":
		Clear()
		debut()
	case "2":
		os.Exit(2)
	}
}
