package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	purchase_options := []string{"店でも", "Webでも"}

	for i := 0; i < 8; i++ {
		fmt.Print(extract(purchase_options))
	}
}

func extract(options []string) string {
	rand.Seed(time.Now().UnixNano())
	randomIndexNum := rand.Intn(len(options))
	return options[randomIndexNum]
}
