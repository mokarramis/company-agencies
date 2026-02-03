package main

import (
	"companyAgancies/delivery/deliveryparam"
	"companyAgancies/service"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("command", os.Args[0])
	message := "default"
	if len(os.Args) > 1 {
		message = os.Args[1]
	}

	const (
		network = "tcp"
		address = "127.0.0.1:9980"
	)
	conn, error := net.Dial(network, address)

	if error != nil {
		log.Println("there is an error", error)
	}
	req := deliveryparam.Request{Command: message}

	if req.Command == "create" {
		var region, name, address, phone, joinedAt, employeeCount string

		fmt.Println("we want to create agency and need some data,")
		fmt.Println("please enter the name of agency")
		_, err := fmt.Scan(&name)
		if err != nil {
			fmt.Println("there is no data to scan")
		}
		fmt.Println("please enter the address of agency")
		_, aErr := fmt.Scan(&address)
		if aErr != nil {
			fmt.Println("there is no data to scan")
		}
		fmt.Println("please enter the phone of agency")
		_, pErr := fmt.Scan(&phone)
		if pErr != nil {
			fmt.Println("there is no data to scan")
		}
		fmt.Println("please enter the joined at of agency")
		_, jErr := fmt.Scan(&joinedAt)
		if jErr != nil {
			fmt.Println("there is no data to scan")
		}
		fmt.Println("please enter the employee count of agency")
		_, eErr := fmt.Scan(&employeeCount)
		if eErr != nil {
			fmt.Println("there is no data to scan")
		}
		fmt.Println("please enter the region agency")
		_, rErr := fmt.Scan(&region)
		if rErr != nil {
			fmt.Println("there is no data to scan")
		}
		fmt.Println(region)
		req.CreateAgencyRequest = service.CreateAgencyRequest{
			Name:          name,
			Region:        region,
			Address:       address,
			Phone:         phone,
			JoinedAt:      joinedAt,
			EmployeeCount: employeeCount,
		}
	}

	if req.Command == "list" {
		var region string

		fmt.Println("please enter the region agency")
		_, rErr := fmt.Scan(&region)
		if rErr != nil {
			fmt.Println("there is no data to scan")
		}
		req.ListAgenciesRequest = service.ListAgenciesRequest{
			Region: region,
		}
	}

	if req.Command == "get" {
		var agencyId int

		fmt.Println("please enter the region agency")
		_, rErr := fmt.Scan(&agencyId)
		if rErr != nil {
			fmt.Println("there is no data to scan")
		}
		req.GetAgencyRequest = service.GetAgencyRequest{
			AgencyId: agencyId,
		}
	}

	if req.Command == "status" {
		var agencyId int

		fmt.Println("please enter the region agency")
		_, rErr := fmt.Scan(&agencyId)
		if rErr != nil {
			fmt.Println("there is no data to scan")
		}
		req.GetAgencyRequest = service.GetAgencyRequest{
			AgencyId: agencyId,
		}
	}

	serializedData, mErr := json.Marshal(&req)
	if mErr != nil {
		log.Fatalln("can't marshal data", mErr)
	}

	numberOfWrittenData, error := conn.Write(serializedData)
	fmt.Println("numberOfWrittenData", numberOfWrittenData)

	var data = make([]byte, 1024)
	_, rError := conn.Read(data)

	if rError != nil {
		log.Fatalln("teher is an error", rError)
	}
}
