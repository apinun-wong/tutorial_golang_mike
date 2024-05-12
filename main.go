package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mikelopster/go-example/book"
	"github.com/mikelopster/go-example/swordmanboy"
)

func main() {
	id := uuid.New()
	fmt.Println("Hello World 2")
	fmt.Println("UUID: %s", id)
	swordmanboy.SayHello("Nun", 2)
	swordmanboy.SayTest()

	fmt.Println("Math: ", swordmanboy.Add(3214, 322))
	book.StartBooks()
}
