package main

import (
	"companyAgancies/delivery/deliveryparam"
	"companyAgancies/entity"
	"companyAgancies/handler"
	"companyAgancies/repository"
	"companyAgancies/service"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func main() {
	const (
		network = "tcp"
		address = ":9980"
	)

	listener, err := net.Listen(network, address)
	fmt.Println("listening to the connection...")

	if err != nil {
		log.Fatalln("there is error in listening...")
	}

	for {
		conn, aErr := listener.Accept()
		if aErr != nil {
			log.Fatalln("there is an error in connecting to client...")
			continue
		}
		var data = make([]byte, 1024)

		numberOfReadBytes, rErr := conn.Read(data)

		if rErr != nil {
			log.Fatalln("can not read from the client...")
			continue
		}

		fmt.Printf("data is: %s", string(data))
		var req = &deliveryparam.Request{}

		if error := json.Unmarshal(data[:numberOfReadBytes], req); error != nil {
			log.Println("there is an error")
			continue
		}

		storage := entity.New("")
		var handler handler.Handler
		repo := repository.Repo{
			StorageInterface: &storage,
		}
		handler.Service = service.New(repo)
		fmt.Println(handler.Service)

		switch req.Command {
		case "create":
			handler.Create(req.CreateAgencyRequest)
		case "list":
			fmt.Println(handler.List(req.ListAgenciesRequest))
		case "get":
			fmt.Println(handler.Get(req.GetAgencyRequest))
		case "status":
			agencyCount, employeeCount := handler.Status(req.StatusAgencyRequest)
			fmt.Println(agencyCount, employeeCount)
		}

	}
}
