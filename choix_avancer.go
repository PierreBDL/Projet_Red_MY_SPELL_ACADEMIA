package MSA

import (
	"math/rand"
	"time"
)

func Avancer() string {
	// Création du générateur d'aléatoire
	rand.Seed(time.Now().UnixNano())

	// Probabilité de tomber sur un ennemis
	arrivee_ennemie_prob := 0.7

	// Apparition ou non d'un ennemi
	if rand.Float64() < arrivee_ennemie_prob {
		return "gobelin"
	} else {
		return "boss"
	}
}
