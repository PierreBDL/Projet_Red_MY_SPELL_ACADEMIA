package main

import (
	"fmt"
	"MSA"
)

func main() {
	fmt.Println(" ###### MY SPELL ACADEMIA ######")

	// Création du personnage
	var name string
	fmt.Println("Quel est votre nom ? ")
	fmt.Scan(&name)

	// Affichage du personnage
	fmt.Println(MSA.initCharacter(name))

	// Boucle infinie
	/*for 0 != 1 {

	}*/
}
