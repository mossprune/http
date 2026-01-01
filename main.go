package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {

	lineChan := make(chan string, 1)

	buf := make([]byte, 8)
	currentLine := ""

	go func() {
		defer f.Close()
		defer close(lineChan)
		for {
			n, err := f.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal()
			}

			chunks := string(buf[:n])
			parts := strings.Split(chunks, "\n")

			for i := 0; i < len(parts)-1; i++ {
				lineChan <- (string(currentLine + parts[i]))
				currentLine = ""
			}
			currentLine += parts[len(parts)-1]
		}

		if currentLine != "" {
			lineChan <- (string(currentLine))
		}
	}()

	return lineChan
}

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal("Unable to open file")
	}

	lineChan := getLinesChannel(file)

	for val := range lineChan {
		fmt.Printf("read: %s\n", val)
	}
}
