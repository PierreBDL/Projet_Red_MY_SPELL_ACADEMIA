package main

import (
	"MSA"
	"bufio"
	"fmt"
	"os"
	"strings"
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
		// Lecture des noms même avec des espaces pour éviter les bugs
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Quel est votre nom ? ")
		name, _ = reader.ReadString('\n') // Lit tout jusqu'à l'appui sur entrée
		name = strings.Trim(name, "\r\n") // enlève le \n à la fin et le \r du début

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
		// Nettoyer la console
		fmt.Print("\033[A\033[2K")
		fmt.Print("\033[A\033[2K")
	}

	// Affichage du personnage avec une majuscule
	majuscule := cases.Title(language.French)
	caracter := MSA.InitCharacter(majuscule.String(name))

	// Nettoyer la console
	MSA.Nettoyage(&caracter)

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
		// toujours vérifier si le joueur est mort
		if caracter.Pv <= 0 {
			// Gestion de la mort
			running = MSA.WasDead(&caracter)
		}

		// Vérifier si on est toujours en jeu (si on est mort, running = false donc ne rien afficher)
		if running {
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
				MSA.Nettoyage(&caracter)
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
			case 2:
				MSA.Nettoyage(&caracter)
				fmt.Println("\nVous entrez dans la ville de Musutafu !")
			case 3:
				running = false
				break
			default:
				break
			}
		}
	}
}
