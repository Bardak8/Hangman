package main

import (
	"bufio"
	"fmt"
	"math/rand"
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

func Readword1() string {
	var word string
	file, _ := os.Open("words.txt")
	scanner := bufio.NewScanner(file)
	rand.Seed(time.Now().UnixNano())
	(rand.Intn(30))
	for scanner.Text() {
		fmt.println("azejazoejzpajejpzeporjzerozreozro")
	}

}
func Readword2() {

	file, _ := os.Open("words2.txt")
	scanner := bufio.NewScanner(file)
}
func Readword3() {

	file, _ := os.Open("words3.txt")
	scanner := bufio.NewScanner(file)
}

func wordToUnderscore(motChoisi string) string {
	str2 := strings.ReplaceAll(motChoisi, "[a-zA-Z]", "_") // remplace chaque lettre "_"
	return str2
}

func replace(hiddenword string, motChoisi string) string {
	INletter := bufio.NewScanner(os.Stdin)
	INletter.Scan()
<<<<<<< HEAD
	i := findIndex(hiddenword, motChoisi)
	var res string
	str1 := hiddenword[i]
	str2 := motChoisi[i]
	str3 := []rune(hiddenword)
	for i, lettre := range hiddenword {
		if i >= 0 {
			res == (str3-str1)+str2
		} else {
			nil
		}
=======
	{

>>>>>>> d02e62277dc72f5bebda56994116419a379ccd5f
	}
	return res
}

func findIndex(hiddenword string, motChoisi string) int {
<<<<<<< HEAD
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
=======

	var lettre bool
	for i, letter := range motChoisi {
		lettre = Contains(motChoisi, rune(letter))
		if lettre == true {
			return i
		}
	}
	return 0
}

func baba() {

}
>>>>>>> d02e62277dc72f5bebda56994116419a379ccd5f
