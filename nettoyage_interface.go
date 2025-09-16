package MSA

import "fmt"

func Nettoyage(joueur *Character_class) {
	for i := 0; i < 90; i++ { // 90 lignes nettoyées
		fmt.Print("\033[A\033[2K")
	}

	fmt.Println(" ###### MY SPELL ACADEMIA ######")

	// Mise à jour des PVs max en fonction du niveau
	if joueur.Niveau > 1 {
		// Calcul basé sur les PV de base selon la classe
		var basePv int
		if joueur.Class == "Sang mélé" {
			basePv = 150 // PV de base pour Harry Potter
		} else if joueur.Class == "Sorcier" {
			basePv = 100
		} else if joueur.Class == "Alchimiste" {
			basePv = 80
		} else {
			basePv = 50
		}

		// Calcul absolu : PV de base + bonus de niveau
		joueur.MaxPv = basePv + 10*(joueur.Niveau-1)

		// S'assurer que les PV actuels ne dépassent pas le maximum
		if joueur.Pv > joueur.MaxPv {
			joueur.Pv = joueur.MaxPv
		}
	}

	// Si le personnage est fait, on affiche sa fiche de stats
	fmt.Printf(`
   o  ◊        Nom : %s      	Argent 💰: %d G
  /|\ |        Classe : %s
 / | \|        Niveau: %d
  / \ |        PVs ❤️: %d / %d
`,
		joueur.Name, joueur.Gold, joueur.Class, joueur.Niveau, joueur.Pv, joueur.MaxPv)
	fmt.Println()
}
