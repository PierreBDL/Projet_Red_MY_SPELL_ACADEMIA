package MSA

type Character_class struct {
	Name    string
	Class   string
	Pv      int
	MaxPv   int
	Attaque int
	Defence int
	Gold    int
	Niveau  int
}

func InitCharacter(nom string, classe int) Character_class {
	if classe == 1 {
		return Character_class{
			Name:    nom,
			Class:   "Sorcier",
			Pv:      100,
			MaxPv:   100,
			Attaque: 20,
			Defence: 5,
			Gold:    100,
			Niveau:  1,
		}
	}
	if classe == 2 {
		return Character_class{
			Name:    nom,
			Class:   "Alchimiste",
			Pv:      80,
			MaxPv:   80,
			Attaque: 30,
			Defence: 7,
			Gold:    100,
			Niveau:  1,
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

func InitEnnemi() Character_class {
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
