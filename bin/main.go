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

	// Choix de la classe du personnage
	choix_classe := 0
	fmt.Println("\nQue voulez-vous être ?\n")
	fmt.Println("1] Sorcier")
	fmt.Println("2] Alchimiste")
	fmt.Print("Quel est votre choix ? ")
	fmt.Scan(&choix_classe)
	fmt.Print("")

	// Variable de classe
	var caracter MSA.Character_class

	// Choix de la classe
	switch choix_classe {
	case 1:
		// Création du perso avec la classe sorcier
		majuscule := cases.Title(language.French)
		caracter = MSA.InitCharacter(majuscule.String(name), 1)
	case 2:
		// Création du perso avec la classe alchimiste
		majuscule := cases.Title(language.French)
		caracter = MSA.InitCharacter(majuscule.String(name), 2)
	case 666:
		// Easter Egg Harry Poter
		majuscule := cases.Title(language.French)
		caracter = MSA.InitCharacter(majuscule.String(""), 666)
	default:
		// Nettoyer la console
		fmt.Print("\033[A\033[2K")
		fmt.Print("\033[A\033[2K")
		fmt.Print("Quel est votre choix ? ")
		fmt.Scan(&choix_classe)

	}

	// Nettoyer la console
	MSA.Nettoyage(&caracter)

	// Tours de jeu
	tour := 1

	// Boucle infini
	running := true

	// Boucle infinie
	for running == true {
		// 1er tour : affichage du message d'accueil et tutoriel
		if tour == 1 {
			fmt.Println("\nVous débutez votre aventure dans le monde de la magie !")
			fmt.Println("\nDébut du tutoriel")
			tour = MSA.Tutoriel(&caracter, tour)
		}
		// toujours vérifier si le joueur est mort
		if caracter.Pv <= 0 {
			// Gestion de la mort
			running = MSA.WasDead(&caracter)
		}

		// Vérifier si on est toujours en jeu (si on est mort, running = false donc ne rien afficher)
		if running {
			fmt.Println("╔════════════════════════════════════════════════════════════════╗")
			fmt.Println("║                    🎮 MY SPELL ACADEMIA 🎮                     ║")
			fmt.Println("╠════════════════════════════════════════════════════════════════╣")
			fmt.Printf("║                        🕒 Tour %-3d                             ║\n", tour)
			fmt.Println("║                                                                ║")
			fmt.Println("║              ⏩ 1] Avancer dans l'aventure                     ║")
			fmt.Println("║              🏙️  2] Aller en ville                              ║")
			fmt.Println("║              🚪 3] Quitter le jeu                              ║")
			fmt.Println("║                                                                ║")
			fmt.Println("╚════════════════════════════════════════════════════════════════╝")
			fmt.Print("➤ Que voulez-vous faire ? ")

			// Annalyse du choix
			choix := 0
			fmt.Scan(&choix)
			fmt.Print("")

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
				// Choix de l'ennemie
				ennemie := MSA.Avancer()
				if ennemie == "gobelin" {
					// Faire apparaître l'ennemi
					ennemie := MSA.InitEnnemi("gobelin")
					// S'il y a un ennemi
					if ennemie.Name != "" {
						tour = MSA.Combat(&caracter, &ennemie, tour)
					}
				}
				if ennemie == "boss" {
					// Faire apparaître l'ennemi
					ennemie := MSA.InitEnnemi("boss")
					// S'il y a un ennemi
					if ennemie.Name != "" {
						tour = MSA.Combat(&caracter, &ennemie, tour)
					}
				}
			case 2:
				MSA.Nettoyage(&caracter)
				MSA.Entree_ville(&caracter)
			case 3:
				running = false
			default:
				break
			}
		}
	}
}
