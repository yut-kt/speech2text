package util

import (
	"log"
	"os"
)

func OpenFile(f string) *os.File {
	fp, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return fp
}

func CloseFile(fp *os.File) {
	log.Printf("Close File: [%v]\n", fp.Name())
	err := fp.Close()
	if err != nil {
		panic(err)
	}
}
