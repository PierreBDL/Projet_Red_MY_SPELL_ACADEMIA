package MSA

type Character_class struct {
	Name       string
	Class      string
	Pv         int
	MaxPv      int
	Attaque    int
	Defence    int
	Gold       int
	Niveau     int
	Inventaire []string
}

func InitCharacter(nom string, classe int) Character_class {
	if classe == 1 {
		return Character_class{
			Name:       nom,
			Class:      "Sorcier",
			Pv:         100,
			MaxPv:      100,
			Attaque:    20,
			Defence:    5,
			Gold:       200,
			Niveau:     1,
			Inventaire: []string{"Potion de soin", "Potion de soin", "Potion de soin", "Potion de soin"},
		}
	}
	if classe == 2 {
		return Character_class{
			Name:       nom,
			Class:      "Alchimiste",
			Pv:         80,
			MaxPv:      80,
			Attaque:    30,
			Defence:    7,
			Gold:       200,
			Niveau:     1,
			Inventaire: []string{"Potion de soin", "Potion de soin", "Potion de soin", "Potion de soin"},
		}
	}

	if classe == 666 {
		// Choix par défaut pour éviter les bugs
		return Character_class{
			Name:       "Harry Poter",
			Class:      "Sang mélé",
			Pv:         150,
			MaxPv:      150,
			Attaque:    50,
			Defence:    25,
			Gold:       1000,
			Niveau:     1,
			Inventaire: []string{"Potion de soin", "Potion de soin", "Potion de soin", "Potion de soin"},
		}
	}

	// Choix par défaut pour éviter les bugs
	return Character_class{
		Name:    nom,
		Class:   "Inconnu",
		Pv:      50,
		MaxPv:   50,
		Attaque: 10,
		Defence: 5,
		Gold:    0,
		Niveau:  1,
	}
}

func InitEnnemi(monstre string) Character_class {
	if monstre == "gobelin" {
		return Character_class{
			Name:    "Gobelin",
			Class:   "Monstre",
			Pv:      50,
			MaxPv:   50,
			Attaque: 10,
			Defence: 5,
			Gold:    10, // Récompense
			Niveau:  1,
		}
	}
	if monstre == "boss" {
		return Character_class{
			Name:    "Boss",
			Class:   "Monstre",
			Pv:      150,
			MaxPv:   150,
			Attaque: 50,
			Defence: 5,
			Gold:    40, // Récompense
			Niveau:  3,
		}
	}
	return Character_class{
		Name:    "Orc",
		Class:   "Monstre",
		Pv:      70,
		MaxPv:   70,
		Attaque: 20,
		Defence: 5,
		Gold:    30, // Récompense
		Niveau:  2,
	}
}

func TutoEnnemi() Character_class {
	return Character_class{
		Name:    "Slime d'entraînement",
		Class:   "Il est consentant, je vous jure",
		Pv:      30,
		MaxPv:   30,
		Attaque: 50,
		Defence: 2,
		Gold:    0,
		Niveau:  1,
	}
}
