package main

import (
	"MSA"
	"fmt"
)

func main() {
	fmt.Println(" ###### MY SPELL ACADEMIA ######")

	// Création du personnage
	var name string
	fmt.Println("Quel est votre nom ? ")
	fmt.Scan(&name)

	// Affichage du personnage
	fmt.Println(MSA.InitCharacter(name))

	// Boucle infinie
	/*for 0 != 1 {

	}*/
}
