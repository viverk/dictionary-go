package main

import (
	"bufio"
	"estiam/dictionary"
	"fmt"
	"os"
)

func main() {
	dict := dictionary.New()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Choisir une action:")
		fmt.Println("1. Add")
		fmt.Println("2. Define")
		fmt.Println("3. Remove")
		fmt.Println("4. List")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Quel est votre choix ? : ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Erreur:", err)
			continue
		}

		switch choice {
		case 1:
			actionAdd(dict, reader)
		case 2:
			actionDefine(dict, reader)
		case 3:
			actionRemove(dict, reader)
		case 4:
			actionList(dict)
		case 5:
			fmt.Println("Retour")
			return
		default:
			fmt.Println("Choix invalide. Entrez un numero entre 1 et 5.")
		}
	}
}

func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Entrer un mot à ajouter dans le dico: ")
	word, _ := reader.ReadString('\n')
	word = word[:len(word)-1]

	fmt.Print("Entrer la definition : ")
	definition, _ := reader.ReadString('\n')
	definition = definition[:len(definition)-1]

	err := d.Add(word, definition)
	if err != nil {
		fmt.Println("Erreur:", err)
	} else {
		fmt.Println("Le mot a bien été ajouté !")
	}
}

func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Quel est le mot dont vous voulez avoir la definition ? : ")
	word, _ := reader.ReadString('\n')
	word = word[:len(word)-1]

	definition, err := d.Get(word)
	if err != nil {
		fmt.Println("le mot est inconnu:", err)
	} else {
		fmt.Printf("Definition de '%s': %s\n", word, definition)
	}
}

func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Entrer le mot à supprimer du dico: ")
	word, _ := reader.ReadString('\n')
	word = word[:len(word)-1]

	err := d.Remove(word)
	if err != nil {
		fmt.Println("Le mot est inconnu du dico:", err)
	} else {
		fmt.Println("Le mot a bien été retiré du dico")
	}
}

func actionList(d *dictionary.Dictionary) {
	entries, err := d.List()
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	fmt.Println("Contenu du dico:")
	for word, definition := range entries {
		fmt.Printf("Mot: %s, Definition: %s\n", word, definition)
	}
}
