package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// shadowing
	var i = 100
	{
		i := 1
		fmt.Println(i)
	}
	fmt.Println(i)

	// precondition
	r8r := rand.New(rand.NewSource(time.Now().UnixNano())) // r8r = randomizer
	if a, b := r8r.Intn(100), r8r.Intn(100); a > b {
		fmt.Println(`Correct!`)
	} else if a%b == 0 {
		fmt.Println(`Excellent!`)
	} else {
		fmt.Println(`Nope!`)
	}

	//fmt.Println(a, b) // DOESN'T WORK

	// switch case
	switch n := rand.Intn(20); n - 10 {
	case 0:
		fmt.Println(n, "minus 10 is 10")
	case 1, 2, 3:
		fmt.Println(n, "minus 10 is less than 10")
		break // is nonsense
	case 4, 5, 6, 7, 8, 9:
		fmt.Println(n, "minus 10 is more less than 10")
	case 10:
		fallthrough
	default:
		fmt.Println(n, "minus 10 is negative")
	}

	// empty switch case
	var x = rand.Intn(100)
	switch /*true*/ {
	case x < 100:
		fmt.Println(x, "is simple")
	//case 57: // DOESN'T WORK
	default:
		fmt.Println(x, "is 100")
	}

	// for loop
	for i := 0; i < 10; i++ {
		//if i%2 == 0 {
		//	break
		//	continue
		//}
		fmt.Println(i)
	}

	var str = "Hello, Світе!"
	for i, b := range []byte(str) {
		fmt.Println(i, ":", b)
	}

	//for /*true*/ {
	// will se later
	//}

	i = 0
Next: // only in namespace
	fmt.Println(i)
	i++
	if i < 5 {
		goto Next
	}

Begin:
	switch n := r8r.Intn(100); n % 9 {
	case 0:
		fmt.Println(n, "is a multiple of 9.")
	case 1, 2, 3:
		fmt.Println(n, "mod 9 is 1, 2 or 3.")
		break Begin
	case 4, 5, 6:
	default:
		fmt.Println(n, "mod 9 is 7 or 8.")
	}
}
