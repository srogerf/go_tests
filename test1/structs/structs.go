package structs

import (
	"log"
)

func init() {
	log.Println("-----------------")
	log.Println("running structs")
	log.Println("-----------------")

}

func Run() {
	initializeArray()
}

// some basic array initialization
func initializeArray() {
	a := [3]int{1: 4, 2: 7}
	a[1] = 5

	b := [2]int{5, 10}
	b[0] = a[1]

	//using calculated size
	c := [...]int{2, 3, 4}
	log.Println(c)

	//array pointers
	d := [4]*int{0: new(int), 1: new(int)}
	*d[0] = 5
	log.Println(d, *d[0])

	//multidimensional
	var e [3][5]int
	log.Println(e)

	testPassing(&a)

}

// get and print an array
func testPassing(myArray *[3]int) {
	log.Printf("Passed array %d", myArray[2])
}
