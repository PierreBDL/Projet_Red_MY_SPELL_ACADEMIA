package MSA

import "fmt"

func Nettoyage(joueur *Character_class) {
	for i := 0; i < 90; i++ { // 50 lignes nettoyées
		fmt.Print("\033[A\033[2K")
	}

	fmt.Println(" ###### MY SPELL ACADEMIA ######")

	// Mise à jour des PVs max en fonction du niveau
	if joueur.Niveau > 1 {
		joueur.MaxPv = 100 + 10*(joueur.Niveau-1)
	}

	// Si le personnage est fait, on affiche sa fiche de stats
	fmt.Printf(`
   o  ◊        Nom : %s      	Argent 💰: %d G
  /|\ |        Classe : %s
 / | \|        Niveau: %d
  / \ |        PVs ❤️: %d / %d
`,
		joueur.Name, joueur.Gold, joueur.Class, joueur.Niveau, joueur.Pv, joueur.MaxPv)
	fmt.Println("\n")
}
