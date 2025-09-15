package main

import (
	"fmt"
	"time"
)

type Player struct {
	inventory map[string]int
}

type Enemy struct {
	name string
	hp   int
}

func NewPlayer() *Player {
	return &Player{
		inventory: map[string]int{
			"potion_de_poison": 1, // Le joueur commence avec 1 potion
		},
	}
}

func NewEnemy(name string, hp int) *Enemy {
	return &Enemy{name: name, hp: hp}
}

func (p *Player) UsePoisonPotion(e *Enemy) {
	if p.inventory["potion_de_poison"] > 0 {
		fmt.Println("🧪 Vous utilisez une potion de poison sur", e.name, "!")
		p.inventory["potion_de_poison"] -= 1

		// Lancement d’un effet poison sur 10 secondes
		go func() {
			for i := 1; i <= 10; i++ {
				time.Sleep(1 * time.Second) // Attente 1 seconde
				damage := 5
				e.hp -= damage
				fmt.Printf("☠️ (%ds) %s subit %d dégâts de poison. PV restants : %d\n", i, e.name, damage, e.hp)
				if e.hp <= 0 {
					fmt.Printf("💀 %s est vaincu par le poison !\n", e.name)
					return
				}
			}
			fmt.Printf("✅ L'effet de poison sur %s s'est dissipé.\n", e.name)
		}()
	} else {
		fmt.Println("❌ Vous n'avez plus de potion de poison.")
	}
}

// Fonction pour afficher l’inventaire avec icônes
func (p *Player) ShowInventory() {
	fmt.Println("📦 Inventaire du joueur :")
	for item, qty := range p.inventory {
		icon := ""
		switch item {
		case "potion_de_poison":
			icon = "🧪"
		default:
			icon = "❔"
		}
		fmt.Printf("%s %s x%d\n", icon, item, qty)
	}
}

func main() {
	// Création joueur et ennemi
	player := NewPlayer()
	enemy := NewEnemy("Gobelin", 50)

	// Affichage inventaire
	player.ShowInventory()

	// Action
	fmt.Println("Voulez-vous utiliser une potion de poison sur l'ennemi ? (oui/non)")
	var choice string
	fmt.Scanln(&choice)

	if choice == "oui" {
		player.UsePoisonPotion(enemy)
	} else {
		fmt.Println("🔚 Vous gardez votre potion pour plus tard.")
	}

	// Attente pour laisser le poison agir
	time.Sleep(12 * time.Second)

	// Inventaire final
	fmt.Println()
	player.ShowInventory()
}
