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
