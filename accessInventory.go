package MSA

import "fmt"

// Voir inventaire ou autre action
func accessInventory(joueur *Character_class, action int) {
	if action == 1 {
		fmt.Println("Dans votre inventaire :")
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
func utiliserObjet(joueur *Character_class, objet string) {
	qty, ok := joueur.Inventaire[objet]
	if !ok || qty == 0 {
		fmt.Println(objet, "n'est pas dans votre inventaire")
		return
	}

	// Exemple pour une Potion de soin
	if objet == "Potion de soin" {
		joueur.Pv += 50
		if joueur.Pv > joueur.MaxPv {
			joueur.Pv = joueur.MaxPv
		}
		// Consommer une potion
		joueur.Inventaire[objet]--
		if joueur.Inventaire[objet] == 0 {
			delete(joueur.Inventaire, objet) // supprime la clé si plus d'objet
		}
		fmt.Println("Vous utilisez une Potion de soin ! PV :", joueur.Pv, "/", joueur.MaxPv)
	}
}
