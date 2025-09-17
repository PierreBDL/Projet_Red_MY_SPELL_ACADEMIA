package MSA

import (
	"bufio"
	"fmt"
	"os"
)

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
	if joueur.Inventaire["Bois de sureau"] < 1 {
		fmt.Println("❌ Il te manque un Bois de sureau.")
		return
	}
	if joueur.Inventaire["Plume de phénix"] < 1 {
		fmt.Println("❌ Il te manque une Plume de phénix.")
		return
	}
	if joueur.Gold < 400 {
		fmt.Println("❌ Tu n'as pas assez d'or.")
		return
	}

	// Retirer les ressources de l'inventaire
	joueur.Inventaire["Bois de sureau"]--
	if joueur.Inventaire["Bois de sureau"] == 0 {
		delete(joueur.Inventaire, "Bois de sureau")
	}

	joueur.Inventaire["Plume de phénix"]--
	if joueur.Inventaire["Plume de phénix"] == 0 {
		delete(joueur.Inventaire, "Plume de phénix")
	}

	// Déduction de l'or
	joueur.Gold -= 400

	// Ajout de la baguette
	joueur.Inventaire["Baguette de sureau"]++
	fmt.Println("✅ La baguette de sureau a été forgée et ajoutée à ton inventaire !")
}

// ==================== FONCTION VILLE FORGE ====================
func Forge(joueur *Character_class) {
	forgeron := NewBlacksmith("Olivenders")

	fmt.Println("\n👋 Bonjour aventurier ! Tu es dans la forge d’Olivenders.")
	fmt.Println("Veux-tu forger la baguette de sureau ? (oui/non)")

	// Nettoyer le buffer d'entrée avant de demander
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') // Vider le buffer des caractères restants
	var response string
	fmt.Scanln(&response)

	if response == "oui" {
		forgeron.CraftWand(joueur)
	} else {
		fmt.Println("🔚 Peut-être une prochaine fois.")
	}

	// Inventaire final
	fmt.Println("\n📦 Inventaire du joueur :")
	for item, qty := range joueur.Inventaire {
		fmt.Printf("- %s x%d\n", item, qty)
	}
	fmt.Printf("💰 Or restant : %d\n", joueur.Gold)
}
