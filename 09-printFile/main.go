package main

import (
	"io"
	"log"
	"os"
)

func main() {
	//fmt.Println(os.Args)
	//fmt.Println(os.Args[1])
	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, f)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
