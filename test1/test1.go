package main

import (
	"github.com/srogerf/go_tests/test1/loops"
	"github.com/srogerf/go_tests/test1/structs"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func testMe() {
	//	log.Fatal("test")
}
func main() {
	log.Println("Test programs")
	var i = 5
	log.Println("%i\n\n", i)

	testMe()
	loops.Run()
	structs.Run()
}
