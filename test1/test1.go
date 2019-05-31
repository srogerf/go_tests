package main

import (
	"log"
	"os"
	"fmt"
)

func init() {
	log.SetOutput(os.Stdout)
	log.Fatal("cc")
}
func main() {
	fmt.Printf("helloi\n\n")
}
