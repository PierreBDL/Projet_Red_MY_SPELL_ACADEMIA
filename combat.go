package MSA

import "fmt"

func Combat(joueur *Character_class, ennemie *Character_class) {
	// Graphisme
	println("  o ◊ 		   o  ")
	println(" /|\\|		 \\/|\\")
	println(" / \\		  / \\")

	// Tant que l'ennemie ou le joueur a des PVs
	for ennemie.Pv > 0 && joueur.Pv > 0 {
		choix_attaque := 0
		println("Un", ennemie.Name, "approche! Faîtes attention")
		fmt.Println("1] Attaquer")
		fmt.Println("2] Regarder dans l'inventaire")
		fmt.Println("3] Fuir")
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
			// Attaque du joueur
			ennemie.Pv -= (joueur.Attaque - ennemie.Defence)
			println("\nVous infligez", joueur.Attaque-ennemie.Defence, "dommages !")
			// Attaque de l'ennemis
			joueur.Pv -= (ennemie.Attaque - joueur.Defence)
			println("Le", ennemie.Name, "vous inflige", ennemie.Attaque-joueur.Defence, "dommages !\n")

			// Si le joueur ou l'ennemie à zéro PVs
			if ennemie.Pv <= 0 {
				ennemie.Pv = 0
				fmt.Println("Félicitation, vous gagnez !")
			}
			if joueur.Pv <= 0 {
				fmt.Println("Game Over !\nVous recommencez avec 50% de vos PVs")
				joueur.Pv = joueur.MaxPv / 2
				return
			}
		case 2:
			fmt.Println("Inventaire vide pour l'instant…\n")
		case 3:
			fmt.Println("Vous fuyez le combat !\n")
			return
		default:
			fmt.Println("Choix invalide")
			fmt.Scan(&choix_attaque)
		}

	}
}
