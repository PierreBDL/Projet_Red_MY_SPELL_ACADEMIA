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

// Craft baguette de sureau
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

// Craft Robe de Bataille des Arcanes
func (b *Blacksmith) CraftRobe(p *Player) {
	fmt.Println("🔨 Pour crafter la Robe de Bataille des Arcanes, il te faut :")
	fmt.Println("- 1 Tissu de lin enchanté")
	fmt.Println("- 1 Fil d'argent runique")
	fmt.Println("- 1 Poudre de pierre lunaire")
	fmt.Println("- 1 Perle d'âme")
	fmt.Println("- 1 Plume de phénix")
	fmt.Println("- 200 or")

	// Vérification des ressources
	materials := []string{"tissu_de_lin_enchante", "fil_d_argent_runique", "poudre_de_pierre_lunaire", "perle_d_ame", "plume_de_phenix"}
	for _, mat := range materials {
		if p.inventory[mat] < 1 {
			fmt.Printf("❌ Il te manque %s.\n", mat)
			return
		}
	}
	if p.gold < 200 {
		fmt.Println("❌ Tu n'as pas assez d'or.")
		return
	}

	// Retrait des matériaux
	for _, mat := range materials {
		p.inventory[mat] -= 1
	}
	p.gold -= 200

	// Ajout robe
	p.inventory["robe_de_bataille_des_arcanes"] += 1
	fmt.Println("✅ La Robe de Bataille des Arcanes a été fabriquée et ajoutée à ton inventaire !")
}

func main() {
	// Joueur avec ressources de base
	player := NewPlayer(1000)
	player.inventory["bois_de_sureau"] = 1
	player.inventory["plume_de_phenix"] = 2
	player.inventory["tissu_de_lin_enchante"] = 1
	player.inventory["fil_d_argent_runique"] = 1
	player.inventory["poudre_de_pierre_lunaire"] = 1
	player.inventory["perle_d_ame"] = 1

	// Forgeron
	olivenders := NewBlacksmith("Olivenders")

	// Dialogue
	fmt.Println("👋 Bonjour aventurier ! Tu es dans la forge d’Olivenders.")
	fmt.Println("Que veux-tu fabriquer ?")
	fmt.Println("1 - Baguette de sureau")
	fmt.Println("2 - Robe de Bataille des Arcanes")

	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		olivenders.CraftWand(player)
	} else if choice == 2 {
		olivenders.CraftRobe(player)
	} else {
		fmt.Println("🔚 Choix invalide, peut-être une prochaine fois.")
	}

	// Inventaire final
	fmt.Println("\n📦 Inventaire du joueur :")
	for item, q := range player.inventory {
		fmt.Printf("- %s x%d\n", item, q)
	}
	fmt.Printf("💰 Or restant : %d\n", player.gold)
}
