package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// fmt.Println(len(os.Args))
	if len(os.Args) == 1 {
		fmt.Println("No input file")
	} else if len(os.Args) > 2 {
		fmt.Println("Invalid number of argument")
	} else {
		file, err := os.Open(os.Args[1])

		if err != nil {
			log.Fatal(err)
		} else {
			io.Copy(os.Stdout, file)
		}
	}
}
