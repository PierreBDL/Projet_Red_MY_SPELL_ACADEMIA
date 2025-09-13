package MSA

import "fmt"

func WasDead(joueur *Character_class) bool {
	for i := 0; i < 40; i++ { // 20 lignes nettoyées
		fmt.Print("\033[A\033[2K")
	}
	fmt.Println(" ###### MY SPELL ACADEMIA ######")
	fmt.Println("💀 Game Over !\n")
	joueur.Pv = 0

	// Arrêt du jeu
	return false
}
