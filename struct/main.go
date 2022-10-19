package main

import "fmt"

import "encoding/json"

type color string

type Rectangle struct {
	X int
	Y int
	C *Rectangle // only as pointer
}

type Circle struct {
	S float64 `json:"Ubabupa"`
}

func (s *Circle) String() string {
	return fmt.Sprintf("%s", s.S)
}

func (s Circle) UpdateArea(a float64) {
	s.S = a
}

func main() {
	// exercise
	type Room struct {
		Area  float64
		Color color
	}
	type MyFlat struct {
		Bathroom Room `json:"bathroom"`
		Lounge   Room `json:"lounge"`
		Bedroom  Room `json:"bedroom"`
	}

	// rekuperator := Rectangle{10, 20, &Rectangle{}}
	// fmt.Printf("%v %T\n", rekuperator, rekuperator)

	// rekuperator.X = 30
	// rekuperator.Y = 50
	// fmt.Printf("%d\n", rekuperator.X)
	// fmt.Printf("%v %T\n", rekuperator, rekuperator)

	c := &Circle{S: 120.5}
	fmt.Println(c)

	// res, _ := json.Marshal(c)
	// fmt.Println(string(res)) // []byte

	c.UpdateArea(200)
	fmt.Println(c)
}
