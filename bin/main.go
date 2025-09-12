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
	for {
		fmt.Print("Quel est votre nom ? ")
		fmt.Scan(&name)

		// Vérifier la validité du nom
		valide := true
		for _, caracter := range name {
			// Vérifier la validité du nom (pas de chiffres ou de caractères spéciaux)
			if !unicode.IsLetter(caracter) && !unicode.IsSpace(caracter) {
				valide = false
				break
			}
		}
		if valide {
			break
		}
		fmt.Print("\033[A\033[2K") // Remonte et efface les dernières lignes
	}

	// Affichage du personnage avec une majuscule
	majuscule := cases.Title(language.French)
	caracter := MSA.InitCharacter(majuscule.String(name))

	println("\n  o ◊		Nom :", caracter.Name, "\n /|\\|		Class :", caracter.Class, "\n / \\		PV:", caracter.Pv, "/", caracter.MaxPv)

	// Tours de jeu
	tour := 1

	// Boucle infini
	running := true

	// Boucle infinie
	for running == true {
		// 1er tour : affichage du message d'accueil
		if tour == 1 {
			fmt.Println("\nVous débutez votre aventure dans le monde de la magie !")
		}
		// Le reste des tours
		choix := 0
		fmt.Println("\nVous êtes au tour", tour, "que voulez vous faire ?\n")
		fmt.Println("1] Avancer")
		fmt.Println("2] Aller en ville")
		fmt.Println("3] Quitter le jeu")
		fmt.Print("Quel est votre choix ? ")
		fmt.Scan(&choix)

		// Annalyse du choix
		for choix < 1 || choix > 3 {
			fmt.Print("Choix invalide. Veuillez recommencer ")
			fmt.Print("\033[A\033[2K") // Remonte et efface les dernières lignes
			fmt.Print("Quel est votre choix ? ")
			fmt.Scan(&choix)
		}
		switch choix {
		case 1:
			// Avancé du tour
			tour++
			// Probabilité d'un ennemi
			if MSA.Avancer() == true {
				// Faire apparaître l'ennemi
				ennemie := MSA.InitEnnemi()
				// S'il y a un ennemi
				if ennemie.Name != "" {
					MSA.Combat(&caracter, &ennemie)
				}
			}
			break
		case 2:
			break
		case 3:
			running = false
			break
		default:
			break
		}
	}
}
