package MSA

import (
	"fmt"
)

type Merchant struct {
	name     string
	gold     int
	potions  map[string]int
	items    map[string]int
	upgrades map[string]int
	armors   map[string]int
}

// Création d’un marchand
func NewMerchant(name string, gold int) *Merchant {
	return &Merchant{
		name: name,
		gold: gold,
		potions: map[string]int{
			"Potion de soin":   200, // ✅ Avec espaces et majuscules
			"Potion de poison": 200, // ✅ Pareil que dans accessInventory.go
		},
		items: map[string]int{
			"bois_de_sureau":        500,
			"plume_de_phenix":       600,
			"baguette_surpuissante": 1000,
		},
		upgrades: map[string]int{
			"augmentation_inventaire": 100,
		},
		armors: map[string]int{
			"Casque en fer":    300,
			"Armure en fer":    400,
			"Jambières en fer": 300,
		},
	}
}

// Affichage des catégories
func (m *Merchant) DisplayCategory(choice int, joueur *Character_class) []string {
	var keys []string
	switch choice {
	case 1:
		fmt.Println("Catégorie : Potions")
		i := 1
		for potion, price := range m.potions {
			// Affichage plus joli pour l'interface
			displayName := potion
			if potion == "Potion de soin" {
				displayName = "🧪 Potion de soin"
			} else if potion == "Potion de poison" {
				displayName = "☠️ Potion de poison"
			}
			fmt.Printf("%d - %s : %d gold\n", i, displayName, price)
			keys = append(keys, potion)
			i++
		}
	case 2:
		fmt.Println("Catégorie : Baguette Magique")
		i := 1
		for item, price := range m.items {
			fmt.Printf("%d - %s : %d gold\n", i, item, price)
			keys = append(keys, item)
			i++
		}
	case 3:
		fmt.Println("Catégorie : Améliorations")
		if joueur.upgradesBought < 3 {
			for upgrade, price := range m.upgrades {
				fmt.Printf("%d - %s : %d gold\n", 1, upgrade, price)
				fmt.Println("💡 Astuce : tu peux augmenter ton inventaire de 1. Max 3 fois.")
				keys = append(keys, upgrade)
			}
		} else {
			fmt.Println("❌ Tu as déjà acheté toutes les améliorations disponibles.")
		}
	case 4:
		fmt.Println("Catégorie : Armures (matériaux)")
		i := 1
		for armor, price := range m.armors {
			fmt.Printf("%d - %s : %d gold\n", i, armor, price)
			keys = append(keys, armor)
			i++
		}
	default:
		fmt.Println("Choix invalide")
	}
	return keys
}

// Achat
func (joueur *Character_class) Buy(m *Merchant, category int, itemName string, quantity int) {
	switch category {
	case 1: // Potions
		if quantity < 1 || quantity > 5 {
			fmt.Println("❌ Vous ne pouvez acheter qu'entre 1 et 5 exemplaires.")
			return
		}
		val, ok := m.potions[itemName]
		if !ok {
			fmt.Println("❌ Potion introuvable.")
			return
		}
		price := val * quantity
		if joueur.Gold < price {
			fmt.Println("❌ Pas assez de gold pour cet achat.")
			return
		}
		joueur.Gold -= price
		joueur.Inventaire[itemName] += quantity
		fmt.Printf("✅ Vous avez acheté %d x %s. Gold restant : %d\n", quantity, itemName, joueur.Gold)

	case 2: // Items (baguette magique)
		if quantity < 1 || quantity > 5 {
			fmt.Println("❌ Vous ne pouvez acheter qu'entre 1 et 5 exemplaires.")
			return
		}
		val, ok := m.items[itemName]
		if !ok {
			fmt.Println("❌ Objet introuvable.")
			return
		}
		price := val * quantity
		if joueur.Gold < price {
			fmt.Println("❌ Pas assez de gold pour cet achat.")
			return
		}
		joueur.Gold -= price
		joueur.Inventaire[itemName] += quantity
		fmt.Printf("✅ Vous avez acheté %d x %s. Gold restant : %d\n", quantity, itemName, joueur.Gold)

	case 3: // Upgrades
		if joueur.upgradesBought >= 3 {
			fmt.Println("❌ Tu ne peux plus acheter cette amélioration.")
			return
		}
		price := m.upgrades[itemName]
		if joueur.Gold < price {
			fmt.Println("❌ Pas assez de gold pour cette amélioration.")
			return
		}
		joueur.Gold -= price
		joueur.InventoryLimit++
		joueur.upgradesBought++
		fmt.Printf("✅ Amélioration achetée ! Nouvelle limite d'inventaire : %d. Gold restant : %d\n",
			joueur.InventoryLimit, joueur.Gold)

	case 4: // Armures
		if quantity < 1 || quantity > 5 {
			fmt.Println("❌ Vous ne pouvez acheter qu'entre 1 et 5 exemplaires.")
			return
		}
		val, ok := m.armors[itemName]
		if !ok {
			fmt.Println("❌ Matériau introuvable.")
			return
		}
		price := val * quantity
		if joueur.Gold < price {
			fmt.Println("❌ Pas assez de gold pour cet achat.")
			return
		}
		joueur.Gold -= price
		joueur.Inventaire[itemName] += quantity

		// Défense et placement dans Equipement
		defenseValue := 0
		slot := -1
		switch itemName {
		case "Casque en fer":
			defenseValue = 10
			slot = 0
		case "Armure en fer":
			defenseValue = 15
			slot = 1
		case "Jambières en fer":
			defenseValue = 10
			slot = 2
		}

		if slot != -1 {
			joueur.Equipement[slot] = []string{itemName, fmt.Sprintf("%d", defenseValue)}
			joueur.Defence += defenseValue
		}

		fmt.Printf("✅ Vous avez acheté %d x %s. Gold restant : %d\n", quantity, itemName, joueur.Gold)

	default:
		fmt.Println("❌ Catégorie invalide.")
	}
}

func Marchand(joueur *Character_class) {
	marchand := NewMerchant("Mandragor", 1000) // ton marchand

	// Présentation
	fmt.Println("\n👋 Bonjour, je suis", marchand.name, "!")
	fmt.Println("Que puis-je faire pour toi ?")
	fmt.Println("1 - Potions")
	fmt.Println("2 - Baguette Magique")
	fmt.Println("3 - Améliorations")
	fmt.Println("4 - Armures")

	var choix int
	fmt.Scanln(&choix)

	// Récupère les objets disponibles dans la catégorie
	items := marchand.DisplayCategory(choix, joueur)
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

	// Quantité à acheter
	quantity := 1
	if choix != 3 { // les améliorations n'ont pas de quantité
		fmt.Println("Combien veux-tu en acheter ? (max 5)")
		fmt.Scanln(&quantity)
	}

	// Achat
	joueur.Buy(marchand, choix, itemName, quantity)

	// Inventaire
	fmt.Println("\n📦 Inventaire du joueur :")
	for item, q := range joueur.Inventaire {
		fmt.Printf("- %s x%d\n", item, q)
	}
	fmt.Printf("💼 Limite d'inventaire : %d\n", joueur.InventoryLimit)
	fmt.Printf("💰 Gold restant : %d\n", joueur.Gold)
}
