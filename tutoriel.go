package MSA

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Tutoriel(joueur *Character_class, tour int) int {
	// Changer la musique pour le combat
	JouerMusique("combat.mp3")

	// Pitch du tutoriel
	fmt.Println("Dans ce tutoriel, vous allez apprendre à mener un combat et à vous soigner !")
	// Pause
	fmt.Println("Appuyez sur Entrée pour continuer...")
	var pause string
	fmt.Scanln(&pause)

	ennemie := TutoEnnemi()

	// Nettoyage et graphismes améliorés
	Nettoyage(joueur) // Un seul appel suffit

	// Graphisme amélioré
	fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                         🎯 TUTORIEL 🎯                         ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║  👤 %-15s                    🎭 %-20s ║\n", "JOUEUR", strings.ToUpper(ennemie.Name))
	fmt.Printf("║  ❤️  %3d/%-3d                           ❤️  %3d/%-3d             ║\n",
		joueur.Pv, joueur.MaxPv, ennemie.Pv, ennemie.MaxPv)
	fmt.Printf("║  ⚔️  %-3d                               ⚔️  %-3d                 ║\n",
		joueur.Attaque, ennemie.Attaque)
	fmt.Printf("║  🛡️  %-3d                               🛡️  %-3d                 ║\n",
		joueur.Defence, ennemie.Defence)
	fmt.Println("╚════════════════════════════════════════════════════════════════╝")

	// ÉTAPE 1 : Apprentissage de l'attaque de base
	fmt.Println("\n📖 LEÇON 1 : Attaque de base")
	fmt.Println("Commençons par apprendre l'attaque à l'épée !")
	fmt.Println("Appuyez sur Entrée pour continuer...")
	fmt.Scanln(&pause)

	choix_attaque := 0
	attaque_faite := false

	for choix_attaque != 1 {
		fmt.Println("Attaquez avec votre épée !\n")
		fmt.Println("1] 🗡️ Attaquer <-- Choisissez cette option")
		fmt.Println("2] 📦 Regarder dans l'inventaire")
		fmt.Println("3] 🕊️ Quittez le terrain d'entraînement")
		fmt.Print("Quel est votre choix ? ")
		fmt.Scan(&choix_attaque)

		if choix_attaque == 1 {
			// Son
			JouerSon("../sounds/slash2.ogg")
			// Attaque du joueur - éviter les dégâts négatifs
			degats_joueur := joueur.Attaque - ennemie.Defence
			if degats_joueur < 1 {
				degats_joueur = 1
			}
			ennemie.Pv -= degats_joueur
			fmt.Println("\n✅ Très bien ! Vous infligez", degats_joueur, "dommages !")
			fmt.Println("L'ennemi vous attaque en retour...")
			// Pause
			fmt.Println("Appuyez sur Entrée pour continuer...")
			fmt.Scanln(&pause)
			// Attaque de l'ennemi - éviter les dégâts négatifs
			degats_ennemi := ennemie.Attaque - joueur.Defence
			if degats_ennemi < 1 {
				degats_ennemi = 1
			}
			joueur.Pv -= degats_ennemi
			fmt.Println("Le", ennemie.Name, "vous inflige", degats_ennemi, "dommages !\n")
			attaque_faite = true
			// Pause
			fmt.Print("Appuyez sur Entrée pour continuer...")
			fmt.Scanln()
		}
	}

	if attaque_faite {
		// Nettoyage
		Nettoyage(joueur)

		// Tour
		tour++

		// ÉTAPE 2 : Apprentissage des soins
		fmt.Println("\n📖 LEÇON 2 : Utilisation des potions")
		fmt.Println("L'ennemi vous a fait des dégâts, il faut vous soigner !")
		fmt.Println("Appuyez sur Entrée pour continuer...")
		fmt.Scanln(&pause)

		choix_soins := 0
		for choix_soins != 2 {
			fmt.Println("Soignez-vous avec l'inventaire !\n")
			fmt.Println("1] 🗡️ Attaquer")
			fmt.Println("2] 📦 Regarder dans l'inventaire <-- Choisissez cette option")
			fmt.Println("3] 🕊️ Fuir")
			fmt.Print("Quel est votre choix ? ")
			fmt.Scan(&choix_soins)

			if choix_soins == 2 {
				accessInventory(joueur, 1)

				// Utilisation des objets
				choix_inv := 0
				println("Utiliser un objet de l'inventaire ?")
				fmt.Println("1] Potion de soin <-- Choisissez cette option")
				fmt.Println("2] Potion de poison")
				fmt.Println("3] Quitter")
				fmt.Print("Quel est votre choix ? ")
				fmt.Scan(&choix_inv)

				// Choix
				switch choix_inv {
				case 1:
					// On force le joueur à utiliser une potion
					utilise := false
					for !utilise {
						var choix_inv int
						println("Vous devez utiliser une potion de soin pour continuer !")
						fmt.Println("1] Potion de soin")
						fmt.Print("Quel est votre choix ? ")
						fmt.Scan(&choix_inv)

						if choix_inv == 1 {
							utilise = utiliserObjetTuto(joueur, "Potion de soin")
						} else {
							fmt.Println("❌ Vous ne pouvez pas quitter, utilisez une potion de soin !")
						}
					}
				case 2:
					fmt.Println("Indisponible pour le tutoriel !")
				case 3:
					fmt.Println("Vous ne pouvez pas quitter le tutoriel !")
				}

				fmt.Println("\n✅ Excellent ! Vous maîtrisez les potions !")
				break
			}
		}

		// Tour
		tour++

		// ÉTAPE 3 : Apprentissage des sorts
		fmt.Println("\n📖 LEÇON 3 : Lancement de sorts")
		fmt.Println("Maintenant, apprenons à utiliser la magie !")
		fmt.Println("Vous allez apprendre à lancer vos sorts.")
		fmt.Println("Appuyez sur Entrée pour continuer...")
		fmt.Scanln(&pause)

		sort_appris := false
		for !sort_appris {
			// Nettoyage et graphismes
			Nettoyage(joueur) // Un seul appel

			fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
			fmt.Println("║                  🪄 APPRENTISSAGE DES SORTS 🪄                 ║")
			fmt.Println("╠════════════════════════════════════════════════════════════════╣")
			fmt.Printf("║  👤 %-15s                    🎭 %-20s ║\n", "JOUEUR", strings.ToUpper(ennemie.Name))
			fmt.Printf("║  ❤️  %3d/%-3d                           ❤️  %3d/%-3d             ║\n",
				joueur.Pv, joueur.MaxPv, ennemie.Pv, ennemie.MaxPv)
			fmt.Printf("║  ⚔️  %-3d                               ⚔️  %-3d                 ║\n",
				joueur.Attaque, ennemie.Attaque)
			fmt.Printf("║  🛡️  %-3d                               🛡️  %-3d                 ║\n",
				joueur.Defence, ennemie.Defence)
			fmt.Println("╚════════════════════════════════════════════════════════════════╝")

			fmt.Println("\nPour lancer un sort, choisissez d'abord 'Attaquer', puis 'Lancer un sort' !")
			fmt.Println("1] 🗡️ Attaquer")
			fmt.Println("2] 📦 Regarder dans l'inventaire")
			fmt.Println("3] 🕊️ Continuer le tutoriel")
			fmt.Print("Quel est votre choix ? ")
			fmt.Scan(&choix_attaque)

			if choix_attaque == 1 {
				// Choisir le type de l'attaque
				choix_type_attaque := 0
				fmt.Println("\nMaintenant, choisissez le type d'attaque :")
				fmt.Println("1] 🗡️  Attaque à l'épée")
				fmt.Println("2] 🪄  Lancer un sort <-- Choisissez cette option pour apprendre")
				fmt.Println("3] 🕊️  Retour")
				fmt.Print("Quel est votre choix ? ")
				fmt.Scan(&choix_type_attaque)

				if choix_type_attaque == 2 {
					// Choisir le sort
					choix_sort := 0
					fmt.Println("\n🌟 Parfait ! Choisissez votre sort :")
					fmt.Println("1]", joueur.Sorts[0][0], "- Dégâts:", joueur.Sorts[0][2])
					fmt.Println("2]", joueur.Sorts[1][0], "- Dégâts:", joueur.Sorts[1][2])
					fmt.Println("3] 🕊️  Retour")
					fmt.Print("Quel est votre choix ? ")
					fmt.Scan(&choix_sort)

					if choix_sort == 1 || choix_sort == 2 {
						// Conversion en int
						var dommages int
						var err error
						if choix_sort == 1 {
							dommages, err = strconv.Atoi(joueur.Sorts[0][2])
							if err != nil {
								dommages = 10
							}
							fmt.Println("\n✨ Vous lancez", joueur.Sorts[0][0])
							fmt.Println("📖", joueur.Sorts[0][1])
						} else {
							// Son
							JouerSon("../sounds/tonnerre.wav")
							dommages, err = strconv.Atoi(joueur.Sorts[1][2])
							if err != nil {
								dommages = 15
							}
							fmt.Println("\n✨ Vous lancez", joueur.Sorts[1][0])
							fmt.Println("📖", joueur.Sorts[1][1])
						}

						// Annimation simple d'attaque
						// Animation
						frames := []string{
							"   O ◊                        O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊  💥                    O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊     💥                  O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊        💥               O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊           💥            O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊              💥         O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊                 💥      O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊                    💥   O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊                    💥💢O\n  /|\\|                         |\\\n  / \\|                        / \\",
							"   O ◊                         O\n  /|\\|                        /|\\\n  / \\|                        / \\",
						}

						for _, frame := range frames {
							// Nettoyage de l’écran
							Nettoyage(joueur)
							fmt.Println(frame)
							time.Sleep(200 * time.Millisecond)
						}
						time.Sleep(1 * time.Second)

						// Nettoyage après l'animation
						Nettoyage(joueur)

						// Attaque magique
						ennemie.Pv -= dommages
						fmt.Println("🎯 Vous infligez", dommages, "dommages magiques !")

						// Attaque de l'ennemi
						degats_ennemi := ennemie.Attaque - joueur.Defence
						if degats_ennemi < 1 {
							degats_ennemi = 1
						}
						joueur.Pv -= degats_ennemi
						fmt.Println("Le", ennemie.Name, "vous inflige", degats_ennemi, "dommages en retour !")

						fmt.Println("\n✅ Excellent ! Vous maîtrisez maintenant les sorts !")
						fmt.Println("Appuyez sur Entrée pour continuer...")
						fmt.Scanln(&pause)
						sort_appris = true
					}
				} else if choix_type_attaque == 1 {
					fmt.Println("⚠️  Vous devez d'abord apprendre à utiliser vos sorts !")
				}
			} else if choix_attaque == 3 {
				// Forcer l'apprentissage des sorts
				fmt.Println("⚠️  Vous devez d'abord apprendre à utiliser vos sorts !")
				fmt.Println("Appuyez sur Entrée pour continuer...")
				fmt.Scanln(&pause)
			}
		}

		// ÉTAPE 4 : Combat final avec toutes les compétences
		fmt.Println("\n📖 LEÇON FINALE : Combat complet")
		fmt.Println("Maintenant que vous maîtrisez les attaques, les sorts et les potions,")
		fmt.Println("battez le", ennemie.Name, "en utilisant toutes vos compétences !")
		fmt.Println("Appuyez sur Entrée pour continuer...")
		fmt.Scanln(&pause)

		// Diminuer l'attaque du monstre pour le combat final
		ennemie.Attaque = 10

		// Boucle pour vaincre avec système complet
		for ennemie.Pv > 0 && joueur.Pv > 0 {
			// Tour
			tour++

			Nettoyage(joueur) // Un seul appel

			// Graphisme amélioré
			fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
			fmt.Println("║                 🏆 COMBAT FINAL D'ENTRAÎNEMENT 🏆              ║")
			fmt.Println("╠════════════════════════════════════════════════════════════════╣")
			fmt.Printf("║  👤 %-15s                    🎭 %-20s ║\n", "JOUEUR", strings.ToUpper(ennemie.Name))
			fmt.Printf("║  ❤️  %3d/%-3d                           ❤️  %3d/%-3d             ║\n",
				joueur.Pv, joueur.MaxPv, ennemie.Pv, ennemie.MaxPv)
			fmt.Printf("║  ⚔️  %-3d                               ⚔️  %-3d                 ║\n",
				joueur.Attaque, ennemie.Attaque)
			fmt.Printf("║  🛡️  %-3d                               🛡️  %-3d                 ║\n",
				joueur.Defence, ennemie.Defence)
			fmt.Println("╚════════════════════════════════════════════════════════════════╝")

			// Poison
			if ennemie.Empoisonne {
				if ennemie.Tour_poison <= 5 {
					ennemie.Pv -= 5
					ennemie.Tour_poison++
					fmt.Println("\nLe", ennemie.Name, "subit 5 points de dégâts de poison !")
					if ennemie.Pv < 0 {
						ennemie.Pv = 0
					}
				}
				if ennemie.Tour_poison >= 5 {
					ennemie.Empoisonne = false
					ennemie.Tour_poison = 0
					fmt.Println("Le", ennemie.Name, " a vaincu le poison !\n")
				}
			}

			// Choix de l'attaque
			choix_attaque := 0
			println("Utilisez toutes vos compétences ! Que voulez-vous faire ?")
			fmt.Println("1] 🗡️ Attaquer")
			fmt.Println("2] 📦 Regarder dans l'inventaire")
			fmt.Println("3] 🕊️ Fuir")
			fmt.Print("Quel est votre choix ? ")
			fmt.Scan(&choix_attaque)

			// Suppression des lignes de choix
			fmt.Print("\033[A\033[2K")
			fmt.Print("\033[A\033[2K")
			fmt.Print("\033[A\033[2K")
			fmt.Print("\033[A\033[2K")
			fmt.Print("\033[A\033[2K")

			// Gestion du combat (identique à combat.go)
			switch choix_attaque {
			case 1:
				// Choisir le type de l'attaque
				choix_type_attaque := 0
				fmt.Println("1] 🗡️  Attaque à l'épée")
				fmt.Println("2] 🪄  Lancer un sort")
				fmt.Println("3] 🕊️  Retour")
				fmt.Print("Quel est votre choix ? ")
				fmt.Scan(&choix_type_attaque)

				if choix_type_attaque == 1 {
					// Son
					JouerSon("../sounds/slash2.ogg")

					// Attaque du joueur - éviter les dégâts négatifs
					degats_joueur := joueur.Attaque - ennemie.Defence
					if degats_joueur < 1 {
						degats_joueur = 1
					}
					ennemie.Pv -= degats_joueur
					fmt.Println("\nVous infligez", degats_joueur, "dommages !")

					// Attaque de l'ennemi - éviter les dégâts négatifs
					degats_ennemi := ennemie.Attaque - joueur.Defence
					if degats_ennemi < 1 {
						degats_ennemi = 1
					}
					joueur.Pv -= degats_ennemi
					fmt.Println("Le", ennemie.Name, "vous inflige", degats_ennemi, "dommages !\n")

					// Pause
					fmt.Print("Appuyez sur Entrée pour continuer...")
					fmt.Scanln()

				} else if choix_type_attaque == 2 {
					// Choisir le sort
					choix_sort := 0
					fmt.Println("\nChoisissez votre sort :")
					fmt.Println("1]", joueur.Sorts[0][0])
					fmt.Println("2]", joueur.Sorts[1][0])
					fmt.Println("3] 🕊️  Retour")
					fmt.Print("Quel est votre choix ? ")
					fmt.Scan(&choix_sort)

					if choix_sort == 1 || choix_sort == 2 {
						var dommages int
						var err error

						if choix_sort == 1 {
							dommages, err = strconv.Atoi(joueur.Sorts[0][2])
							if err != nil {
								dommages = 10
							}
							fmt.Println("\nVous lancez", joueur.Sorts[0][0], "!")
							fmt.Println(ennemie.Name, "", joueur.Sorts[0][1])
						} else {
							// Son
							JouerSon("../sounds/tonnerre.wav")
							dommages, err = strconv.Atoi(joueur.Sorts[1][2])
							if err != nil {
								dommages = 15
							}
							fmt.Println("\nVous lancez", joueur.Sorts[1][0], "!")
							fmt.Println(ennemie.Name, "", joueur.Sorts[1][1])
						}

						// Animation
						frames := []string{
							"   O ◊                      O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊  💥                     O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊     💥                  O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊        💥               O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊           💥            O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊              💥         O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊                 💥      O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊                    💥   O\n  /|\\|                        /|\\\n  / \\|                        / \\",
							"   O ◊                    💥💢O\n  /|\\|                        |\\\n  / \\|                        / \\",
							"   O ◊                         O\n  /|\\|                        /|\\\n  / \\|                        / \\",
						}

						for _, frame := range frames {
							// Nettoyage de l’écran
							Nettoyage(joueur)
							fmt.Println(frame)
							time.Sleep(200 * time.Millisecond)
						}
						time.Sleep(1 * time.Second)

						// Nettoyage après l'animation
						Nettoyage(joueur)

						// Attaque magique
						ennemie.Pv -= dommages
						fmt.Println("Vous infligez", dommages, "dommages magiques !")

						// Attaque de l'ennemi
						degats_ennemi := ennemie.Attaque - joueur.Defence
						if degats_ennemi < 1 {
							degats_ennemi = 1
						}
						joueur.Pv -= degats_ennemi
						fmt.Println("Le", ennemie.Name, "vous inflige", degats_ennemi, "dommages !\n")

						// Pause
						fmt.Print("Appuyez sur Entrée pour continuer...")
						fmt.Scanln()
					}
				}

			case 2:
				accessInventory(joueur, 1)

				// Utilisation des objets
				inventaire_choix := 0
				println("Utiliser un objet de l'inventaire ?")
				fmt.Println("1] Potion de soin")
				fmt.Println("2] Potion de poison")
				fmt.Println("3] Quitter")
				fmt.Print("Quel est votre choix ? ")
				fmt.Scan(&inventaire_choix)

				switch inventaire_choix {
				case 1:
					utiliserObjet(joueur, &ennemie, "Potion de soin")
					// Attaque de l'ennemi après utilisation de la potion
					degats_ennemi := ennemie.Attaque - joueur.Defence
					if degats_ennemi < 1 {
						degats_ennemi = 1
					}
					joueur.Pv -= degats_ennemi
					fmt.Println("Le", ennemie.Name, "vous inflige", degats_ennemi, "dommages !\n")

					fmt.Print("Appuyez sur Entrée pour continuer...")
					fmt.Scanln()
				case 2:
					utiliserObjet(joueur, &ennemie, "Potion de poison")
					fmt.Println("Vous infligez 30 dommages de poison !")
					// L'ennemi attaque après
					degats_ennemi := ennemie.Attaque - joueur.Defence
					if degats_ennemi < 1 {
						degats_ennemi = 1
					}
					joueur.Pv -= degats_ennemi
					fmt.Println("Le", ennemie.Name, "vous inflige", degats_ennemi, "dommages !\n")

					fmt.Print("Appuyez sur Entrée pour continuer...")
					fmt.Scanln()

					// Poison
					ennemie.Empoisonne = true
				}

			case 3:
				fmt.Println("Vous ne pouvez pas fuir le tutoriel !")
				fmt.Print("Appuyez sur Entrée pour continuer...")
				fmt.Scanln()

			default:
				fmt.Println("Choix invalide")
			}

			// Vérification fin de combat
			if ennemie.Pv <= 0 {
				ennemie.Pv = 0
				Nettoyage(joueur)

				fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
				fmt.Println("║                      🏆 TUTORIEL TERMINÉ ! 🏆                  ║")
				fmt.Println("╠════════════════════════════════════════════════════════════════╣")
				fmt.Printf("║  👤 %-15s                    🎭 %-20s ║\n", "JOUEUR", strings.ToUpper(ennemie.Name))
				fmt.Printf("║  ❤️  %3d/%-3d                           ❤️  %3d/%-3d             ║\n",
					joueur.Pv, joueur.MaxPv, ennemie.Pv, ennemie.MaxPv)
				fmt.Printf("║  ⚔️  %-3d                               ⚔️  %-3d                 ║\n",
					joueur.Attaque, ennemie.Attaque)
				fmt.Printf("║  🛡️  %-3d                               🛡️  %-3d                 ║\n",
					joueur.Defence, ennemie.Defence)
				fmt.Println("╚════════════════════════════════════════════════════════════════╝")

				fmt.Println("\n🎉 Félicitations ! Vous avez terminé le tutoriel!")
				fmt.Println("✅ Vous maîtrisez maintenant :")
				fmt.Println("   • Les attaques à l'épée")
				fmt.Println("   • Les sorts magiques")
				fmt.Println("   • L'utilisation des potions")
				fmt.Println("\nVous êtes prêt pour de vrais combats !")

				fmt.Println("Appuyez sur Entrée pour continuer...")
				fmt.Scanln(&pause)
				Nettoyage(joueur)
				Nettoyage(joueur)

				// Vie réinitialisée
				joueur.Pv = joueur.MaxPv
				return tour + 1
			}

			if joueur.Pv <= 0 {
				fmt.Println("\n💀 Vous êtes mort pendant le tutoriel !")
				fmt.Println("🔄 Pas de souci, c'est le tutoriel! Retentez votre chance au terrain d'entraînement de la ville !")
				joueur.Pv = joueur.MaxPv
				ennemie.Pv = 50 // Reset l'ennemi aussi
				fmt.Print("Appuyez sur Entrée pour continuer...")
				fmt.Scanln()
			}
		}
	}
	return tour + 1
}
