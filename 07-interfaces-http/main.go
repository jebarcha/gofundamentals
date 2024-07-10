package main

import (
	"fmt"
	"io"

	// "log"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	res, err := http.Get("http://google.com/")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999) //byte slice of 99999 elements
	// res.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, res.Body)
	//io.Copy(os.Stdout, res.Body)

	// body, err := io.ReadAll(res.Body)
	// res.Body.Close()
	// if res.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes", len(bs))

	return len(bs), nil
}
