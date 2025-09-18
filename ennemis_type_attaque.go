package MSA

import (
	"math/rand"
	"time"
)

func Ennemis_type_attaque() int {
	// Création du générateur d'aléatoire
	rand.Seed(time.Now().UnixNano())
	val := rand.Float64()
	// Attaque aléatoire
	if val <= 0.5 {
		return 0
	} else if val <= 0.8 {
		return 1
	} else if val <= 1 {
		return 2
	}

	return 3
}
