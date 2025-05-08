package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Définition de la structure de réponse JSON
type Greeting struct {
	Message string `json:"message"`
}

// Handler REST : retourne un message JSON
func apiGreetHandler(w http.ResponseWriter, r *http.Request) {
	// On définit le header Content-Type comme JSON
	w.Header().Set("Content-Type", "application/json")

	// Création de l'objet réponse
	response := Greeting{Message: "Hello from the Go API! 👋"}

	// Encodage en JSON et envoi
	json.NewEncoder(w).Encode(response)
}

// Route racine (toujours dispo)
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go server! Try /api/greet")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api/greet", apiGreetHandler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
