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
	Mana           int
	MaxMana        int
	// Ennemis
	Type_attaque [][]string // attaques ennemis
	// Poison
	Empoisonne  bool
	Tour_poison int
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
			Sorts: [][]string{
				{"🪄  Wingardium Leviosa !", "est envoyé loin et perd 10 PV.", "10"},
				{"🔥 Incendio !", "est en feu et perd 15 PV.", "15"},
				{"⚡ Stupefix !", "est étourdi et perd 20 PV.", "20"},
			},
			// Tenue + Arme
			Equipement: make([][]string, 3), // Casque, Armure, Jambières
			Armes:      make([]string, 1),   // Arme
			// Mana
			Mana:    50,
			MaxMana: 50,
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
			Gold:           300,
			Niveau:         1,
			Inventaire:     map[string]int{"Potion de soin": 2, "Potion de poison": 6},
			upgradesBought: 0,
			InventoryLimit: 10,
			// Sorts
			Sorts: [][]string{
				{"🪄  Wingardium Leviosa !", "est envoyé loin et perd 10 PV.", "10"},
				{"🔥 Incendio !", "est en feu et perd 15 PV.", "15"},
				{"⚡ Stupefix !", "est étourdi et perd 20 PV.", "20"},
			},
			// Tenue + Arme
			Equipement: make([][]string, 3), // Casque, Armure, Jambières
			Armes:      make([]string, 1),   // Arme
			// Mana
			Mana:    30,
			MaxMana: 30,
		}
	}

	// Easter Egg Harry Potter (HP)
	if classe == 7 {
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
			Sorts: [][]string{
				{"🪄  Wingardium Leviosa !", "est envoyé loin et perd 10 PV.", "10"},
				{"🔥 Incendio !", "est en feu et perd 15 PV.", "15"},
				{"⚡ Stupefix !", "est étourdi et perd 20 PV.", "20"},
			},
			// Tenue + Arme
			Equipement: make([][]string, 3), // Casque, Armure, Jambières
			Armes:      make([]string, 1),   // Arme
			// Mana
			Mana:    90,
			MaxMana: 90,
		}
	}

	// Choix par défaut pour éviter les bugs
	return Character_class{
		Name:       nom,
		Class:      "Inconnu",
		Pv:         50,
		MaxPv:      50,
		Attaque:    10,
		Defence:    5,
		Gold:       0,
		Niveau:     1,
		Inventaire: map[string]int{},
		Mana:       20,
		MaxMana:    20,
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
			Type_attaque: [][]string{
				{"Griffe", "Le gobelin vous griffe vicieusement", "10"},
				{"Coup de massue", "Le gobelin vous assène un coup de massue", "15"},
				{"Coup de pied", "Le gobelin vous donne un coup de pied", "20"},
			},
			Empoisonne:  false,
			Tour_poison: 0,
		}
	}
	if monstre == "boss" {
		return Character_class{
			Name:    "Boss Final",
			Class:   "Monstre",
			Pv:      100,
			MaxPv:   100,
			Attaque: 30,
			Defence: 5,
			Gold:    40, // Récompense
			Niveau:  3,
			Type_attaque: [][]string{
				{"Griffe puissante", "Le boss vous griffe avec une force terrible", "15"},
				{"Coup de massue brutal", "Le boss vous assène un coup de massue dévastateur", "20"},
				{"Coup de pied écrasant", "Le boss vous écrase avec son pied", "25"},
				{"Charge dévastatrice", "Le boss vous charge de toutes ses forces", "30"},
			},
			Empoisonne:  false,
			Tour_poison: 0,
		}
	}
	// Ennemi par défaut
	return Character_class{
		Name:    "Orc",
		Class:   "Monstre",
		Pv:      60,
		MaxPv:   60,
		Attaque: 20,
		Defence: 5,
		Gold:    30, // Récompense
		Niveau:  2,
		Type_attaque: [][]string{
			{"Coup d'épée", "L'orc vous frappe avec son épée", "18"},
			{"Rugissement", "L'orc rugit et vous intimide", "12"},
		},
		Empoisonne:  false,
		Tour_poison: 0,
	}
}

func TutoEnnemi() Character_class {
	return Character_class{
		Name:    "Slime d'entraînement",
		Class:   "Il est consentant, je vous jure",
		Pv:      70,
		MaxPv:   70,
		Attaque: 15, // Réduit pour le tutoriel
		Defence: 2,
		Gold:    0,
		Niveau:  1,
		Type_attaque: [][]string{
			{"Coup de slime", "Le slime vous frappe mollement", "8"},
			{"Jet d'acide faible", "Le slime crache un peu d'acide", "12"},
		},
		Empoisonne:  false,
		Tour_poison: 0,
	}
}
