package main

import (
	"fmt"
	"module/repos" // ваш пакет
)

func main() {
	if sum := repos.Add(1, 2); sum != 3 {
		panic(fmt.Sprintf("sum expected to be 3; got %d", sum))
	}

	fmt.Println("Well done!")
}
