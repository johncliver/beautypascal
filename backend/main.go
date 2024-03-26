package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type User struct {
	ID               int               `json:"id"`
	ShoppingBags     []ShoppingBag     `json:"shoppingbag"`
	ShippingAddressS []ShippingAddress `json:"shippingaddress"`
	ShippingMethods  string            `json:"shippingmethod"`
	Currencies       string            `json:"currency"`
	Card             []CredictCard     `json:"card"`
}

type ShoppingBag struct {
	Photo       string  `json:"photo"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float32 `json:"price"`
	SubTotal    float32 `json:"subtotal"`
}

type ShippingAddress struct {
	Email        string `json:"email"`
	Country      string `json:"country"`
	Name         string `json:"name"`
	LastName     string `json:"lastname"`
	AddressLine1 string `json:"AddressLine1"`
	AddressLine2 string `json:"AddressLine2"`
	Zip          string `json:"zip"`
	City         string `json:"city"`
	State        string `json:"state"`
	Mobile       string `json:"mobile"`
}

type CredictCard struct {
	Credict string `json:"credict"`
	Date    string `json:"date"`
	Cvv     string `json:"cvv"`
}

type ListUser struct {
	Users []User `json:"users"`
}

var users = ListUser{
	Users: []User{},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	// Log para depuración
	fmt.Println("Received POST request to /users")

	// Decodifica el cuerpo JSON de la solicitud y almacena los datos en newUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		// Log para depuración
		fmt.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Realiza validaciones adicionales si es necesario

	// Asigna un nuevo ID al usuario
	newUser.ID = len(users.Users) + 1

	// Agrega el nuevo usuario a la lista
	users.Users = append(users.Users, newUser)

	// Devuelve el nuevo usuario como respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	// Manejar solicitudes OPTIONS para el preflight CORS
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users", optionsHandler).Methods("OPTIONS") // Manejar OPTIONS para CORS

	// Crea un nuevo manejador CORS y úsalo para envolver el enrutador
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Actualiza según tus necesidades
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "application/json"},
		AllowCredentials: true,
		Debug:            true,
	})

	// Usa el manejador CORS envuelto en el enrutador
	handler := c.Handler(router)

	fmt.Printf("Listen And Server and Port: 8080")
	http.ListenAndServe(":5000", handler)
}
