package main

// exercise
func SquaresOfSumAndDiff1(a, b int64) (s, d int64) {
	x, y := a+b, a-b
	s = x * x
	d = y * y
	return // <=> return s, d
}

func SquaresOfSumAndDiff2(a int64, b int64) (int64, int64) {
	return (a + b) * (a + b), (a - b) * (a - b)
}

func f() (x int, y bool) {
	println(x, y) // 0 false
	return
}

func SquaresOfSumAndDiff(a, b int64) (s, d int64) {
	return (a + b) * (a + b), (a - b) * (a - b)
	//s = (a+b) * (a+b); d = (a-b) * (a-b); return
}

func CompareLower4bits(m, n uint32) (r bool) {
	r = m&0xF > n&0xF
	return
	// The above two lines is equivalent to
	// the following line.
	/*
		return m&0xF > n&0xF
	*/
}

// This function has no parameters. The result
// declaration list is composed of only one
// anonymous result declaration, so it is not
// required to be enclosed within ().
func VersionString() string {
	return "go1.0"
}

// This function has no results. And all of
// its parameters are anonymous, for we don't
// care about them. Its result declaration
// list is blank (so omitted).
func doNothing(string, int) {
}

// Initialize a package-level variable
// with a function call.
var v = VersionString()

func main() {
	// exercise
	fnAsVal := SquaresOfSumAndDiff1

	println(v) // v1.0
	x1, y1 := SquaresOfSumAndDiff(3, 6)
	println(x1, y1) // 81 9
	b := CompareLower4bits(uint32(x1), uint32(y1))
	println(b) // false
	// "Go" is deduced as a string,
	// and 1 is deduced as an int32.
	doNothing("Go", 1)

	// This anonymous function has no parameters
	// but has two results.
	x, y := func() (int, int) {
		println("This function has no parameters.")
		return 3, 4
	}() // Call it. No arguments are needed.

	// The following anonymous function have no results.

	func(a, b int) {
		// The following line prints: a*a + b*b = 25
		println("a*a + b*b =", a*a+b*b)
	}(x, y) // pass argument x and y to parameter a and b.

	func(x int) {
		// The parameter x shadows the outer x.
		// The following line prints: x*x + y*y = 32
		println("x*x + y*y =", x*x+y*y)
	}(y) // pass argument y to parameter x.

	func() {
		// The following line prints: x*x + y*y = 25
		println("x*x + y*y =", x*x+y*y)
	}() // no arguments are needed.

	// exercise
	func(fn func(a, b int64) (s, d int64)) {
		fnAsVal(3, 6)
	}(fnAsVal)
}
