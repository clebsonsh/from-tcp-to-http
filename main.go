package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	output := make(chan string)

	go func() {
		defer func() {
			err := f.Close()
			if err != nil {
				log.Fatal("Error", err)
			}
		}()
		defer close(output)

		currentLine := ""

		for {
			data := make([]byte, 8)

			count, err := f.Read(data)
			if err != nil {
				break
			}

			parts := strings.Split(string(data[:count]), "\n")

			currentLine += parts[0]

			if len(parts) > 1 {
				output <- currentLine
				currentLine = parts[1]
			}
		}
	}()

	return output
}

func main() {
	fileHandler, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Error", err)
	}

	lines := getLinesChannel(fileHandler)

	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}
