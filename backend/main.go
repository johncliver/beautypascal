package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID                int               `json:"id"`
	ShoppingBags      []ShoppingBag     `json:"shoppingbag"`
	ShippingAddresses []ShippingAddress `json:"shippingaddress"`
	ShippingMethod    string            `json:"shippingmethod"`
	Currency          string            `json:"currency"`
	Cards             []CreditCard      `json:"card"`
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

type CreditCard struct {
	Credit string `json:"credict"`
	Date   string `json:"date"`
	Cvv    string `json:"cvv"`
}

type ListUser struct {
	Users []User `json:"users"`
}

var users = ListUser{
	Users: []User{},
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser User

	// Bind JSON body to User struct
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign a new ID to the user
	newUser.ID = len(users.Users) + 1

	// Add the new user to the list
	users.Users = append(users.Users, newUser)

	// Return the new user as response
	c.JSON(http.StatusCreated, newUser)
}

func main() {
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	router.GET("/users", getUsers)
	router.POST("/users", createUser)

	router.Run(":8080")
}
