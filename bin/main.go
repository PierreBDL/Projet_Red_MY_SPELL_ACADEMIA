package main

import (
	"MSA"
	"fmt"
	"unicode"

	// Majuscule
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	fmt.Println(" ###### MY SPELL ACADEMIA ######")

	// Création du personnage
	var name string
	fmt.Print("Quel est votre nom ? ")
	fmt.Scan(&name)

	// Vérifier la validité du nom (pas de chiffres ou de caractères spéciaux)
	for _, caracter := range name {
		if !unicode.IsLetter(caracter) && !unicode.IsSpace(caracter) {
			fmt.Print("Quel est votre nom ? ")
			fmt.Scan(&name)
		}
	}

	// Affichage du personnage avec une majuscule
	majuscule := cases.Title(language.French)
	caracter := MSA.InitCharacter(majuscule.String(name))

	println("  o		Nom :", caracter.Name, "\n /|\\|		Class :", caracter.Class, "\n / \\		PV:", caracter.Pv, "/", caracter.MaxPv)

	// Tours de jeu
	tour := 1

	// Boucle infinie
	for 0 != 1 {
		// 1er tour : affichage du message d'accueil
		if tour == 1 {
			fmt.Println("\n Vous débutez votre aventure dans le monde de la magie !")
		}
		// Le reste des tours
		//fmt.Println("Vous êtes")
		tour++
	}
}
