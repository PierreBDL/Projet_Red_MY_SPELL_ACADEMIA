package MSA

import "fmt"

// Voir inventaire ou autre action
func accessInventory(joueur *Character_class, action int) {
	if action == 1 {
		fmt.Println("")
		fmt.Println("Dans votre inventaire :")
		fmt.Println("")
		if len(joueur.Inventaire) == 0 {
			fmt.Println("Inventaire vide")
			return
		}
		// Affichage des objets avec leur quantité
		for objet, nb := range joueur.Inventaire {
			fmt.Printf("x%d : %s\n", nb, objet)
		}
	}
}

// Utiliser un objet
func utiliserObjet(joueur, ennemie *Character_class, objet string) {
	qty, ok := joueur.Inventaire[objet]
	if !ok || qty == 0 {
		fmt.Println(objet, "n'est pas dans votre inventaire")
		return
	}

	// Exemple pour une Potion de soin
	if objet == "Potion de soin" {
		if joueur.Inventaire["Potion de soin"] > 0 {
			joueur.Pv += 50
			if joueur.Pv > joueur.MaxPv {
				joueur.Pv = joueur.MaxPv
			}
			joueur.Inventaire["Potion de soin"]-- // on enlève 1 potion
			fmt.Println("Vous utilisez une Potion de soin ! PV :", joueur.Pv, "/", joueur.MaxPv)
		} else {
			fmt.Println("❌ Vous n'avez pas de Potion de soin !")
		}
	}
	// Utiliser Potion de poison
	if objet == "Potion de poison" {
		if joueur.Inventaire["Potion de poison"] > 0 {
			ennemie.Pv -= 30
			if ennemie.Pv < 0 {
				ennemie.Pv = 0
			}
			joueur.Inventaire["Potion de poison"]-- // on enlève 1 potion
			fmt.Println("Vous utilisez une Potion de poison ! PV ennemi :", ennemie.Pv, "/", ennemie.MaxPv)
		} else {
			fmt.Println("❌ Vous n'avez pas de Potion de poison !")
		}
	}
}

// Utiliser un objet
func utiliserObjetTuto(joueur *Character_class, objet string) bool {
	qty, ok := joueur.Inventaire[objet]
	if !ok || qty == 0 {
		fmt.Println(objet, "n'est pas dans votre inventaire")
		return false
	}

	// Exemple pour une Potion de soin
	if objet == "Potion de soin" {
		if joueur.Inventaire["Potion de soin"] > 0 {
			joueur.Pv += 50
			if joueur.Pv > joueur.MaxPv {
				joueur.Pv = joueur.MaxPv
			}
			joueur.Inventaire["Potion de soin"]-- // on enlève 1 potion
			fmt.Println("Vous utilisez une Potion de soin ! PV :", joueur.Pv, "/", joueur.MaxPv)
			return true
		} else {
			fmt.Println("❌ Vous n'avez pas de Potion de soin !")
		}
	}
	return false
}
