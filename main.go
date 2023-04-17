package main

import (
	"fmt"
	"github.com/ShinyShield/AbyShopServ/config"
	"github.com/ShinyShield/AbyShopServ/pkg/server"
	"github.com/ShinyShield/AbyShopServ/pkg/store"
	//"github.com/labstack/gommon/log"
)

type Input struct {
	Name     string
	Email    string
	Password string
}

func main() {
	// Capture connection properties.

	config.Init()
	fmt.Println("Connected!")

	server.Init()
	store.Init()
	// Connect to database
	//	var err = store.Init()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer store.Close()
	//
	//	// Test registration
	//	input := Input{
	//		Name:     "John Doe",
	//		Email:    "johndoe@example.com",
	//		Password: "password123",
	//	}
	//	output, err := Register(input)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Printf("New user registered with ID %d, name %s, and token %s\n", output.ID, output.Name, output.Token)
	//
	//	// Test login with new user
	//	loginInput := Input{
	//		Email:    "johndoe@example.com",
	//		Password: "password123",
	//	}
	//	loginOutput, err := Exec(loginInput)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Printf("User logged in with ID %d, name %s, and token %s\n", loginOutput.ID, loginOutput.Name, loginOutput.Token)
	//}
}
