package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// for loop 1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//for loop 2
	var i = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Printf("\nand: %d\n\n", i)
	for i < 20 {
		fmt.Println(i)
		i++
	}

	// for loop 3
	for i := 0; ; i++ { // <=> for i := 0; true; i++ {
		if i >= 10 {
			// "break" statement will be explained below.
			break
		}
		fmt.Println(i)
	}

	//for loop 4
	for true {
	}
	for true {
	}
	for {
	}
	for /*true*/ {
	}

	//for loop 5
	for i := 0; ; i++ {
		if i >= 10 {
			break
		}
		fmt.Println(i)
	}

	//for loop 6
	for i := 0; i < 3; i++ {
		fmt.Print(i)
		i := i
		i = 10
		_ = i
	}

	//for loop 7
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i)
	}

	//for loop 8
	var str = "world"
	for i, b := range []byte(str) {
		fmt.Println(i, ":", b)
	}

	//for loop 8
	s := "éक्षिaπ囧"
	for i, b := range []byte(s) {
		fmt.Printf("The byte at index %v: 0x%x \n", i, b)
	}

	//for loop 8.1
	s = "éक्षिaπ囧"
	for i, b := range s {
		fmt.Printf("The byte at index %v: 0x%x \n", i, b)
	}

	i = 0
Next: // only in namespace
	fmt.Println(i)
	i++
	if i < 5 {
		goto Next
	}

	rand.Seed(time.Now().UnixNano())
Begin:
	switch n := rand.Intn(100); n % 9 {
	case 0:
		fmt.Println(n, "is a multiple of 9.")
	case 1, 2, 3:
		fmt.Println(n, "mod 9 is 1, 2 or 3.")
		break Begin
	case 4, 5, 6:
	default:
		fmt.Println(n, "mod 9 is 7 or 8.")
	}

	//exercise
	// 1. scan diagnosis from keyboard (fever, leaking nose, etc.) - suggest the illness. Use if/else
	// 2. scan the year season from keyboard. Suggest clothes. Use switch/case
	// 3. make a game, where user must guess the number. Program answers "more" or "less" to lead user to the right answer. Use for loop and branching
}
