package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	dict := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Choose an action:")
		fmt.Println("1. Add")
		fmt.Println("2. Define")
		fmt.Println("3. Remove")
		fmt.Println("4. List")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("quel est votre choix ? : ")
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
			fmt.Println("Retour...")
			return
		default:
			fmt.Println("Choix invalide. Entrez un numero entre 1 et 5.")
		}
	}
}

func actionAdd(dict map[string]string, reader *bufio.Reader) {
	fmt.Print("Entrer un mot à ajouter dans le dico: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	fmt.Print("Entrer la definition : ")
	definition, _ := reader.ReadString('\n')
	definition = strings.TrimSpace(definition)

	dict[word] = definition
	fmt.Println("Le mot a bien été ajouté !")
}

func actionDefine(dict map[string]string, reader *bufio.Reader) {
	fmt.Print("Quel est le mot dont vous voulez avoir la definition: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word) // Nettoyer la chaîne entrée

	definition, found := dict[word]
	if found {
		fmt.Printf("Definition de '%s': %s\n", word, definition)
	} else {
		fmt.Printf("'%s' le mot est inconnu.\n", word)
	}
}

func actionRemove(dict map[string]string, reader *bufio.Reader) {
	fmt.Print("Entrer le mot à supprimer du dico: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word) // Nettoyer la chaîne entrée

	_, found := dict[word]
	if found {
		delete(dict, word)
		fmt.Println("Le mot a bien été retiré du dico")
	} else {
		fmt.Printf("'%s' Le mot est inconnu du dico.\n", word)
	}
}

func actionList(dict map[string]string) {
	fmt.Println("Contenu du dico:")
	if len(dict) == 0 {
		fmt.Println("le dico est vide.")
		return
	}
	for word, definition := range dict {
		fmt.Printf("Mot: %s, Definition: %s\n", word, definition)
	}
}
