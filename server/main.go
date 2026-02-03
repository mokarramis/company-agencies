package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	const (
		address = "127.0.0.1:2022"
		network = "tcp"
	)

	listener, err := net.Listen(network, address)
	fmt.Println("listener created...")

	if err != nil {
		log.Println("there can not be connected", address, network)
	}

	for {
		conn, aErr := listener.Accept()
		if aErr != nil {
			log.Println("there can not create a connection", aErr)
			continue
		}

		data := make([]byte, 1024)
		numberOfBytes, cErr := conn.Read(data)
		if cErr != nil {
			log.Println("there is an error in reading bytes", cErr)
			continue
		}
		fmt.Println(numberOfBytes)
		fmt.Println(string(data))

		defer listener.Close()

	}

}
