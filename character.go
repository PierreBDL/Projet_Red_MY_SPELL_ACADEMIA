package MSA

type Character_class struct {
	Name           string
	Class          string
	Pv             int
	MaxPv          int
	Attaque        int
	Defence        int
	Gold           int
	Niveau         int
	Inventaire     map[string]int
	upgradesBought int
	InventoryLimit int
	Sorts          [][]string // Nom, description, dégats
	Equipement     [][]string // Casque, Armure, Jambières
	Armes          []string   // Arme
	// Ennemis
	Type_attaque [][]string // attaques ennemis
}

func InitCharacter(nom string, classe int) Character_class {
	if classe == 1 {
		return Character_class{
			Name:           nom,
			Class:          "Sorcier",
			Pv:             100,
			MaxPv:          100,
			Attaque:        20,
			Defence:        5,
			Gold:           200,
			Niveau:         1,
			Inventaire:     map[string]int{"Potion de soin": 4, "Potion de poison": 2},
			upgradesBought: 0,
			InventoryLimit: 10,
			// Sorts
			Sorts: [][]string{{"🪄  Wingardium Leviosa !", "est envoyé loin et perd 10 PV.", "10"}, {"🔥 Incendio !", "est en feu et perd 15 PV.", "15"}},
			// Tenue + Arme
			Equipement: make([][]string, 3), // Casque, Armure, Jambières
			Armes:      make([]string, 1),   // Arme
		}
	}
	if classe == 2 {
		return Character_class{
			Name:           nom,
			Class:          "Alchimiste",
			Pv:             80,
			MaxPv:          80,
			Attaque:        30,
			Defence:        7,
			Gold:           200,
			Niveau:         1,
			Inventaire:     map[string]int{"Potion de soin": 4, "Potion de poison": 2},
			upgradesBought: 0,
			InventoryLimit: 10,
			// Sorts
			Sorts: [][]string{{"🪄  Wingardium Leviosa !", "est envoyé loin et perd 10 PV.", "10"}, {"🔥 Incendio !", "est en feu et perd 15 PV.", "15"}},
		}
	}

	if classe == 666 {
		// Choix par défaut pour éviter les bugs
		return Character_class{
			Name:           "Harry Potter",
			Class:          "Sang mélé",
			Pv:             150,
			MaxPv:          150,
			Attaque:        50,
			Defence:        25,
			Gold:           1000,
			Niveau:         1,
			Inventaire:     map[string]int{"Potion de soin": 4, "Potion de poison": 2},
			upgradesBought: 0,
			InventoryLimit: 10,
			// Sorts
			Sorts: [][]string{{"🪄  Wingardium Leviosa !", "est envoyé loin et perd 10 PV.", "10"}, {"🔥 Incendio !", "est en feu et perd 15 PV.", "15"}},
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
			Type_attaque: [][]string{{"Griffe", "Le gobelin vous griffe et vous inflige 10 points de dégâts.", "10"},
				{"Coup de massue", "Le gobelin vous assène un coup de massue et vous inflige 15 points de dégâts.", "15"},
				{"Coup de pied", "Le gobelin vous assène un coup de pied et vous inflige 20 points de dégâts.", "20"}},
		}
	}
	if monstre == "boss" {
		return Character_class{
			Name:    "Boss",
			Class:   "Monstre",
			Pv:      100,
			MaxPv:   100,
			Attaque: 30,
			Defence: 5,
			Gold:    40, // Récompense
			Niveau:  3,
			Type_attaque: [][]string{{"Griffe", "Le boss vous griffe et vous inflige 10 points de dégâts.", "10"}, {"Coup de massue", "Le boss vous assène un coup de massue et vous inflige 15 points de dégâts.", "15"},
				{"Coup de pied", "Le boss vous assène un coup de pied et vous inflige 20 points de dégâts.", "20"},
				{"Charge", "Le boss vous charge et vous inflige 25 points de dégâts.", "25"}},
		}
	}
	return Character_class{
		Name:    "Orc",
		Class:   "Monstre",
		Pv:      60,
		MaxPv:   60,
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
		Pv:      70,
		MaxPv:   70,
		Attaque: 50,
		Defence: 2,
		Gold:    0,
		Niveau:  1,
	}
}
