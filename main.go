package main

import (
<<<<<<< HEAD
=======
	"bufio"
>>>>>>> 6fc3d783bce466975b9d0c02eb9e71552a6d047c
	"fmt"
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

<<<<<<< HEAD
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
=======
	scan := bufio.NewScanner(file)

	for scan.Scan() {
		line := scan.Text()
		fmt.Println(line)

	}
	if err := scan.Err(); err != nil {
		log.Fatal("err")
	}
>>>>>>> 6fc3d783bce466975b9d0c02eb9e71552a6d047c

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
