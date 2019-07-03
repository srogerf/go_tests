package files

import (
	"io/ioutil"
	"log"
)

func init() {
	log.Println("files: starting")
	ListDir()
	log.Println("files: ending")
}

func ListDir() {
	log.Println("files: listdir")
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		log.Printf("files: >%s\n", f.Name())
	}
}
