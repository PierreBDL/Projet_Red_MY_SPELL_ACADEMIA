package MSA

import "fmt"

// ==================== STRUCTURE FORGERON ====================
type Blacksmith struct {
	Name string
}

// Crée un nouveau forgeron
func NewBlacksmith(name string) *Blacksmith {
	return &Blacksmith{Name: name}
}

// ==================== FONCTION DE CRAFT ====================
func (b *Blacksmith) CraftWand(joueur *Character_class) {
	fmt.Println("🔨 Bienvenue dans la forge d'", b.Name, "!")
	fmt.Println("Pour crafter la baguette de sureau, il te faut :")
	fmt.Println("- 1 Bois de sureau")
	fmt.Println("- 1 Plume de phénix")
	fmt.Println("- 400 or")

	// Vérification des ressources
	boisOK, plumeOK := false, false
	for _, item := range joueur.Inventaire {
		if item == "Bois de sureau" && !boisOK {
			boisOK = true
		}
		if item == "Plume de phénix" && !plumeOK {
			plumeOK = true
		}
	}

	if !boisOK {
		fmt.Println("❌ Il te manque un Bois de sureau.")
		return
	}
	if !plumeOK {
		fmt.Println("❌ Il te manque une Plume de phénix.")
		return
	}
	if joueur.Gold < 400 {
		fmt.Println("❌ Tu n'as pas assez d'or.")
		return
	}

	// Retirer les ressources de l'inventaire
	for i, item := range joueur.Inventaire {
		if item == "Bois de sureau" {
			joueur.Inventaire = append(joueur.Inventaire[:i], joueur.Inventaire[i+1:]...)
			break
		}
	}
	for i, item := range joueur.Inventaire {
		if item == "Plume de phénix" {
			joueur.Inventaire = append(joueur.Inventaire[:i], joueur.Inventaire[i+1:]...)
			break
		}
	}

	// Déduction de l'or
	joueur.Gold -= 400

	// Ajout de la baguette
	joueur.Inventaire = append(joueur.Inventaire, "Baguette de sureau")
	fmt.Println("✅ La baguette de sureau a été forgée et ajoutée à ton inventaire !")
}

// ==================== FONCTION VILLE FORGE ====================
func Forge(joueur *Character_class) {
	forgeron := NewBlacksmith("Olivenders")

	fmt.Println("\n👋 Bonjour aventurier ! Tu es dans la forge d’Olivenders.")
	fmt.Println("Veux-tu forger la baguette de sureau ? (oui/non)")

	var response string
	fmt.Scanln(&response)

	if response == "oui" {
		forgeron.CraftWand(joueur)
	} else {
		fmt.Println("🔚 Peut-être une prochaine fois.")
	}

	// Inventaire final
	fmt.Println("\n📦 Inventaire du joueur :")
	for _, item := range joueur.Inventaire {
		fmt.Println("-", item)
	}
	fmt.Printf("💰 Or restant : %d\n", joueur.Gold)
}
