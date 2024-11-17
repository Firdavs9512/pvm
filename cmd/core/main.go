package main

import (
	"fmt"

	"github.com/Firdavs9512/pvm/core/php"
)

func main() {
	fmt.Println("Hello, World!")

	manager, err := php.NewManager()
	if err != nil {
		fmt.Println("Error creating manager:", err)
		return
	}

	fmt.Println(manager.Versions())
}
