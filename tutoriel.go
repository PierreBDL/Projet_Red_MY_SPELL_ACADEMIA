package MSA

import (
	"fmt"
)

func Tutoriel(joueur *Character_class) {
	// Pitch du tutoriel
	fmt.Println("Dans ce tutoriel, vous allez apprendre à mener un combat")
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
	fmt.Println("Attaquez !")
	choix_attaque := 0
	fmt.Println("1] 🗡️  Attaquer <--")
	fmt.Println("2] 📦 Regarder dans l'inventaire")
	fmt.Println("3] 🕊️  Fuir")
	fmt.Print("Quel est votre choix ? ")
	fmt.Scan(&choix_attaque)

	switch choix_attaque {
	case 1:
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
	default:
		fmt.Println("Attaquez !")
		fmt.Println("1] 🗡️  Attaquer <--")
		fmt.Println("2] 📦 Regarder dans l'inventaire")
		fmt.Println("3] 🕊️  Fuir")
		fmt.Print("Quel est votre choix ? ")
		fmt.Scan(&choix_attaque)
	}

	// Tuto soins
	fmt.Println("L'ennemi vous a fais des dégats, il faut vous soigner !")
	fmt.Println("Appuyez sur Entrée pour continuer...")
	var pause string
	fmt.Scanln(&pause)
	fmt.Println("Ouvrez votre inventaire")

	fmt.Println("Soignez-vous !\n")
	choix_soins := 0
	fmt.Println("1] 🗡️  Attaquer")
	fmt.Println("2] 📦 Regarder dans l'inventaire <--")
	fmt.Println("3] 🕊️  Fuir")
	fmt.Print("Quel est votre choix ? ")
	fmt.Scan(&choix_soins)

	switch choix_attaque {
	case 2:
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

		fmt.Println("Bravo ! Le tutoriel est fini !")
		return

	default:
		fmt.Println("Soignez-vous !\n")
		fmt.Println("1] 🗡️  Attaquer")
		fmt.Println("2] 📦 Regarder dans l'inventaire <--")
		fmt.Println("3] 🕊️  Fuir")
		fmt.Print("Quel est votre choix ? ")
		fmt.Scan(&choix_soins)
	}
}
