package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player struct {
	inventory      map[string]int
	maxItems       int
	upgradesBought int
}

func NewPlayer() *Player {
	return &Player{
		inventory:      map[string]int{},
		maxItems:       10,
		upgradesBought: 0,
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

// Fonction pour afficher inventaire avec icônes
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

// Fonction pour améliorer la limite d'inventaire
func (p *Player) UpgradeInventory() {
	if p.upgradesBought >= 3 {
		fmt.Println("❌ Vous avez déjà atteint la limite maximale d'améliorations.")
		return
	}
	p.maxItems++
	p.upgradesBought++
	fmt.Printf("✅ Limite d'inventaire augmentée ! Nouvelle limite : %d\n", p.maxItems)
}

// Fonction pour supprimer un objet existant
func (p *Player) RemoveItem(item string, qty int) {
	if currentQty, ok := p.inventory[item]; ok {
		if qty >= currentQty {
			delete(p.inventory, item)
		} else {
			p.inventory[item] -= qty
		}
		fmt.Printf("🗑️ %d %s supprimé(s) de l'inventaire.\n", qty, item)
	} else {
		fmt.Println("❌ Objet introuvable dans l'inventaire.")
	}
}

// Fonction pour ajouter un objet avec gestion de l'espace
func (p *Player) AddItem(item string, qty int) {
	for p.TotalItems()+qty > p.maxItems {
		fmt.Printf("❌ Inventaire plein ! Vous essayez d'ajouter %d %s.\n", qty, item)
		p.ShowInventory()

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Voulez-vous jeter un objet pour faire de la place ? (oui/non) : ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))

		if answer != "oui" {
			fmt.Println("❌ Ajout annulé.")
			return
		}

		fmt.Print("Quel objet voulez-vous jeter ? ")
		removeItem, _ := reader.ReadString('\n')
		removeItem = strings.TrimSpace(removeItem)

		fmt.Print("Combien en voulez-vous jeter ? ")
		var removeQty int
		fmt.Scanln(&removeQty)

		p.RemoveItem(removeItem, removeQty)
	}

	p.inventory[item] += qty
	fmt.Printf("✅ %d %s ajouté(s) à votre inventaire.\n", qty, item)
}

func main() {
	player := NewPlayer()

	// Exemple d'ajout d'objets
	player.AddItem("potion_de_poison", 3)
	player.AddItem("bois_de_sureau", 5)
	player.AddItem("plume_de_phenix", 3) // Va demander si on veut jeter quelque chose

	fmt.Println()
	player.ShowInventory()

	// Exemple d'amélioration de l'inventaire
	fmt.Println("\n🎁 Achat d'une amélioration d'inventaire :")
	player.UpgradeInventory()
	player.UpgradeInventory()
	player.UpgradeInventory()
	player.UpgradeInventory() // Devrait bloquer, max 3 fois

	fmt.Println()
	player.ShowInventory()
}
