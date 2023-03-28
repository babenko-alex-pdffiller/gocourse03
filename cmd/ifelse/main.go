package main

import (
	"flowcontrol/pkg"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var i = 100
	{
		i := 1
		fmt.Println(i)
	}
	fmt.Println(i)

	rand.Seed(time.Now().UnixNano())
	if pkg.IsItCorrect(rand.Int(), rand.Int()) {
		fmt.Println(`Correct!`)
	}

	if a, b := rand.Int(), rand.Int(); pkg.IsItCorrect(a, b) {
		fmt.Println(`Correct!`)
	} else if a%b == 0 {
		fmt.Println(`Not correct!`)
	} else {
		fmt.Println(`Nothing!`)
	}
}
