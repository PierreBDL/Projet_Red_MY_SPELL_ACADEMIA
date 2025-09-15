package main

import (
	"fmt"
)

type Player struct {
	name string
}

type Enemy struct {
	name string
	hp   int
}

func NewPlayer(name string) *Player {
	return &Player{name: name}
}

func NewEnemy(name string, hp int) *Enemy {
	return &Enemy{name: name, hp: hp}
}

// Méthode pour lancer un sort
func (p *Player) CastSpell(e *Enemy) {
	fmt.Println("Choisissez un sort :")
	fmt.Println("1️⃣ Wingardium Leviosa (envoie loin le gobelin et lui inflige 10 PV)")
	fmt.Println("2️⃣ Incendio (met le gobelin en feu et lui enlève la moitié de ses PV)")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		e.hp -= 10
		fmt.Printf("🪄 Wingardium Leviosa ! %s est envoyé loin et perd 10 PV. PV restants : %d\n", e.name, e.hp)
	case 2:
		loss := e.hp / 2
		e.hp -= loss
		fmt.Printf("🔥 Incendio ! %s est en feu et perd %d PV. PV restants : %d\n", e.name, loss, e.hp)
	default:
		fmt.Println("❌ Sort inconnu. Vous ratez votre tour !")
	}

	if e.hp <= 0 {
		fmt.Printf("💀 %s est vaincu !\n", e.name)
	}
}

func main() {
	player := NewPlayer("Harry")
	enemy := NewEnemy("Gobelin", 50)

	fmt.Printf("⚔️ Combat contre %s ! PV : %d\n", enemy.name, enemy.hp)

	// Le joueur lance un sort
	player.CastSpell(enemy)
}
