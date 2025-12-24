package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal("Unable to open file")
		return
	}
	defer file.Close()

	buff := make([]byte, 8)

	for {
		readFile, err := file.Read(buff)
		if readFile > 0 {
			fmt.Println(string(buff[:readFile]))
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Println("Error Reading :(")
		}
	}

}
