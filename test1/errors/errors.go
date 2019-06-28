package errors

import (
	"log"
	"os"
)

func init() {
	testErrors()
}

func testErrors() {
	f, err := os.Open("dont_exist")
	if err == nil {
		f.Close()
	} else {
		log.Printf("got_error: %s\n", err.Error())
	}
}
