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
var guessedletter string = ""

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
	SlowPrint("Quelle bibliothèque de mot souhaitez vous choisir ? \n")
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
		if testmot() || !Contains(wordhidden, '_') {
			displayWinMessage()
			Retry()
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
			println(line)
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
			fmt.Println("Il vous reste", deathCount, "essais")
			return wordhidden
			// mettre à jour le score
		}
		if deathCount == 1 {
			deathCount--
			deathCountStage()
			Retry()
		}
	} else {
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
	lettreoumot = strings.ToLower(lettreoumot)
	// peret à l'utilisateur de savoir qu'il ne doit mettre que des lettres contenues dans l'alphabet latin
	isALetter, err := regexp.MatchString("^[a-zA-Z]", lettreoumot)
	if lettreoumot == guessedletter {
		fmt.Println("vous avez utilisé les lettres :", guessedletter)
		fmt.Println("vous avez deja rentré cette lettre")
	} else {
		if err != nil {
			fmt.Printf("Malheureusement cela ne marche pas ")
			fmt.Printf("Partir %v", lettreoumot)
			return testmot()
		}
		if !isALetter {
			fmt.Printf("Ce n'est pas une lettre !\n")
			return testmot()
		}
		if len(lettreoumot) == 1 {
			guessedletter += lettreoumot
			fmt.Println("vous avez utilisé les lettres :", guessedletter)
			findAndReplace(lettreoumot)
		} else if lettreoumot == word {
			return true
		} else if (len(lettreoumot) == len(word)) && wordhidden == word {
			return true
		} else {
			fmt.Println("Vous n'avez pas trouvé le bon mot")
			deathCount -= 2
			deathCountStage()
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

// Compte le nombre de tour
func countPrint() {
	if count == 1 {
		fmt.Println("------------", count, "er tour", "-------------")
	}
	if count > 1 {
		fmt.Println("------------", count, "ème tour", "-------------")
	}
}

func Retry() {
	count = 0
	deathCount = 10
	displayLoseMessage()
	SlowPrint("Voulez vous recommencer? \n")
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
	if !Contains(o, '1') || !Contains(o, '2') {
		fmt.Println("Veuillez saisir une des réponses proposées")
		Retry()
	}
}

func displayWinMessage() {
	fmt.Println()
	fmt.Println("Tu as découvert le bon mot en ", count, " essai")
	fmt.Println("Votre mot était: ", word)
	fmt.Println("Bravo, vous avez sauvé le pendu")
}

func displayLoseMessage() {
	fmt.Println()
	fmt.Println("Raté ! Tu n'as pas réussi à découvrir le mot")
	fmt.Println("Votre mot choisi était : ", word)
	fmt.Println("Vous essaierez de sauver le pendu une autre fois")
}
