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

func InitCharacter(nom string) Character_class {
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
