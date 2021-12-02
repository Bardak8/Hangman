package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	SlowPrint("Bonjour ")
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
		Readword1()
	case "2":
		Readword2()
	case "3":
		Readword3()
	}
}

func Readword1() {

	content, err := ioutil.ReadFile("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
func Readword2() {

	content, err := ioutil.ReadFile("words2.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
func Readword3() {

	content, err := ioutil.ReadFile("words3.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}

func wordToUnderscore(motChoisi string) string {
	str2 := strings.ReplaceAll(motChoisi, "[a-zA-Z]", "_") // remplace chaque lettre "_"
	return str2
}

func Guesstheword(hiddenword string, motChoisi string) string{
	INletter := bufio.NewScanner(os.Stdin)
	INletter.Scan()
	letter := Capitalize(INletter.Text())	
	hiddenword[i] = motChoisi[i]
	for_, INletter := range motChoisi {
		if Contains(motChoisi, INletter) {
			strings.Replace()
		} else {
			return true
		}
	}

}
