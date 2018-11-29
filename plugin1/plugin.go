package main

import (
	"fmt"

	"github.com/qaisjp/go-plugins-experiment/structs"
)

var Context *structs.Context

func init() {
	fmt.Println("Loader.go init loaded")
}

func main() {}

func OnLoad() {
	fmt.Println("main.go loaded")
	fmt.Printf("Host is: %s\n", Context.GetHost())
}
