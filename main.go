package main

import (
	"bufio"
	"fmt"
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

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		line := scan.Text()
		fmt.Println(line)

	}
	if err := scan.Err(); err != nil {
		log.Fatal("err")
	}

}
