package main

import (
	"fmt"
)

type Merchant struct {
	name    string
	gold    int
	potions map[string]int
	items   map[string]int
	upgrades map[string]int
}

type Player struct {
	inventory      map[string]int
	gold           int
	inventoryLimit int
	upgradesBought int
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
		upgrades: map[string]int{
			"augmentation_inventaire": 100, // coût pour augmenter l'inventaire
		},
	}
}

func NewPlayer(gold int) *Player {
	return &Player{
		inventory:      make(map[string]int),
		gold:           gold,
		inventoryLimit: 10, // limite initiale de l'inventaire
		upgradesBought: 0,
	}
}

// Retourne la liste des objets et affiche avec un index
func (m *Merchant) DisplayCategory(choice int, player *Player) []string {
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
	case 3:
	fmt.Println("Catégorie : Améliorations")
	if player.upgradesBought < 3 {
		for upgrade, price := range m.upgrades {
			// Suppression du ✅ ici
			fmt.Printf("%d - %s : %d or\n", 1, upgrade, price)
			fmt.Println("💡 Astuce : tu peux augmenter ton inventaire de 1. Max 3 fois.")
			keys = append(keys, upgrade)
		}
	} else {
		fmt.Println("❌ Tu as déjà acheté toutes les améliorations disponibles.")
	}
	default:
		fmt.Println("Choix invalide")
	}
	return keys
}

func (p *Player) Buy(m *Merchant, category int, itemName string, quantity int) {
	if category == 3 {
		if p.upgradesBought >= 3 {
			fmt.Println("❌ Tu ne peux plus acheter cette amélioration.")
			return
		}
		price := m.upgrades[itemName]
		if p.gold < price {
			fmt.Println("❌ Pas assez d'or pour cette amélioration.")
			return
		}
		p.gold -= price
		p.inventoryLimit++
		p.upgradesBought++
		fmt.Printf("✅ Amélioration achetée ! Nouvelle limite d'inventaire : %d. Or restant : %d\n", p.inventoryLimit, p.gold)
		return
	}

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
	fmt.Printf(" Vous avez acheté %d x %s. Or restant : %d\n", quantity, itemName, p.gold)
}

func main() {
	merchant := NewMerchant("Mandragor", 1000)
	player := NewPlayer(2000)

	// Présentation
	fmt.Println("👋 Bonjour, je suis", merchant.name, "!")
	fmt.Println("Que puis-je faire pour toi ?")
	fmt.Println("1 - Potions")
	fmt.Println("2 - Baguette Magique")
	fmt.Println("3 - Améliorations")

	var choice int
	fmt.Scanln(&choice)

	// On récupère les objets de la catégorie
	items := merchant.DisplayCategory(choice, player)
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

	quantity := 1
	if choice != 3 { // pour les améliorations, pas besoin de quantité
		fmt.Println("Combien veux-tu en acheter ? (max 5)")
		fmt.Scanln(&quantity)
	}

	// Achat
	player.Buy(merchant, choice, itemName, quantity)

	// Inventaire
	fmt.Println("📦 Inventaire du joueur :")
	for item, q := range player.inventory {
		fmt.Printf("- %s x%d\n", item, q)
	}
	fmt.Printf("💼 Limite d'inventaire : %d\n", player.inventoryLimit)
	fmt.Printf("💰 Or restant : %d\n", player.gold)
}
