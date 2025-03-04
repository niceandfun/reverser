package main

import (
	"fmt"

	"reverser/internal/wisdom"
)

func main() {
	w := wisdom.New()
	fmt.Println(*w)
	w.UpdateWisdom()
	fmt.Println(*w)
	w.CleanWisdom()
	fmt.Println(*w)
}
