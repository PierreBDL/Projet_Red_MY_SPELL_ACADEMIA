package MSA

import "fmt"

func Nettoyage(joueur *Character_class) {
	for i := 0; i < 40; i++ { // 20 lignes nettoyées
		fmt.Print("\033[A\033[2K")
	}
	fmt.Println(" ###### MY SPELL ACADEMIA ######")

	// Si le personnage est fait, on affiche sa fiche de stats
	println("\n  o ◊		Nom :", joueur.Name, "\n /|\\|		Classe :", joueur.Class, "\n / \\		PV:", joueur.Pv, "/", joueur.MaxPv)

}
