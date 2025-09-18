package MSA

import (
	"math/rand"
	"time"
)

func Ennemis_type_attaque(ennemi string) int {
	// Création du générateur d'aléatoire
	rand.Seed(time.Now().UnixNano())

	// Attaque aléatoire
	if rand.Float64() < 0.5 {
		return 0
	} else if rand.Float64() > 0.5 && rand.Float64() < 0.8 {
		return 1
	} else if rand.Float64() > 0.8 && rand.Float64() < 1 {
		return 2
	}

	return 3
}
