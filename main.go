package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(stream io.ReadCloser) <-chan string {
	output := make(chan string)

	go func() {
		defer func() {
			err := stream.Close()
			if err != nil {
				log.Fatal("Error", err)
			}
		}()
		defer close(output)

		currentLine := ""

		for {
			data := make([]byte, 8)

			count, err := stream.Read(data)
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
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("Error", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Error", err)
	}

	lines := getLinesChannel(conn)

	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}
