package main

import "fmt"

func main() {
	// 1 definition
	// languageAges := map[string]int{"C": 1972, "Go": 2009}
	// var someMap1 map[string]int
	// someMap2 := make(map[string]int, 10)
	// someMap3 := *new(map[string]int)
	//
	// fmt.Printf("%T, %v\n", languageAges, languageAges)
	// fmt.Printf("%T, %v\n", someMap1, someMap1)
	// fmt.Printf("%T, %v\n", someMap2, someMap2)
	// fmt.Printf("%T, %v\n", someMap3, someMap3)

	// exercise // writing
	// myFirstMap := map[string]string{
	// 	`mom name`: `Name`,
	// 	`dad name`: `Name`,
	// 	`pet name`: `Name`,
	// }
	// fmt.Printf("%T, %v\n", myFirstMap, myFirstMap)

	//// unaddressable
	// someMap := map[string]int{"C": 1972, "Go": 2009}
	// somSlice := make([]int, 10, 10)
	// _ = &somSlice[1]
	// _ = &someMap[`C`]

	// exercise  // read it
	// pm1_1 := map[[2]int]int{{1, 2}: 1972, {3, 4}: 2009}
	// pm1_2 := map[[2]int]int{[2]int{1, 2}: 1972, [2]int{3, 4}: 2009}
	// pm2 := map[bool]string{} // why is bad?

	// incomparable
	//slice types
	//map types
	//function types
	//struct array with incomparable fields/elements
	//// nope
	// pm1 := map[[]int]string{}
	// pm2 := map[map[string]any]string{}

	// 2 init and update
	// type person struct{ age, height int }
	// persons := map[string]person{}
	// persons["John"] = person{age: 29}

	// someIntegers := map[int][5]int{}
	// someIntegers[1] = [5]int{1: 789}
	// fmt.Printf("%v, %v\n", persons, someIntegers)

	// exercise: update age
	// relatives := make(map[string]int, 3)
	// relatives[`mom name`] = 50
	// relatives[`dad name`] = 60
	// relatives[`pet name`] = 10
	// fmt.Printf("%T, %v\n", relatives, relatives)

	// 3 // read it (різниця між типом і змінною) (Як можна спростити запис знизу)
	langs := map[struct{ dynamic, strong bool }]map[string]int{
		{true, false}:  {"JavaScript": 1995},
		{false, true}:  {"Go": 2009},
		{false, false}: {"C": 1972},
	}
	destination1 := map[*struct{ dynamic, strong bool }]*map[string]int{}
	for category, langInfo := range langs {
		destination1[&category] = &langInfo // check if it will be effects on langInfo (value)
		category.dynamic, category.strong = true, true
	}
	for category, langInfo := range langs {
		fmt.Println(category, langInfo)
	}

	destination2 := map[struct{ dynamic, strong bool }]map[string]int{}
	for category, langInfo := range destination1 {
		destination2[*category] = *langInfo
	}
	// destination1 and destination2 both contain only one entry.
	fmt.Println(len(destination1), len(destination2)) // 1 1
	fmt.Println(destination2)                         // map[{true true}:map[C:1972]]
}
