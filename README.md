# 🧙 My Spell Academia

Ce jeu a été développé par l'équipe MSA (MY SPELL ACADEMIA) qui est composée de Pierre BOURGEOIS DE LAVERGNE et de Ethan LEGLISE. Vous incarnez un un sorcier qui doit combattre des monstres (gobelins). 

---

## 🎮 Fonctionnalités

- Choix dans le terminal (CMD).
- Création et personnalisation de votre sorcier.
- Combats au tour par tour contre des monstres.
- Economie dans le jeu pour acheter au marché et forgeron
- ASCII Art + Emoji pour les graphismes.
- Sons pour les attaques

---

## 📦 Installation

Assurez-vous d’avoir **Go** installé (>= 1.20).  
Clonez ce dépôt et compilez le projet :

```bash
# Clonez le dépôt
git clone https://github.com/ton-pseudo/my-spell-academia.git

# Accédez au dossier
cd PROJET_RED_MY_SPELL_ACADEMIA

# Installer les outils pour les sons
sudo apt install pulseaudio-utils
sudo apt install libasound2-dev pkg-config

# Lancez le jeu
go run ./bin/main.go
