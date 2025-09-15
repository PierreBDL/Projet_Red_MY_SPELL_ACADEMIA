package main

import (
	"fmt"
)

type Merchant struct {
	name    string
	gold    int
	potions map[string]int
	items   map[string]int
}

type Player struct {
	inventory map[string]int
	gold      int
}

func NewMerchant(name string, gold int) *Merchant {
	return &Merchant{
		name: name,
		gold: gold,
		potions: map[string]int{
			"potion_de_soin":   200,
			"potion_de_poison": 200,
		},
		items: map[string]int{
			"bois_de_sureau":        500,
			"plume_de_phenix":       600,
			"baguette_surpuissante": 1000,
		},
	}
}

func NewPlayer(gold int) *Player {
	return &Player{
		inventory: make(map[string]int),
		gold:      gold,
	}
}

// Retourne la liste des objets et affiche avec un index
func (m *Merchant) DisplayCategory(choice int) []string {
	var keys []string
	switch choice {
	case 1:
		fmt.Println("Catégorie : Potions")
		i := 1
		for potion, price := range m.potions {
			fmt.Printf("%d - %s : %d or\n", i, potion, price)
			keys = append(keys, potion)
			i++
		}
	case 2:
		fmt.Println("Catégorie : Baguette Magique")
		i := 1
		for item, price := range m.items {
			fmt.Printf("%d - %s : %d or\n", i, item, price)
			keys = append(keys, item)
			i++
		}
	default:
		fmt.Println("Choix invalide")
	}
	return keys
}

func (p *Player) Buy(m *Merchant, category int, itemName string, quantity int) {
	if quantity < 1 || quantity > 5 {
		fmt.Println("❌ Vous ne pouvez acheter qu'entre 1 et 5 exemplaires.")
		return
	}

	var price int
	if category == 1 {
		val, ok := m.potions[itemName]
		if !ok {
			fmt.Println("❌ Potion introuvable.")
			return
		}
		price = val * quantity
	} else if category == 2 {
		val, ok := m.items[itemName]
		if !ok {
			fmt.Println("❌ Objet introuvable.")
			return
		}
		price = val * quantity
	} else {
		fmt.Println("❌ Catégorie invalide.")
		return
	}

	if p.gold < price {
		fmt.Println("❌ Pas assez d'or pour cet achat.")
		return
	}

	// Déduction or + ajout inventaire
	p.gold -= price
	p.inventory[itemName] += quantity
	fmt.Printf("✅ Vous avez acheté %d x %s. Or restant : %d\n", quantity, itemName, p.gold)
}

func main() {
	merchant := NewMerchant("Mandragor", 1000)
	player := NewPlayer(2000)

	// Présentation
	fmt.Println("👋 Bonjour, je suis", merchant.name, "!")
	fmt.Println("Que puis-je faire pour toi ?")
	fmt.Println("1 - Potions")
	fmt.Println("2 - Baguette Magique")

	var choice int
	fmt.Scanln(&choice)

	// On récupère les objets de la catégorie
	items := merchant.DisplayCategory(choice)
	if len(items) == 0 {
		return
	}

	// Choix de l’item par numéro
	fmt.Println("Choisis un objet en entrant son numéro :")
	var itemIndex int
	fmt.Scanln(&itemIndex)

	if itemIndex < 1 || itemIndex > len(items) {
		fmt.Println("❌ Choix invalide.")
		return
	}
	itemName := items[itemIndex-1]

	// Quantité
	fmt.Println("Combien veux-tu en acheter ? (max 5)")
	var qty int
	fmt.Scanln(&qty)

	// Achat
	player.Buy(merchant, choice, itemName, qty)

	// Inventaire
	fmt.Println("📦 Inventaire du joueur :")
	for item, q := range player.inventory {
		fmt.Printf("- %s x%d\n", item, q)
	}
}
