package MSA

import "fmt"

func Nettoyage(joueur *Character_class) {
	for i := 0; i < 40; i++ { // 20 lignes nettoyées
		fmt.Print("\033[A\033[2K")
	}
	fmt.Println(" ###### MY SPELL ACADEMIA ######")

	// Mise à jour des PVs max en fonction du niveau
	if joueur.Niveau > 1 {
		joueur.MaxPv = 100 + 10*(joueur.Niveau-1)
	}

	// Si le personnage est fait, on affiche sa fiche de stats
	println("\n   o  ◊		Nom :", joueur.Name, "		Argent:", joueur.Gold, "G\n  /|\\ |		Classe :", joueur.Class, "\n / | \\|		Niveau:", joueur.Niveau, "\n  / \\ |		PV:", joueur.Pv, "/", joueur.MaxPv)

}
