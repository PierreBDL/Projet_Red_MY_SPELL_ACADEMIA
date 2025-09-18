package MSA

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Combat(joueur *Character_class, ennemie *Character_class, tour int) int {

	// Musique de combat
	// Choisir aléartoirement une musique de combat
	rand.Seed(int64(time.Now().UnixNano()))
	combatMusic := []string{
		"../musics/combat/Tetris Thème Officiel.mp3",
		"../musics/combat/Joshua McLean - Mountain Trials ♫ NO COPYRIGHT 8-bit Music.mp3",
		"../musics/combat.mp3",
	}
	err := JouerMusique(combatMusic[rand.Intn(len(combatMusic))])
	if err != nil {
		fmt.Printf("🎵 Erreur musique combat: %v\n", err)
	}

	// Tant que l'ennemie ou le joueur a des PVs
	for ennemie.Pv > 0 && joueur.Pv > 0 {
		// Tour
		tour++

		// Nettoyage de la console
		Nettoyage(joueur)
		Nettoyage(joueur)

		// Graphisme amélioré
		fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                        ⚔️  COMBAT  ⚔️                          ║")
		fmt.Println("╠════════════════════════════════════════════════════════════════╣")
		fmt.Printf("║  👤 %-15s                    🔥 %-15s      ║\n", "JOUEUR", strings.ToUpper(ennemie.Name))
		fmt.Printf("║  ❤️  %3d/%-3d                           ❤️  %3d/%-3d             ║\n",
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
			for {
				fmt.Println("1] 🗡️  Attaque à l'épée")
				fmt.Println("2] 🪄  Lancer un sort")
				fmt.Println("3] 🕊️  Retour")
				fmt.Print("Quel est votre choix ? ")
				fmt.Scan(&choix_type_attaque)
				if choix_type_attaque == 1 || choix_type_attaque == 2 || choix_type_attaque == 3 {
					break
				}
				fmt.Println("❌ Choix invalide, recommencez.")
			}

			if choix_type_attaque == 1 {
				// Tour + 1
				tour++

				// Son
				JouerSon("../sounds/slash2.ogg")

				// Attaque du joueur
				degats_joueur := joueur.Attaque - ennemie.Defence
				if degats_joueur < 1 {
					degats_joueur = 1 // Minimum 1 dégât
				}
				ennemie.Pv -= degats_joueur
				fmt.Println("\nVous infligez", degats_joueur, "dommages !")

				// Attaque de l'ennemi avec système amélioré
				var degats_ennemi int
				aleatoire_attaque_ennemie := Ennemis_type_attaque()
				if aleatoire_attaque_ennemie >= 0 && aleatoire_attaque_ennemie < len(ennemie.Type_attaque) {
					// Affichage du titre et de la description de l'attaque
					fmt.Printf(" %s : %s\n", ennemie.Type_attaque[aleatoire_attaque_ennemie][0], ennemie.Type_attaque[aleatoire_attaque_ennemie][1])

					// Conversion string vers int
					degatsStr := ennemie.Type_attaque[aleatoire_attaque_ennemie][2]
					degats, err := strconv.Atoi(degatsStr)
					if err != nil {
						// En cas d'erreur, utiliser l'attaque de base
						degats_ennemi = ennemie.Attaque - joueur.Defence
					} else {
						degats_ennemi = (degats + ennemie.Attaque) - joueur.Defence
					}
				} else {
					fmt.Printf(" Attaque de base : Le %s vous attaque !\n", ennemie.Name)
					degats_ennemi = ennemie.Attaque - joueur.Defence
				}

				// IMPORTANT : Protection contre les dégâts négatifs
				if degats_ennemi < 1 {
					degats_ennemi = 1 // Minimum 1 dégât
				}

				joueur.Pv -= degats_ennemi
				fmt.Println("Le", ennemie.Name, "vous inflige", degats_ennemi, "dommages !\n")

				// Pause
				fmt.Print("Appuyez sur Entrée pour continuer...")
				fmt.Scanln()

			} else if choix_type_attaque == 2 {
				// Choisir le type de l'attaque
				choix_sort := 0
				fmt.Println("\nChoississez votre sort :")

				// Afficher les sorts avec leur coût en mana
				sorts := InitSpells()
				for i := 0; i < len(joueur.Sorts) && i < len(sorts); i++ {
					if joueur.Mana >= sorts[i].Mana_cost {
						fmt.Printf("%d] %s (Coût: %d 💧)\n", i+1, joueur.Sorts[i][0], sorts[i].Mana_cost)
					} else {
						fmt.Printf("%d] %s (Coût: %d 💧) Pas assez de mana\n", i+1, joueur.Sorts[i][0], sorts[i].Mana_cost)
					}
				}
				fmt.Println("4] 🕊️  Retour")
				fmt.Print("Quel est votre choix ? ")
				fmt.Scan(&choix_sort)

				// Validation du choix
				for choix_sort != 1 && choix_sort != 2 && choix_sort != 3 && choix_sort != 4 {
					fmt.Print("\033[A\033[A\033[A\033[A\033[2K")
					fmt.Println("Choix invalide !")
					for i := 0; i < len(joueur.Sorts) && i < len(sorts); i++ {
						if joueur.Mana >= sorts[i].Mana_cost {
							fmt.Printf("%d] %s (Coût: %d 💧)\n", i+1, joueur.Sorts[i][0], sorts[i].Mana_cost)
						} else {
							fmt.Printf("%d] %s (Coût: %d 💧)Pas assez de mana\n", i+1, joueur.Sorts[i][0], sorts[i].Mana_cost)
						}
					}
					fmt.Println("4] 🕊️  Retour")
					fmt.Print("Quel est votre choix ? ")
					fmt.Scan(&choix_sort)
				}

				if choix_sort == 1 {
					// Vérifier si le joueur a assez de mana
					if joueur.Mana >= sorts[0].Mana_cost {
						// Tour + 1
						tour++

						// Utiliser le mana
						joueur.Mana -= sorts[0].Mana_cost

						// Attaque du joueur
						dommages := sorts[0].Degats
						ennemie.Pv -= dommages
						fmt.Printf("\nVous lancez %s!\n", joueur.Sorts[0][0])
						fmt.Printf("%s %s\n", ennemie.Name, joueur.Sorts[0][1])
						fmt.Printf("Vous infligez %d dommages magiques !\n", dommages)

						// Attaque de l'ennemi
						var degats_ennemi int
						aleatoire_attaque_ennemie := Ennemis_type_attaque()
						if aleatoire_attaque_ennemie >= 0 && aleatoire_attaque_ennemie < len(ennemie.Type_attaque) {
							fmt.Printf("%s : %s\n", ennemie.Type_attaque[aleatoire_attaque_ennemie][0], ennemie.Type_attaque[aleatoire_attaque_ennemie][1])
							degatsStr := ennemie.Type_attaque[aleatoire_attaque_ennemie][2]
							degats, err := strconv.Atoi(degatsStr)
							if err != nil {
								degats_ennemi = ennemie.Attaque - joueur.Defence
							} else {
								degats_ennemi = degats - joueur.Defence
							}
						} else {
							fmt.Printf("Attaque de base : Le %s vous attaque !\n", ennemie.Name)
							degats_ennemi = ennemie.Attaque - joueur.Defence
						}
						if degats_ennemi < 1 {
							degats_ennemi = 1
						}
						joueur.Pv -= degats_ennemi
						fmt.Println("Le", ennemie.Name, "vous inflige", degats_ennemi, "dommages !\n")

						// Pause
						fmt.Print("Appuyez sur Entrée pour continuer...")
						fmt.Scanln()
					} else {
						fmt.Print("Pas assez de mana !")
						fmt.Print("Appuyez sur Entrée pour continuer...")
						fmt.Scanln()
						continue // Retour au menu d'attaque
					}
				} else if choix_sort == 2 {
					// Vérifier si le joueur a assez de mana
					if joueur.Mana >= sorts[1].Mana_cost {
						// Tour + 1
						tour++

						// Utiliser le mana
						joueur.Mana -= sorts[1].Mana_cost

						// Son pour le sort Incendio
						JouerSon("../sounds/tonnerre.wav")

						// Attaque du joueur
						dommages := sorts[1].Degats
						ennemie.Pv -= dommages
						fmt.Printf("\nVous lancez %s!\n", joueur.Sorts[1][0])
						fmt.Printf("%s %s\n", ennemie.Name, joueur.Sorts[1][1])
						fmt.Printf("Vous infligez %d dommages magiques !\n", dommages)
						// Attaque de l'ennemi
						var degats_ennemi int
						aleatoire_attaque_ennemie := Ennemis_type_attaque()
						if aleatoire_attaque_ennemie >= 0 && aleatoire_attaque_ennemie < len(ennemie.Type_attaque) {
							fmt.Printf(" %s : %s\n", ennemie.Type_attaque[aleatoire_attaque_ennemie][0], ennemie.Type_attaque[aleatoire_attaque_ennemie][1])
							degatsStr := ennemie.Type_attaque[aleatoire_attaque_ennemie][2]
							degats, err := strconv.Atoi(degatsStr)
							if err != nil {
								degats_ennemi = ennemie.Attaque - joueur.Defence
							} else {
								degats_ennemi = degats - joueur.Defence
							}
						} else {
							fmt.Printf("Attaque de base : Le %s vous attaque !\n", ennemie.Name)
							degats_ennemi = ennemie.Attaque - joueur.Defence
						}
						if degats_ennemi < 1 {
							degats_ennemi = 1
						}
						joueur.Pv -= degats_ennemi
						fmt.Println("Le", ennemie.Name, "vous inflige", degats_ennemi, "dommages !\n")

						// Pause
						fmt.Print("Appuyez sur Entrée pour continuer...")
						fmt.Scanln()
					} else {
						fmt.Print("Pas assez de mana !\n")
						fmt.Print("Appuyez sur Entrée pour continuer...")
						fmt.Scanln()
						continue // Retour au menu d'attaque
					}
				} else if choix_sort == 3 {
					// Vérifier si le joueur a assez de mana
					if joueur.Mana >= sorts[2].Mana_cost {
						// Tour + 1
						tour++

						// Utiliser le mana
						joueur.Mana -= sorts[2].Mana_cost

						// Son pour le sort Incendio
						JouerSon("../sounds/tonnerre.wav")

						// Attaque du joueur
						dommages := sorts[2].Degats
						ennemie.Pv -= dommages
						fmt.Printf("\nVous lancez %s!\n", joueur.Sorts[2][0])
						fmt.Printf("%s %s\n", ennemie.Name, joueur.Sorts[2][1])
						fmt.Printf("Vous infligez %d dommages magiques !\n", dommages)

						// Attaque de l'ennemi
						var degats_ennemi int
						aleatoire_attaque_ennemie := Ennemis_type_attaque()
						if aleatoire_attaque_ennemie >= 0 && aleatoire_attaque_ennemie < len(ennemie.Type_attaque) {
							fmt.Printf(" %s : %s\n", ennemie.Type_attaque[aleatoire_attaque_ennemie][0], ennemie.Type_attaque[aleatoire_attaque_ennemie][1])
							degatsStr := ennemie.Type_attaque[aleatoire_attaque_ennemie][2]
							degats, err := strconv.Atoi(degatsStr)
							if err != nil {
								degats_ennemi = ennemie.Attaque - joueur.Defence
							} else {
								degats_ennemi = degats - joueur.Defence
							}
						} else {
							fmt.Printf("Attaque de base : Le %s vous attaque !\n", ennemie.Name)
							degats_ennemi = ennemie.Attaque - joueur.Defence
						}
						if degats_ennemi < 1 {
							degats_ennemi = 1
						}
						joueur.Pv -= degats_ennemi
						fmt.Println("Le", ennemie.Name, "vous inflige", degats_ennemi, "dommages !\n")

						// Pause
						fmt.Print("Appuyez sur Entrée pour continuer...")
						fmt.Scanln()
					} else {
						fmt.Print("Pas assez de mana !\n")
						fmt.Print("Appuyez sur Entrée pour continuer...")
						fmt.Scanln()
						continue // Retour au menu d'attaque
					}
				} else if choix_type_attaque == 4 {
					continue // Retour au menu principal
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
				fmt.Println("║                        ⚔️  COMBAT  ⚔️                          ║")
				fmt.Println("╠════════════════════════════════════════════════════════════════╣")
				fmt.Printf("║  👤 %-15s                    🔥 %-15s      ║\n", "JOUEUR", strings.ToUpper(ennemie.Name))
				fmt.Printf("║  ❤️  %3d/%-3d                           ❤️  %3d/%-3d             ║\n",
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
				fmt.Println("║                        ⚔️  COMBAT  ⚔️                          ║")
				fmt.Println("╠════════════════════════════════════════════════════════════════╣")
				fmt.Printf("║  👤 %-15s                    🔥 %-15s      ║\n", "JOUEUR", strings.ToUpper(ennemie.Name))
				fmt.Printf("║  ❤️  %3d/%-3d                           ❤️  %3d/%-3d             ║\n",
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
				fmt.Println("║                        ⚔️  COMBAT  ⚔️                          ║")
				fmt.Println("╠════════════════════════════════════════════════════════════════╣")
				fmt.Printf("║  👤 %-15s                    🔥 %-15s      ║\n", "JOUEUR", strings.ToUpper(ennemie.Name))
				fmt.Printf("║  ❤️  %3d/%-3d                           ❤️  %3d/%-3d             ║\n",
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
