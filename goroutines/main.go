package main

import (
	"fmt"
	"time"
)

func addPrint(toAdd *int) {
	*toAdd += 1
	fmt.Println(*toAdd)
}

func main() {
	var first int = 0
	var second int = 0

	for i := 0; i < 10; i++ {
		time.Sleep(1000 * time.Millisecond)
		go addPrint(&first)
		addPrint(&second)
	}
}
