package MSA

import (
	"fmt"
	// Atoi
	"strconv"
	"strings"
)

func Combat(joueur *Character_class, ennemie *Character_class, tour int) int {

	// Tant que l'ennemie ou le joueur a des PVs
	for ennemie.Pv > 0 && joueur.Pv > 0 {
		// Nettoyage de la console
		Nettoyage(joueur)
		Nettoyage(joueur)

		// Graphisme amélioré
		fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                        ⚔️  COMBAT  ⚔️                            ║")
		fmt.Println("╠════════════════════════════════════════════════════════════════╣")
		fmt.Printf("║  👤 %-15s                    🔥 %-15s      ║\n", "JOUEUR", strings.ToUpper(ennemie.Name))
		fmt.Printf("║  ❤️  %3d/%-3d                           ❤️  %3d/%-3d               ║\n",
			joueur.Pv, joueur.MaxPv, ennemie.Pv, ennemie.MaxPv)
		fmt.Printf("║  ⚔️  %-3d                               ⚔️  %-3d           	 ║\n",
			joueur.Attaque, ennemie.Attaque)
		fmt.Printf("║  🛡️  %-3d                               🛡️  %-3d           	 ║\n",
			joueur.Defence, ennemie.Defence)
		fmt.Println("╚════════════════════════════════════════════════════════════════╝")

		// Choix de l'attaque
		choix_attaque := 0
		println("Un", ennemie.Name, "approche! Faîtes attention")
		fmt.Println("1] 🗡️  Attaquer")
		fmt.Println("2] 📦 Regarder dans l'inventaire")
		fmt.Println("3] 🕊️  Fuir")
		fmt.Print("Quel est votre choix ? ")
		fmt.Scan(&choix_attaque)

		// Suppression des lignes de choix
		fmt.Print("\033[A\033[2K") // supprime la ligne du "Quel est votre choix ?"
		fmt.Print("\033[A\033[2K") // supprime la ligne "3] Fuir"
		fmt.Print("\033[A\033[2K") // supprime la ligne "2] Regarder dans l'inventaire"
		fmt.Print("\033[A\033[2K") // supprime la ligne "1] Attaquer"
		fmt.Print("\033[A\033[2K") // supprime la ligne "Un Ennemi approche!"

		// Gestion du combat
		switch choix_attaque {
		case 1:
			// Choisir le type de l'attaque
			choix_type_attaque := 0
			fmt.Println("1] 🗡️  Attaque à l'épée")
			fmt.Println("2] 🪄  Lancer un sort")
			fmt.Println("3] 🕊️  Retour")
			fmt.Print("Quel est votre choix ? ")
			fmt.Scan(&choix_type_attaque)

			// Mauvaise saisie
			if choix_type_attaque != 1 && choix_type_attaque != 2 && choix_type_attaque != 3 {
				fmt.Print("\033[A\033[A\033[A\033[A\033[2K")
				// Choisir le type de l'attaque
				choix_type_attaque := 0
				fmt.Println("1] 🗡️  Attaque à l'épée")
				fmt.Println("2] 🪄  Lancer un sort")
				fmt.Println("3] 🕊️  Retour")
				fmt.Print("Quel est votre choix ? ")
				fmt.Scan(&choix_type_attaque)
			} else {
				if choix_type_attaque == 1 {
					// Tour + 1
					tour++

					// Attaque du joueur - éviter les dégâts négatifs
					degats_joueur := joueur.Attaque - ennemie.Defence
					if degats_joueur < 1 {
						degats_joueur = 1 // Minimum 1 dégât
					}
					ennemie.Pv -= degats_joueur
					fmt.Println("\nVous infligez", degats_joueur, "dommages !")

					// Attaque de l'ennemi - éviter les dégâts négatifs
					degats_ennemi := ennemie.Attaque - joueur.Defence
					if degats_ennemi < 1 {
						degats_ennemi = 1 // Minimum 1 dégât
					}
					joueur.Pv -= degats_ennemi
					fmt.Println("Le", ennemie.Name, "vous inflige", degats_ennemi, "dommages !\n")
				} else if choix_type_attaque == 2 {
					// Choisir le type de l'attaque
					choix_sort := 0
					fmt.Println("\nChoississez votre sort :")
					fmt.Println("1]", joueur.Sorts[0][0])
					fmt.Println("2]", joueur.Sorts[1][0])
					fmt.Println("3] 🕊️  Retour")
					fmt.Print("Quel est votre choix ? ")
					fmt.Scan(&choix_sort)

					// Mauvaise saisie
					if choix_sort != 1 && choix_sort != 2 && choix_sort != 3 {
						fmt.Print("\033[A\033[A\033[A\033[A\033[2K")
						// Choisir le type de l'attaque
						choix_sort := 0
						fmt.Println("1]", joueur.Sorts[0][0])
						fmt.Println("2] ", joueur.Sorts[1][0])
						fmt.Println("3] 🕊️  Retour")
						fmt.Print("Quel est votre choix ? ")
						fmt.Scan(&choix_sort)
					} else {
						if choix_sort == 1 {
							// Conversion en int
							dommages, err := strconv.Atoi(joueur.Sorts[0][2])
							if err != nil {
								fmt.Println("Erreur de conversion :", err)
								dommages = 10 // ou autre valeur par défaut
							}
							// Tour + 1
							tour++

							// Attaque du joueur avec description du sort
							ennemie.Pv -= (dommages)
							fmt.Print("\nVous lancez", joueur.Sorts[0][0])
							fmt.Println(ennemie.Name, "", joueur.Sorts[0][1]) // Description du sort
							// Attaque de l'ennemis
							joueur.Pv -= (ennemie.Attaque - joueur.Defence)
							fmt.Println("Le", ennemie.Name, "vous inflige", ennemie.Attaque-joueur.Defence, "dommages !\n")

							// Pause pour permettre au joueur de lire
							fmt.Print("Appuyez sur Entrée pour continuer...")
							fmt.Scanln()
						} else if choix_sort == 2 {
							// Conversion en int
							dommages, err := strconv.Atoi(joueur.Sorts[1][2])
							if err != nil {
								fmt.Println("Erreur de conversion :", err)
								dommages = 15 // ou autre valeur par défaut
							}
							// Tour + 1
							tour++

							// Attaque du joueur avec description du sort
							ennemie.Pv -= (dommages)
							fmt.Print("\nVous lancez", joueur.Sorts[1][0], "!")
							fmt.Println(ennemie.Name, "", joueur.Sorts[1][1]) // Description du sort
							// Attaque de l'ennemis
							joueur.Pv -= (ennemie.Attaque - joueur.Defence)
							fmt.Println("Le", ennemie.Name, "vous inflige", ennemie.Attaque-joueur.Defence, "dommages !\n")

							// Pause pour permettre au joueur de lire
							fmt.Print("Appuyez sur Entrée pour continuer...")
							fmt.Scanln()
						} else if choix_sort == 3 {
							break
						}
					}
				}
			}

			// Si le joueur ou l'ennemie à zéro PVs
			if ennemie.Pv <= 0 {
				ennemie.Pv = 0
				// Nettoyage de la console
				Nettoyage(joueur)
				Nettoyage(joueur)

				// Graphisme
				fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
				fmt.Println("║                        ⚔️  COMBAT  ⚔️                            ║")
				fmt.Println("╠════════════════════════════════════════════════════════════════╣")
				fmt.Printf("║  👤 %-15s                    🔥 %-15s      ║\n", "JOUEUR", strings.ToUpper(ennemie.Name))
				fmt.Printf("║  ❤️  %3d/%-3d                           ❤️  %3d/%-3d               ║\n",
					joueur.Pv, joueur.MaxPv, ennemie.Pv, ennemie.MaxPv)
				fmt.Printf("║  ⚔️  %-3d                               ⚔️  %-3d           	 ║\n",
					joueur.Attaque, ennemie.Attaque)
				fmt.Printf("║  🛡️  %-3d                               🛡️  %-3d           	 ║\n",
					joueur.Defence, ennemie.Defence)
				fmt.Println("╚════════════════════════════════════════════════════════════════╝")
				fmt.Println("🎉 Félicitation, vous gagnez", ennemie.Gold, "Or et", ennemie.Niveau, "niveau(x)!")

				// Ajout des récompenses (Or, expérience et vie)
				joueur.Gold += ennemie.Gold
				joueur.Niveau += ennemie.Niveau

				// Retour au menu
				return tour

			} else if joueur.Pv <= 0 {
				// Nettoyage de la console
				Nettoyage(joueur)
				Nettoyage(joueur)

				// Graphisme
				fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
				fmt.Println("║                        ⚔️  COMBAT  ⚔️                            ║")
				fmt.Println("╠════════════════════════════════════════════════════════════════╣")
				fmt.Printf("║  👤 %-15s                    🔥 %-15s      ║\n", "JOUEUR", strings.ToUpper(ennemie.Name))
				fmt.Printf("║  ❤️  %3d/%-3d                           ❤️  %3d/%-3d               ║\n",
					joueur.Pv, joueur.MaxPv, ennemie.Pv, ennemie.MaxPv)
				fmt.Printf("║  ⚔️  %-3d                               ⚔️  %-3d           	 ║\n",
					joueur.Attaque, ennemie.Attaque)
				fmt.Printf("║  🛡️  %-3d                               🛡️  %-3d           	 ║\n",
					joueur.Defence, ennemie.Defence)
				fmt.Println("╚════════════════════════════════════════════════════════════════╝")
				// Revoi au menu où se trouve la gestion de la mort
				return tour
			} else if joueur.Pv == 0 && ennemie.Pv == 0 {
				// Nettoyage de la console
				Nettoyage(joueur)
				Nettoyage(joueur)

				// Graphisme
				fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
				fmt.Println("║                        ⚔️  COMBAT  ⚔️                            ║")
				fmt.Println("╠════════════════════════════════════════════════════════════════╣")
				fmt.Printf("║  👤 %-15s                    🔥 %-15s      ║\n", "JOUEUR", strings.ToUpper(ennemie.Name))
				fmt.Printf("║  ❤️  %3d/%-3d                           ❤️  %3d/%-3d               ║\n",
					joueur.Pv, joueur.MaxPv, ennemie.Pv, ennemie.MaxPv)
				fmt.Printf("║  ⚔️  %-3d                               ⚔️  %-3d           	 ║\n",
					joueur.Attaque, ennemie.Attaque)
				fmt.Printf("║  🛡️  %-3d                               🛡️  %-3d           	 ║\n",
					joueur.Defence, ennemie.Defence)
				fmt.Println("╚════════════════════════════════════════════════════════════════╝")
				// Revoi au menu où se trouve la gestion de la mort
				return tour
			}
		case 2:
			// Tour + 1
			tour++

			accessInventory(joueur, 1)

			// Utilisation des objets
			inventaire_choix := 0
			println("Utiliser un objet de l'inventaire ?")
			fmt.Println("1] Potion de soin")
			fmt.Println("2] Potion de poison")
			fmt.Println("3] Quitter")
			fmt.Print("Quel est votre choix ? ")
			fmt.Scan(&inventaire_choix)

			for inventaire_choix != 1 && inventaire_choix != 2 && inventaire_choix != 3 {
				fmt.Print("\033[A\033[A\033[A\033[A\033[2K")
				// Utilisation des objets
				inventaire_choix := 0
				println("Utiliser un objet de l'inventaire ?")
				fmt.Println("1] Potion de soin")
				fmt.Println("2] Potion de poison")
				fmt.Println("3] Quitter")
				fmt.Print("Quel est votre choix ? ")
				fmt.Scan(&inventaire_choix)
			}

			// Choix
			switch inventaire_choix {
			case 1:
				utiliserObjet(joueur, ennemie, "Potion de soin")
				// Attaque de l'ennemi après utilisation de la potion
				joueur.Pv -= (ennemie.Attaque - joueur.Defence)
				fmt.Println("Le", ennemie.Name, "vous inflige", ennemie.Attaque-joueur.Defence, "dommages !\n")

				// Pause pour permettre au joueur de lire
				fmt.Print("Appuyez sur Entrée pour continuer...")
				fmt.Scanln()
			case 2:
				utiliserObjet(joueur, ennemie, "Potion de poison")
				fmt.Println("Vous infligez 30 dommages de poison !")
				// L'ennemi attaque après que vous ayez utilisé la potion de poison
				joueur.Pv -= (ennemie.Attaque - joueur.Defence)
				fmt.Println("Le", ennemie.Name, "vous inflige", ennemie.Attaque-joueur.Defence, "dommages !\n")

				// Pause pour permettre au joueur de lire
				fmt.Print("Appuyez sur Entrée pour continuer...")
				fmt.Scanln()
			case 3:
				break
			}

		case 3:
			// Tour + 1
			tour++

			fmt.Println("Vous fuyez le combat !\n")
			return tour
		default:
			fmt.Println("Choix invalide")
			println("Un", ennemie.Name, "approche! Faîtes attention")
			fmt.Println("1] 🗡️  Attaquer")
			fmt.Println("2] 📦 Regarder dans l'inventaire")
			fmt.Println("3] 🕊️  Fuir")
			fmt.Print("Quel est votre choix ? ")
			fmt.Scan(&choix_attaque)
		}

	}
	return tour
}
