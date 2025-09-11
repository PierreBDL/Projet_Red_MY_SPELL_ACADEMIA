package MSA

import "strconv"

func InitCharacter(name string) []string {
	nom := name
	class := "Sorcier"
	pv := 100
	maxPv := 100
	attaque := 20
	defence := 5

	// Préparation à l'envoi
	caracter := make([]string, 6)
	caracter[0] = nom
	caracter[1] = class
	caracter[2] = strconv.Itoa(pv)
	caracter[3] = strconv.Itoa(maxPv)
	caracter[4] = strconv.Itoa(attaque)
	caracter[5] = strconv.Itoa(defence)

	return caracter
}
