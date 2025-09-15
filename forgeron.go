package main

import (
	"fmt"
)

type Player struct {
	inventory map[string]int
	gold      int
}

func NewPlayer(gold int) *Player {
	return &Player{
		inventory: make(map[string]int),
		gold:      gold,
	}
}

type Blacksmith struct {
	name string
}

func NewBlacksmith(name string) *Blacksmith {
	return &Blacksmith{name: name}
}

func (b *Blacksmith) CraftWand(p *Player) {
	fmt.Println("🔨 Bienvenue dans la forge d'", b.name, "!")
	fmt.Println("Pour crafter la baguette de sureau, il te faut :")
	fmt.Println("- 1 bois de sureau")
	fmt.Println("- 1 plume de phénix")
	fmt.Println("- 400 or")

	// Vérification ressources
	if p.inventory["bois_de_sureau"] < 1 {
		fmt.Println("❌ Il te manque un bois de sureau.")
		return
	}
	if p.inventory["plume_de_phenix"] < 1 {
		fmt.Println("❌ Il te manque une plume de phénix.")
		return
	}
	if p.gold < 400 {
		fmt.Println("❌ Tu n'as pas assez d'or.")
		return
	}

	// Retrait des ressources
	p.inventory["bois_de_sureau"] -= 1
	p.inventory["plume_de_phenix"] -= 1
	p.gold -= 400

	// Ajout baguette
	p.inventory["baguette_de_sureau"] += 1
	fmt.Println("✅ La baguette de sureau a été forgée et ajoutée à ton inventaire !")
}

func main() {
	// Joueur avec ressources de base
	player := NewPlayer(1000)
	player.inventory["bois_de_sureau"] = 1
	player.inventory["plume_de_phenix"] = 1

	// Forgeron
	olivenders := NewBlacksmith("Olivenders")

	// Dialogue
	fmt.Println("👋 Bonjour aventurier ! Tu es dans la forge d’Olivenders.")
	fmt.Println("Veux-tu forger la baguette de sureau ? (oui/non)")

	var response string
	fmt.Scanln(&response)

	if response == "oui" {
		olivenders.CraftWand(player)
	} else {
		fmt.Println("🔚 Peut-être une prochaine fois.")
	}

	// Inventaire final
	fmt.Println("\n📦 Inventaire du joueur :")
	for item, q := range player.inventory {
		fmt.Printf("- %s x%d\n", item, q)
	}
	fmt.Printf("💰 Or restant : %d\n", player.gold)
}
