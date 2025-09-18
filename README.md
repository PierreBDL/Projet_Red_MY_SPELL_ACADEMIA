# 🧙 My Spell Academia

Ce jeu a été développé par l'équipe MSA (MY SPELL ACADEMIA) qui est composée de Pierre BOURGEOIS DE LAVERGNE et de Ethan LEGLISE. Vous incarnez un un sorcier qui doit combattre des monstres (gobelins). 

---

## 🎮 Fonctionnalités

- Choix dans le terminal (CMD).
- Création et personnalisation de votre perso.
- Combats avec du tour par tour contre des ennemis (gobelin et boss).
- Economie dans le jeu pour acheter au marché, forgeron ou passer une nuit à l'auberge
- ASCII Art + animations + Emoji pour les graphismes.
- Sons pour les attaques
- Musiques

## 🥚 Easter Egg
- Harry Potter est caché dans le jeu

---

## 📦 Installation

### Prérequis
Avoir **Go** installé (Version 1.20 ou supérieure).  
Avoir cloné ce dépôt git :


### Dans un dossier que vous avez choisis

```bash
# Accédez au dossier
cd mon/chemin/vers/mon/dossier/

# Clonez le dépôt
git clone https://github.com/PierreBDL/Projet_Red_MY_SPELL_ACADEMIA

# Aller dans le dossier
cd PROJET_RED_MY_SPELL_ACADEMIA

# Installer les outils pour les sons
sudo apt install pulseaudio-utils
sudo apt install libasound2-dev pkg-config

# Lancez le jeu
go run ./bin/main.go
