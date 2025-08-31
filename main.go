package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileHandler, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Error", err)
	}

	for {
		data := make([]byte, 8)

		count, err := fileHandler.Read(data)
		if err != nil {
			break
		}

		fmt.Printf("read: %s\n", string(data[:count]))
	}
}
