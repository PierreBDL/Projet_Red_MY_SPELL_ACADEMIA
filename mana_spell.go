package MSA

type Spell struct {
	Name        string
	Description string
	Degats      int
	Mana_cost   int
}

// InitSpell retourne une liste de sorts disponibles
func InitSpells() []Spell {
	return []Spell{
		{
			Name:        "🪄  Wingardium Leviosa !",
			Description: "est envoyé loin et perd 10 PV.",
			Degats:      10,
			Mana_cost:   10,
		},
		{
			Name:        "🔥 Incendio !",
			Description: "est en feu et perd 15 PV.",
			Degats:      15,
			Mana_cost:   20,
		},
		{
			Name:        "⚡ Stupefix !",
			Description: "est étourdi et perd 20 PV.",
			Degats:      20,
			Mana_cost:   30,
		},
	}
}

// GetSpellByIndex retourne un sort par son index
func GetSpellByIndex(index int) (Spell, bool) {
	spells := InitSpells()
	if index >= 0 && index < len(spells) {
		return spells[index], true
	}
	return Spell{}, false
}

// GetSpellByName retourne un sort par son nom
func GetSpellByName(name string) (Spell, bool) {
	spells := InitSpells()
	for _, spell := range spells {
		if spell.Name == name {
			return spell, true
		}
	}
	return Spell{}, false
}

// InitWingardiumLeviosa retourne le sort Wingardium Leviosa
func InitWingardiumLeviosa() Spell {
	return Spell{
		Name:        "🪄  Wingardium Leviosa !",
		Description: "est envoyé loin et perd 10 PV.",
		Degats:      10,
		Mana_cost:   5,
	}
}

// InitIncendio retourne le sort Incendio
func InitIncendio() Spell {
	return Spell{
		Name:        "🔥 Incendio !",
		Description: "est en feu et perd 15 PV.",
		Degats:      15,
		Mana_cost:   10,
	}
}

// InitStupefix retourne le sort Stupefix
func InitStupefix() Spell {
	return Spell{
		Name:        "⚡ Stupefix !",
		Description: "est étourdi et perd 20 PV.",
		Degats:      20,
		Mana_cost:   15,
	}
}
