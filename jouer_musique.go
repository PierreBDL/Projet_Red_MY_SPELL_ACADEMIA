package MSA

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
	streamer           beep.StreamSeekCloser
	format             beep.Format
	ctrl               *beep.Ctrl
	speakerInitialized bool
)

// JouerMusique joue la musique en boucle
func JouerMusique(cheminFichier string) error {
	// Arrêter la musique précédente
	ArreterMusique()

	// Vérifier si le fichier existe
	if _, err := os.Stat(cheminFichier); os.IsNotExist(err) {
		// Essayer d'autres chemins
		alternatives := []string{
			"./musics/" + cheminFichier,
			"musics/" + cheminFichier,
			cheminFichier,
		}

		found := false
		for _, alt := range alternatives {
			if _, err := os.Stat(alt); err == nil {
				cheminFichier = alt
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("fichier audio non trouvé: %s", cheminFichier)
		}
	}

	// Ouvrir le fichier audio
	f, err := os.Open(cheminFichier)
	if err != nil {
		return fmt.Errorf("erreur ouverture fichier: %v", err)
	}

	// Décoder le mp3
	newStreamer, newFormat, err := mp3.Decode(f)
	if err != nil {
		f.Close()
		return fmt.Errorf("erreur décodage MP3: %v", err)
	}

	// Initialiser le speaker une seule fois
	if !speakerInitialized {
		speaker.Init(newFormat.SampleRate, newFormat.SampleRate.N(time.Second/10))
		speakerInitialized = true
	}

	// Stocker les références
	streamer = newStreamer
	format = newFormat

	// Contrôleur pour pouvoir stopper (boucle infinie)
	ctrl = &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}

	// Jouer en arrière-plan
	speaker.Play(ctrl)

	return nil
}

// ArreterMusique stoppe la musique
func ArreterMusique() {
	if ctrl != nil {
		speaker.Lock()
		ctrl.Paused = true
		speaker.Unlock()
		speaker.Clear() // Vider la queue
	}
	if streamer != nil {
		streamer.Close()
		streamer = nil
	}
	ctrl = nil
}

// Fonction pour compatibilité avec l'ancien code
func JouerMusiqueMenu(nom string) error {
	return JouerMusique("../musics/" + nom)
}
