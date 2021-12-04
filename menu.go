package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"regexp"

	// "math/rand"
	"os"
	"time"
)

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
	wordfound := Readword(filename, nbword)
	wordfoundunderscore := wordToUnderscore(wordfound)
	fmt.Println(wordfoundunderscore)
	for {
		if testmot(wordfound) {
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

func wordToUnderscore(motChoisi string) string {
	sampleRegexp := regexp.MustCompile("[a-z,A-Z]")

	input := motChoisi

	result := sampleRegexp.ReplaceAllString(input, "_")
	return (string(result))
}

func replace(hiddenword string, motChoisi string) string {
	INletter := bufio.NewScanner(os.Stdin)
	INletter.Scan()
	i := findIndex(hiddenword, motChoisi)
	//var res string
	//str1 := hiddenword[i]
	//str2 := motChoisi[i]
	//str3 := []rune(hiddenword)
	//for i, lettre := range hiddenword {
	if i >= 0 {
		//		res == (str3-str1)+str2
	} else {
		//		nil
	}
	//}
	return ""
}

func findIndex(hiddenword string, motChoisi string) int {
	var lettre bool
	for i, letter := range motChoisi {
		lettre = Contains(motChoisi, rune(letter))
		if lettre == false {
			fmt.Print("raté")
		}
		return i
	}
	return -1
}

func testmot(wordfound string) bool {

	fmt.Println("Veuillez saisir une lettre ou un mot")
	// créer une var scanner qui va lire ce que l'utilisateur va écrire
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan() // l'utilisateur input dans la console

	// lis ce que l'utilisateur a écrit
	fmt.Println(wordfound)
	lettreoumot := scanner.Text()
	if len(lettreoumot) == 1 {
		testlettre()
	} else {
		if lettreoumot == wordfound {
			return true
		}
	}
	return false
}

func testlettre() bool {
	return false
}
