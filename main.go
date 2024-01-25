package main

import (
	"estiam/dictionary"
	"estiam/middleware"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the dictionary with BoltDB
	dict, err := dictionary.New()
	if err != nil {
		fmt.Printf("Error initializing dictionary: %v\n", err)
		return
	}
	defer dict.Close()

	r := mux.NewRouter()

	// Add the authentication middleware to the router
	r.Use(middleware.AuthenticationMiddleware)

	// Add the logging middleware to the router
	r.Use(middleware.LoggingMiddleware)

	// Define routes
	r.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		actionAdd(dict, w, r)
	}).Methods("POST")

	r.HandleFunc("/define/{word}", func(w http.ResponseWriter, r *http.Request) {
		actionDefine(dict, w, r)
	}).Methods("GET")

	r.HandleFunc("/remove/{word}", func(w http.ResponseWriter, r *http.Request) {
		actionRemove(dict, w, r)
	}).Methods("DELETE")

	r.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		actionList(dict, w, r)
	}).Methods("GET")

	// Start the server
	http.Handle("/", r)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func actionAdd(d *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	word := r.FormValue("word")
	definition := r.FormValue("definition")

	err := d.Add(word, definition)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "Le mot a bien été ajouté !")
}

func actionDefine(d *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	definition, err := d.Get(word)
	if err != nil {
		http.Error(w, fmt.Sprintf("Le mot est inconnu: %s", err), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Definition de '%s': %s\n", word, definition)
}

func actionRemove(d *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	err := d.Remove(word)
	if err != nil {
		http.Error(w, fmt.Sprintf("Le mot est inconnu du dico: %s", err), http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Le mot a bien été retiré du dico")
}

func actionList(d *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	entries, err := d.List()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur: %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Contenu du dico:")
	for word, definition := range entries {
		fmt.Fprintf(w, "Mot: %s, Definition: %s\n", word, definition)
	}
}
