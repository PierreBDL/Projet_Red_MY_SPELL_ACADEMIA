package MSA

import (
	"fmt"
)

func Entree_ville(joueur *Character_class) {
	// Le reste des tours
	choix_ville := 0
	fmt.Println("\nVous entrez dans la ville de Musutafu !")
	fmt.Println("1] Marché")
	fmt.Println("2] Forge")
	fmt.Println("3] Quitter la ville")
	fmt.Print("Quel est votre choix ? ")
	fmt.Scan(&choix_ville)

	// Annalyse du choix
	for choix_ville != 1 && choix_ville != 2 && choix_ville != 3 {
		fmt.Print("Choix invalide. Veuillez recommencer ")
		fmt.Print("\033[A\033[2K") // Remonte et efface les dernières lignes
		fmt.Print("Quel est votre choix ? ")
		fmt.Scan(&choix_ville)
	}
	switch choix_ville {
	case 1:
		Nettoyage(joueur)
		Ville(joueur)
	case 2:
		Nettoyage(joueur)
		Forge(joueur)
	case 3:
		return
	default:
		break
	}
}
