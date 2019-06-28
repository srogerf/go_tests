package structs

import "log"

// slices can grow and shrink and provide garage collection

func init() {
	log.Println("Running slices")
	slices_101_ints()
	slices_102_strings()
	slices_103_vs_arrays()
	slices_104_nil()
	slices_105_manipulate()
	log.Println("Running slices: done")
}
func slices_101_ints() {
	//slice of ints len 3, cap 5
	// len = inital cap, cap = max cap
	testMe := make([]int, 3, 5)
	log.Printf("Slice 101_1: %d\n", testMe)

	//slice of len 45, can init just used items
	testMe2 := []int{45: 33}
	log.Printf("Slice 101_2: %d\n", testMe2)
}
func slices_102_strings() {
	//slice of strings cap 5
	testMe := make([]string, 5)
	log.Printf("Slice 102_1: %s\n", testMe)

	//slice of cap and len 5
	testMe2 := []string{"a", "b", "c", "d", "e"}
	log.Printf("Slice 102_2: %s\n", testMe2)
}

func slices_103_vs_arrays() {
	//array
	testMe := [3]int{1, 2, 3}
	log.Printf("Array 103_1: %d\n", testMe)

	//slice
	testMe2 := []int{1, 2, 3}
	log.Printf("Slice 101_2: %d\n", testMe2)
}
func slices_104_nil() {
	//create empty
	var testMe []int
	//testMe[2] = 1 error
	log.Printf("Slice 104_1: %d\n", testMe)

	//use make
	testMe2 := make([]int, 0)
	//testMe2[2] = 4 error
	log.Printf("Slice 104_2: %d\n", testMe2)

	//idiom
	testMe3 := []int{}
	//testMe3[2] = 7 error
	log.Printf("Slice 104_3: %d\n", testMe3)
}

func slices_105_manipulate() {
	//slice of cap and len 5
	testMe := []string{"a", "b", "c", "d", "e"}
	testMe2 := testMe[1:3]
	log.Printf("Slice 105_1: %s, sliced %s\n", testMe, testMe2)

	//still same storage
	testMe2[1] = "f"
	log.Printf("Slice 105_2: %s, sliced %s\n", testMe, testMe2)

	testMe2 = append(testMe2, "h")
	log.Printf("Slice 105_3: %s, sliced %s\n", testMe, testMe2)

	testMe2 = append(testMe2, "i")
	testMe2 = append(testMe2, "j")
	testMe = append(testMe2, "z")
	log.Printf("Slice 105_4: %s, sliced %s\n", testMe, testMe2)
}
