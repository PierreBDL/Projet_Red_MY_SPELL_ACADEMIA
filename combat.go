package MSA

import (
	"fmt"
)

func Combat(joueur *Character_class, ennemie *Character_class, tour int) int {

	// Tant que l'ennemie ou le joueur a des PVs
	for ennemie.Pv > 0 && joueur.Pv > 0 {
		// Nettoyage de la console
		Nettoyage(joueur)

		// Graphisme
		fmt.Println("\n\n 	  o ◊ 		  o  ")
		fmt.Println(" 	 /|\\|		 \\/|\\")
		fmt.Println(" 	 / \\		  / \\")
		// Affichage des stats
		fmt.Println("        Joueur	       ", ennemie.Name)
		fmt.Println("❤️:  ", joueur.Pv, "/", joueur.MaxPv, "   		   ", ennemie.Pv, "/", ennemie.MaxPv)
		fmt.Println("🗡️:  	 ", joueur.Attaque, "	 		 ", ennemie.Attaque)
		fmt.Println("🛡️:	 ", joueur.Defence, "			 ", ennemie.Defence, "\n")

		fmt.Println("\nVous êtes au tour", tour, "\n")

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
			// Tour + 1
			tour++

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
				fmt.Println("🎉 Félicitation, vous gagnez", ennemie.Gold, "Or et", ennemie.Niveau, "niveau(x)!")

				// Ajout des récompenses (Or, expérience et vie)
				joueur.Gold += ennemie.Gold
				joueur.Niveau += ennemie.Niveau

				// Retour au menu
				return tour

			} else if joueur.Pv <= 0 {
				// Nettoyage de la console
				Nettoyage(joueur)

				// Graphisme
				fmt.Println("\n\n 	  o ◊ 		   o  ")
				fmt.Println(" 	 /|\\|		 \\/|\\")
				fmt.Println(" 	 / \\		  / \\")
				// Affichage des stats
				fmt.Println("        Joueur	       ", ennemie.Name)
				fmt.Println("PVs:  ", joueur.Pv, "/", joueur.MaxPv, "      ", ennemie.Pv, "/", ennemie.MaxPv)
				fmt.Println("Atq:  	 ", joueur.Attaque, "	 	 ", ennemie.Attaque)
				fmt.Println("Def:	 ", joueur.Defence, "		 ", ennemie.Defence, "\n")
				// Revoi au menu où se trouve la gestion de la mort
				return tour
			} else if joueur.Pv == 0 && ennemie.Pv == 0 {
				// Nettoyage de la console
				Nettoyage(joueur)

				// Graphisme
				fmt.Println("\n\n 	  o ◊ 		   o  ")
				fmt.Println(" 	 /|\\|		 \\/|\\")
				fmt.Println(" 	 / \\		  / \\")
				// Affichage des stats
				fmt.Println("        Joueur	       ", ennemie.Name)
				fmt.Println("PVs:  ", joueur.Pv, "/", joueur.MaxPv, "      ", ennemie.Pv, "/", ennemie.MaxPv)
				fmt.Println("Atq:  	 ", joueur.Attaque, "	 	 ", ennemie.Attaque)
				fmt.Println("Def:	 ", joueur.Defence, "		 ", ennemie.Defence, "\n")
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
			fmt.Println("2] Quitter")
			fmt.Print("Quel est votre choix ? ")
			fmt.Scan(&inventaire_choix)

			// Choix
			switch inventaire_choix {
			case 1:
				utiliserObjet(joueur, "Potion de soin")
			case 2:
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
