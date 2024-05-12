package swordmanboy

import (
	"fmt"
)

func SayHello(name string, round int) {
	for i := 0; i < round; i++ {
		fmt.Println("Hello ", name, ", round: ", i)
	}
}
