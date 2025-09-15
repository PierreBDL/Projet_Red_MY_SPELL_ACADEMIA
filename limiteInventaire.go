package MSA

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Adaptation pour Character_class
// On suppose que Character_class possède :
// Pv, MaxPv, Or, Inventaire (map[string]int), InventoryLimit, upgradesBought

// Ajouter un objet dans l'inventaire
func AddItem(joueur *Character_class, item string, qty int) {
	for TotalItems(joueur)+qty > joueur.InventoryLimit {
		fmt.Printf("❌ Inventaire plein ! Vous essayez d'ajouter %d %s.\n", qty, item)
		ShowInventory(joueur)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Voulez-vous jeter un objet pour faire de la place ? (oui/non) : ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))

		if answer != "oui" {
			fmt.Println("❌ Ajout annulé.")
			return
		}

		fmt.Print("Quel objet voulez-vous jeter ? ")
		removeItem, _ := reader.ReadString('\n')
		removeItem = strings.TrimSpace(removeItem)

		fmt.Print("Combien en voulez-vous jeter ? ")
		var removeQty int
		fmt.Scanln(&removeQty)

		RemoveItem(joueur, removeItem, removeQty)
	}

	joueur.Inventaire[item] += qty
	fmt.Printf("✅ %d %s ajouté(s) à votre inventaire.\n", qty, item)
}

// Supprimer un objet
func RemoveItem(joueur *Character_class, item string, qty int) {
	if currentQty, ok := joueur.Inventaire[item]; ok {
		if qty >= currentQty {
			delete(joueur.Inventaire, item)
		} else {
			joueur.Inventaire[item] -= qty
		}
		fmt.Printf("🗑️ %d %s supprimé(s) de l'inventaire.\n", qty, item)
	} else {
		fmt.Println("❌ Objet introuvable dans l'inventaire.")
	}
}

// Compter le nombre total d'objets
func TotalItems(joueur *Character_class) int {
	total := 0
	for _, qty := range joueur.Inventaire {
		total += qty
	}
	return total
}

// Afficher l'inventaire avec icônes
func ShowInventory(joueur *Character_class) {
	fmt.Println("📦 Inventaire du joueur :")
	for item, qty := range joueur.Inventaire {
		icon := ""
		switch item {
		case "potion_de_soin":
			icon = "🧪"
		case "potion_de_poison":
			icon = "🧪"
		case "bois_de_sureau":
			icon = "🌳"
		case "plume_de_phenix":
			icon = "🪶"
		case "baguette_surpuissante":
			icon = "✨"
		default:
			icon = "❔"
		}
		fmt.Printf("%s %s x%d\n", icon, item, qty)
	}
	fmt.Printf("Total objets : %d/%d\n", TotalItems(joueur), joueur.InventoryLimit)
}

// Améliorer la limite d'inventaire
func UpgradeInventory(joueur *Character_class) {
	if joueur.upgradesBought >= 3 {
		fmt.Println("❌ Vous avez déjà atteint la limite maximale d'améliorations.")
		return
	}
	joueur.InventoryLimit++
	joueur.upgradesBought++
	fmt.Printf("✅ Limite d'inventaire augmentée ! Nouvelle limite : %d\n", joueur.InventoryLimit)
}
