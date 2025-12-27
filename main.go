package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal("Unable to open file")
		return
	}
	defer file.Close()

	buf := make([]byte, 8)
	currentLine := ""

	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal()
		}

		chunks := string(buf[:n])
		parts := strings.Split(chunks, "\n")

		for i := 0; i < len(parts)-1; i++ {
			fmt.Printf("read: %s\n", currentLine+parts[i])
			currentLine = ""
		}
		currentLine += parts[len(parts)-1]
	}
	if currentLine != "" {
		fmt.Printf("read: %s\n", currentLine)
	}
}
