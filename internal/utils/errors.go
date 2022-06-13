package utils

import (
	"log"
	"os"
)

func Check(e error, message string) {
	if e != nil {
		log.Fatal(message + e.Error())
		panic(e)
		os.Exit(1)
	}
}
