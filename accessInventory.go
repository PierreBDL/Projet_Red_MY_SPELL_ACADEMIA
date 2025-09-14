package MSA

import "fmt"

func accessInventory(joueur *Character_class, action int) {
	// Voir inventaire
	if action == 1 {
		// Si l'inventaire est vide
		fmt.Println("Dans votre inventaire :")
		if len(joueur.Inventaire) == 0 {
			fmt.Println("Inventaire vide")
			return
		}
		// Création d'un dictionnaire pour savoir combien il y a d'objet de chaque type
		compte := make(map[string]int)
		for _, objet := range joueur.Inventaire {
			compte[objet]++
		}
		// Affichage
		for objet, nb := range compte {
			fmt.Printf("x%d : %s\n", nb, objet)
		}
	}
}

func utiliserObjet(joueur *Character_class, objet string) {
	// Chercher si on a l'objet
	recherche := objet
	trouve := false
	var position_objet int

	for i, potion := range joueur.Inventaire {
		if potion == recherche {
			trouve = true
			// Sera utile pour consommer l'objet
			position_objet = i
			break
		}
	}

	if trouve {
		if objet == "Potion de soin" {
			joueur.Pv += 50
			// Eviter dépassement des PVs max
			if joueur.Pv > joueur.MaxPv {
				joueur.Pv = joueur.MaxPv
			}
			// Consommer objet
			joueur.Inventaire = append(joueur.Inventaire[:position_objet], joueur.Inventaire[position_objet+1:]...)
			fmt.Println("Vous utilisez une Potion de soin ! PV :", joueur.Pv, "/", joueur.MaxPv)
			return
		}
	} else {
		fmt.Println(recherche, "n'est pas dans votre inventaire")
	}
}
