package main

import "fmt"

type Age int

func (age Age) LargerThan(a Age) bool {
	return age > a
}
func (age *Age) IsNil() bool {
	return age == nil
}
func (age *Age) Increase() {
	*age++
}

// Receiver of custom defined function type.
type FilterFunc func(in int) bool

func (ff FilterFunc) Filter(in int) bool {
	return ff(in)
}

// Receiver of custom defined map type.
type StringSet map[string]struct{}

func (ss StringSet) Has(key string) bool {
	_, present := ss[key]
	return present
}
func (ss StringSet) Add(key string) {
	ss[key] = struct{}{}
}
func (ss StringSet) Remove(key string) {
	delete(ss, key)
}

// Receiver of custom defined struct type.
type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

func (b Book) SetAnotherPages(pages int) {
	b.pages = pages
}

type Books []Book

func (books Books) Modify() {
	// Modifications on the underlying part of
	// the receiver will be reflected to outside
	// of the method.
	books[0].pages = 500
	// Modifications on the direct part of the
	// receiver will not be reflected to outside
	// of the method.
	books = append(books, Book{789})
}

func (books *Books) ModifyAnother() {
	*books = append(*books, Book{789})
	(*books)[0].pages = 500
}

//func Book.Pages(b Book) int {
//	// The body is the same as the Pages method.
//	return b.pages
//}
//
//func (*Book).SetPages(b *Book, pages int) {
//	// The body is the same as the SetPages method.
//	b.pages = pages
//}

func main() {
	// exercise
	// MyFlat.Area() float64
	// MyFlat.Address() string
	// MyFlat.UpdateBathroomArea(float64)
	var book Book

	fmt.Printf("%T \n", book.Pages)       // func() int
	fmt.Printf("%T \n", (&book).SetPages) // func(int)
	// &book has an implicit method.
	fmt.Printf("%T \n", (&book).Pages) // func() int

	// Call the three methods.
	(&book).SetPages(123)
	book.SetPages(123)           // equivalent to the last line
	fmt.Println(book.Pages())    // 123
	fmt.Println((&book).Pages()) // 123

	var m map[string]int // nil

	_ = (StringSet(nil)).Has   // will not panic
	_ = ((*Age)(nil)).IsNil    // will not panic
	_ = ((*Age)(nil)).Increase // will not panic
	//
	//_ = (StringSet(nil)).Has("key") // will not panic
	//_ = ((*Age)(nil)).IsNil()       // will not panic
	//
	//// This following line will panic. But the
	//// panic is not caused by invoking the method.
	//// It is caused by the nil pointer dereference
	//// within the method body.
	// ((*Age)(nil)).Increase()

	//var b Book
	//b.SetPages(123)
	//fmt.Println(b.pages) // 0

	//var books = Books{{123}, {456}}
	//books.Modify()
	//fmt.Println(books) // [{500} {456}]

	//var books = Books{{123}, {456}}
	//books.ModifyAnother()
	//fmt.Println(books) // [{500} {456} {789}]
}
