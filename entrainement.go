package MSA

import (
	"fmt"
)

func Entrainement(joueur *Character_class) {
	// Pitch du tutoriel
	fmt.Println("Dans cet entraînement, vous allez apprendre à mener un combat et à vous soigner !")
	// Pause
	fmt.Println("Appuyez sur Entrée pour continuer...")
	var pause string
	fmt.Scanln(&pause)

	ennemie := TutoEnnemi()

	// Arrivée de l'ennemi
	fmt.Println("Reprenons les bases !")
	fmt.Println(" 	  o ◊ 			   o  ")
	fmt.Println(" 	 /|\\|			 \\/|\\")
	fmt.Println(" 	 / \\|			  / \\")
	// Affichage des stats
	fmt.Println("        Joueur	       ", ennemie.Name)
	fmt.Println("❤️:   ", joueur.Pv, "/", joueur.MaxPv, "   	       ", ennemie.Pv, "/", ennemie.MaxPv)
	fmt.Println("🗡️:   	", joueur.Attaque, "	 	 	 ", ennemie.Attaque)
	fmt.Println("🛡️:   	", joueur.Defence, "		 	 ", ennemie.Defence, "\n")

	// Tuto attaque
	choix_attaque := 0
	attaque_faite := false

	for choix_attaque != 1 {
		fmt.Println("Attaquez !\n")
		fmt.Println("1] 🗡️  Attaquer <--")
		fmt.Println("2] 📦 Regarder dans l'inventaire")
		fmt.Println("3] 🕊️  Quittez le terrain d'entraînement")
		fmt.Print("Quel est votre choix ? ")
		fmt.Scan(&choix_attaque)

		if choix_attaque == 1 {
			// Attaque du joueur
			ennemie.Pv -= (joueur.Attaque - ennemie.Defence)
			fmt.Println("\nVous infligez", joueur.Attaque-ennemie.Defence, "dommages !")
			fmt.Print("L'ennemi vous attaque aussi")
			// Pause
			fmt.Println("Appuyez sur Entrée pour continuer...")
			var pause string
			fmt.Scanln(&pause)
			// Attaque de l'ennemi
			joueur.Pv -= (ennemie.Attaque - joueur.Defence)
			println("Le", ennemie.Name, "vous inflige", ennemie.Attaque-joueur.Defence, "dommages !\n")
			attaque_faite = true
		}

		// Quitter
		if choix_attaque == 3 {
			return
		}

		if attaque_faite == true {
			// Tuto soins
			fmt.Println("L'ennemi vous a fais des dégats, il faut vous soigner !")
			fmt.Println("Appuyez sur Entrée pour continuer...")
			var pause string
			fmt.Scanln(&pause)
			fmt.Println("")

			choix_soins := 0
			for choix_soins != 2 {
				fmt.Println("Soignez-vous !\n")
				fmt.Println("1] 🗡️  Attaquer")
				fmt.Println("2] 📦 Regarder dans l'inventaire <--")
				fmt.Println("3] 🕊️  Fuir")
				fmt.Print("Quel est votre choix ? ")
				fmt.Scan(&choix_soins)

				if choix_soins == 2 {
					accessInventory(joueur, 1)

					// Utilisation des objets
					choix_inv := 0
					println("Utiliser un objet de l'inventaire ?")
					fmt.Println("1] Potion de soin")
					fmt.Println("2] Quitter")
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
						// Empêcher de quitter dans le tutoriel
						fmt.Println("Vous quittez le terrain d'entraînement")
						return
					}

					fmt.Println("")
					fmt.Println("Bravo ! Essayez de vaincre le monstre !")
					fmt.Println("Appuyez sur Entrée pour continuer...")
					var pause string
					fmt.Scanln(&pause)

					// Vaincre le monstre
					// Diminuer l'attaque du monstre
					ennemie.Attaque = 10
					// Boucle pour vaincre
					for ennemie.Pv > 0 {
						Nettoyage(joueur)
						Nettoyage(joueur)
						// Tant que l'ennemie ou le joueur a des PVs

						// Graphisme
						fmt.Println("\n\n 	  o ◊ 		  o  ")
						fmt.Println(" 	 /|\\|		 \\/|\\")
						fmt.Println(" 	 / \\		  / \\")
						// Affichage des stats
						fmt.Println("        Joueur	       ", ennemie.Name)
						fmt.Println("❤️:  ", joueur.Pv, "/", joueur.MaxPv, "   		   ", ennemie.Pv, "/", ennemie.MaxPv)
						fmt.Println("🗡️:  	 ", joueur.Attaque, "	 		 ", ennemie.Attaque)
						fmt.Println("🛡️:	 ", joueur.Defence, "			 ", ennemie.Defence, "\n")

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
							fmt.Print("\033[A\033[2K") // supprime la ligne attaque du joueur
							fmt.Print("\033[A\033[2K") // supprime la ligne attaque de l'ennemi
							// Attaque du joueur
							ennemie.Pv -= (joueur.Attaque - ennemie.Defence)
							println("\nVous infligez", joueur.Attaque-ennemie.Defence, "dommages !")
							// Attaque de l'ennemis
							joueur.Pv -= (ennemie.Attaque - joueur.Defence)
							println("Le", ennemie.Name, "vous inflige", ennemie.Attaque-joueur.Defence, "dommages !\n")

							// Si le joueur ou l'ennemie à zéro PVs
							if ennemie.Pv <= 0 {
								ennemie.Pv = 0
								// Nettoyage de la console
								Nettoyage(joueur)

								// Graphisme
								fmt.Println("\n\n 	  o ◊ 		   o  ")
								fmt.Println(" 	 /|\\|		 \\/|\\")
								fmt.Println(" 	 / \\		  / \\")
								// Affichage des stats
								fmt.Println("        Joueur	       ", ennemie.Name)
								fmt.Println("❤️ :  ", joueur.Pv, "/", joueur.MaxPv, "      ", ennemie.Pv, "/", ennemie.MaxPv)
								fmt.Println("🗡️ :  	 ", joueur.Attaque, "	 	 ", ennemie.Attaque)
								fmt.Println("🛡️ :	 ", joueur.Defence, "		 ", ennemie.Defence, "\n")

								// Quitter le tuto
								fmt.Println("Bravo ! L'entraînement est fini !")
								fmt.Println("Appuyez sur Entrée pour continuer...")
								var pause string
								fmt.Scanln(&pause)
								Nettoyage(joueur)
								Nettoyage(joueur)

								// Vie réinitialisée
								joueur.Pv = joueur.MaxPv

								return
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

							for inventaire_choix != 1 && inventaire_choix != 2 && inventaire_choix != 3 {
								fmt.Print("\033[A\033[A\033[A\033[A\033[2K")
								// Utilisation des objets
								inventaire_choix = 0
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
								utiliserObjet(joueur, &ennemie, "Potion de soin")
								// Attaque de l'ennemi après utilisation de la potion
								joueur.Pv -= (ennemie.Attaque - joueur.Defence)
								fmt.Println("Le", ennemie.Name, "vous inflige", ennemie.Attaque-joueur.Defence, "dommages !\n")

								// Pause pour permettre au joueur de lire
								fmt.Print("Appuyez sur Entrée pour continuer...")
								fmt.Scanln()
							case 2:
								utiliserObjet(joueur, &ennemie, "Potion de poison")
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
				}
			}
		}
	}
}
