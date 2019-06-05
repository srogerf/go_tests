package main

import (
	_ "fmt"
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
	"log"
	_ "log"
	"os"
	_ "os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.Fatal("init")
}
func main() {
	log.SetOutput(os.Stdout)
	log.Fatal("helloi\n\n")
	i := 5
	log.Fatal("%s\n\n", i)
	search.Run("rest")
}
