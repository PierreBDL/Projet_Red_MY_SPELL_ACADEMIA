package MSA

import (
	"fmt"
	"os/exec"
	"runtime"
)

func jouerSon(fichier string) {
	switch runtime.GOOS {
	case "windows":
		exec.Command("start", fichier).Start()
	case "linux":
		// Essayer différents lecteurs audio Linux
		if err := exec.Command("aplay", fichier).Start(); err != nil {
			if err := exec.Command("mpg123", fichier).Start(); err != nil {
				if err := exec.Command("paplay", fichier).Start(); err != nil {
					fmt.Println("Impossible de jouer le son:", err)
				}
			}
		}
	case "darwin": // macOS
		exec.Command("afplay", fichier).Start()
	default:
		fmt.Println("OS non supporté pour les sons")
	}
}
