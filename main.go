package main

import (
	"companyAgancies/entity"
	"companyAgancies/handler"
	"companyAgancies/repository"
	"companyAgancies/service"
	"fmt"
)

func main() {
	var handler handler.Handler
	storage := entity.New("")
	repo := repository.Repo{
		StorageInterface: &storage,
	}
	handler.Service = service.New(repo)
	fmt.Println("Welcome to the company agencies app!")
	fmt.Println("--------------------------------------")
	region, command := handler.GetInputFlags()

	switch command {
	case "create":
		var agency = service.CreateAgencyRequest{
			Name:          "name",
			Region:        region,
			Address:       "address",
			Phone:         "0911",
			JoinedAt:      "1980",
			EmployeeCount: "458",
		}
		handler.Create(agency)
	case "list":
		var req = service.ListAgenciesRequest{
			Region: region,
		}
		fmt.Println(handler.List(req))
	case "get":
		var req = service.GetAgencyRequest{AgencyId: 0}
		fmt.Println(handler.Get(req))
	case "status":
		var req = service.StatusAgencyRequest{Region: region}
		agencyCount, employeeCount := handler.Status(req)
		fmt.Println(agencyCount, employeeCount)
	}
}
