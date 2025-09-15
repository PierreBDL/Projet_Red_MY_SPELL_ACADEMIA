package main

import (
	"fmt"
)

type Player struct {
	inventory map[string]int
	maxItems  int
}

func NewPlayer() *Player {
	return &Player{
		inventory: map[string]int{},
		maxItems:  10,
	}
}

// Fonction pour compter le nombre total d'objets dans l'inventaire
func (p *Player) TotalItems() int {
	total := 0
	for _, qty := range p.inventory {
		total += qty
	}
	return total
}

// Fonction pour ajouter un objet avec limite
func (p *Player) AddItem(item string, qty int) {
	if p.TotalItems()+qty > p.maxItems {
		fmt.Println("❌ Inventaire plein ! Vous devez jeter des objets avant d'ajouter", item)
		p.ShowInventory()
		return
	}

	p.inventory[item] += qty
	fmt.Printf("✅ %d %s ajouté(s) à votre inventaire.\n", qty, item)
}

// Affichage inventaire avec icônes
func (p *Player) ShowInventory() {
	fmt.Println("📦 Inventaire du joueur :")
	for item, qty := range p.inventory {
		icon := ""
		switch item {
		case "potion_de_poison":
			icon = "🧪"
		case "bois_de_sureau":
			icon = "🌳"
		case "plume_de_phenix":
			icon = "🪶"
		default:
			icon = "❔"
		}
		fmt.Printf("%s %s x%d\n", icon, item, qty)
	}
	fmt.Printf("Total objets : %d/%d\n", p.TotalItems(), p.maxItems)
}

func main() {
	player := NewPlayer()

	// Exemple d'ajout d'objets
	player.AddItem("potion_de_poison", 3)
	player.AddItem("bois_de_sureau", 5)
	player.AddItem("plume_de_phenix", 3) // Devrait bloquer car dépasse 10

	fmt.Println()
	player.ShowInventory()
}
