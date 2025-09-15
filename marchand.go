package MSA

import "fmt"

// ==================== STRUCTURE MARCHAND ====================
type Merchant struct {
	Name    string
	Gold    int
	Potions map[string]int
	Items   map[string]int
}

// Crée un nouveau marchand
func NewMerchant(name string, gold int) *Merchant {
	return &Merchant{
		Name: name,
		Gold: gold,
		Potions: map[string]int{
			"Potion de soin":   200,
			"Potion de poison": 200,
		},
		Items: map[string]int{
			"Bois de sureau":        500,
			"Plume de phénix":       600,
			"Baguette surpuissante": 1000,
		},
	}
}

// Affiche les objets d'une catégorie et retourne la liste des noms
func (m *Merchant) DisplayCategory(choice int) []string {
	var keys []string
	switch choice {
	case 1:
		fmt.Println("Catégorie : Potions")
		i := 1
		for potion, price := range m.Potions {
			fmt.Printf("%d - %s : %d or\n", i, potion, price)
			keys = append(keys, potion)
			i++
		}
	case 2:
		fmt.Println("Catégorie : Objets Magiques")
		i := 1
		for item, price := range m.Items {
			fmt.Printf("%d - %s : %d or\n", i, item, price)
			keys = append(keys, item)
			i++
		}
	default:
		fmt.Println("Choix invalide")
	}
	return keys
}

// ==================== FONCTION D'ACHAT ====================
func Buy(joueur *Character_class, m *Merchant, category int, itemName string, quantity int) {
	// Vérifie la quantité
	if quantity < 1 || quantity > 5 {
		fmt.Println("❌ Vous ne pouvez acheter qu'entre 1 et 5 exemplaires.")
		return
	}

	var price int
	if category == 1 {
		val, ok := m.Potions[itemName]
		if !ok {
			fmt.Println("❌ Potion introuvable.")
			return
		}
		price = val * quantity
	} else if category == 2 {
		val, ok := m.Items[itemName]
		if !ok {
			fmt.Println("❌ Objet introuvable.")
			return
		}
		price = val * quantity
	} else {
		fmt.Println("❌ Catégorie invalide.")
		return
	}

	// Vérifie si le joueur a assez d'or
	if joueur.Gold < price {
		fmt.Println("❌ Pas assez d'or pour cet achat.")
		return
	}

	// Déduit l'or et ajoute l'objet à l'inventaire du joueur
	joueur.Gold -= price
	for i := 0; i < quantity; i++ {
		joueur.Inventaire = append(joueur.Inventaire, itemName)
	}

	fmt.Printf("✅ Vous avez acheté %d x %s. Or restant : %d\n", quantity, itemName, joueur.Gold)
}

// ==================== FONCTION VILLE ====================
func Ville(joueur *Character_class) {
	merchant := NewMerchant("Mandragor", 1000)

	fmt.Println("👋 Bonjour, je suis", merchant.Name, "! Que puis-je faire pour toi ?")
	fmt.Println("1 - Potions")
	fmt.Println("2 - Objets Magiques")

	var choice int
	fmt.Scanln(&choice)

	// Récupère les objets de la catégorie choisie
	items := merchant.DisplayCategory(choice)
	if len(items) == 0 {
		return
	}

	// Choix de l'objet par numéro
	fmt.Println("Choisis un objet en entrant son numéro :")
	var itemIndex int
	fmt.Scanln(&itemIndex)
	if itemIndex < 1 || itemIndex > len(items) {
		fmt.Println("❌ Choix invalide.")
		return
	}
	itemName := items[itemIndex-1]

	// Choix de la quantité
	fmt.Println("Combien veux-tu en acheter ? (max 5)")
	var qty int
	fmt.Scanln(&qty)

	// Achat
	Buy(joueur, merchant, choice, itemName, qty)

	// Affiche l'inventaire du joueur
	fmt.Println("📦 Inventaire du joueur :")
	for _, item := range joueur.Inventaire {
		fmt.Println("-", item)
	}
}
