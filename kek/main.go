package main

import (
	"fmt"

	"github.com/niceandfun/reverser/internal/wisdom"
)

func main() {
	w := wisdom.New()
	fmt.Println(w)
	w.updateWisdom()
	fmt.Println(w)
	w.cleanWisdom()
	fmt.Println(w)
}
