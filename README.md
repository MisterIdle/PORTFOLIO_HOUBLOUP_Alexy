# Projet Web en Go

## Description

Ce projet est une application web développée en Go (Golang) avec une base de données SQLite. Il permet de gérer et d'afficher des données de plusieurs catégories (contacts, formations, expériences, compétences) via une interface web.
L'application utilise le framework Gin pour gérer les routes et les requêtes HTTP. Les pages web sont générées à l'aide de templates HTML, avec des données insérées dynamiquement depuis la base de données.

## Fonctionnalités

- Gestion des contacts, expériences, compétences, et formations.
- Ajout, suppression, et modification de données via une interface web.
- Affichage des données sous forme de tableaux dynamiques.
- API REST pour récupérer les données par catégorie.
  
## Prérequis

- **Go** (version 1.21 ou supérieure)
- **SQLite** installé

## Installation

1. Clonez le dépôt sur votre machine locale :
2. Exécutez la commande suivante dans le terminal pour compiler le programme :
   
   ```bash
   go run main.go
