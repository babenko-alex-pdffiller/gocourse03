package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(100); n % 9 {
	case 0:
		fmt.Println(n, "is a multiple of 9.")
	case 1, 2, 3:
		fmt.Println(n, "mod 9 is 1, 2 or 3.")
		break // is nonsense
	case 4, 5, 6, 7:
		fmt.Println(n, "mod 9 is 4, 5 or 6.")
	default:
		fmt.Println(n, "mod 9 is 7 or 8.")
	}

	rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(100) % 5; n {
	case 0, 1, 2, 3, 4:
		fmt.Println("n =", n)
		fallthrough
	case 5, 6, 7, 8:
		n := 99               // for current branch code block
		fmt.Println("n =", n) // 99
		fallthrough
		//_ = n
	default:
		fmt.Println("n =", n)
	}

	switch n := 5; n {
	}

	switch 5 {
	}

	switch _ = 5; {
	}

	switch /*true*/ {
	}

	switch /*cond*/ {
	case true:
		fmt.Println("hello")
	default:
		fmt.Println("bye")
	}

	rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(3); n {
	default:
		//fmt.Println("n == 2")
		fallthrough
	case 0:
		fmt.Println("n == 0")
	case 1:
		fmt.Println("n == 1")
	}
}
