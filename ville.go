package MSA

import (
	"fmt"
)

func Entree_ville(joueur *Character_class) {
	Nettoyage(joueur)

	fmt.Println("╔════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    🏰 VILLE DE MUSUTAFU 🏰                     ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════╣")
	fmt.Println("║                                                                ║")
	fmt.Println("║    🏪 1] 🤝 Marché          🏠 3] 🛌 Auberge (10G/nuit)        ║")
	fmt.Println("║    ⚒️  2] ⚒️  Forge           🎯 4] 🎯 Terrain d'entraînement    ║")
	fmt.Println("║               5] 🚪🚶 Quitter la ville                         ║")
	fmt.Println("║                                                                ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════╝")
	fmt.Print("➤ Votre choix : ")

	// Le reste des tours
	choix_ville := 0
	fmt.Scan(&choix_ville)
	fmt.Print("")

	// Annalyse du choix
	for choix_ville != 1 && choix_ville != 2 && choix_ville != 3 && choix_ville != 4 && choix_ville != 5 {
		fmt.Print("Choix invalide. Veuillez recommencer ")
		fmt.Print("\033[A\033[2K") // Remonte et efface les dernières lignes
		fmt.Print("Quel est votre choix ? ")
		fmt.Scan(&choix_ville)
		fmt.Print("")
	}
	switch choix_ville {
	case 1:
		Marchand(joueur)
	case 2:
		Nettoyage(joueur)
		Forge(joueur)
	case 3:
		if joueur.Gold >= 10 {
			ancien_pv := joueur.Pv
			joueur.Pv = joueur.Pv + 10
			// Eviter débordement
			if joueur.Pv > joueur.MaxPv {
				joueur.Pv = joueur.MaxPv
			}
			pv_soignes := joueur.Pv - ancien_pv
			fmt.Println("\nVous passez une bonne nuit de sommeil,", pv_soignes, "PVs soignés !")
			// Prix
			joueur.Gold -= 10
			fmt.Println("Vous avez payez 10G, il vous reste", joueur.Gold, "G")
		} else {
			fmt.Println("Vous n'avez pas assez d'or, vous avez", joueur.Gold, "G")
		}

	case 4:
		Nettoyage(joueur)
		Entrainement(joueur)
	case 5:
		return
	default:
		break
	}
}
