package repository

import (
	"companyAgancies/contract"
	"companyAgancies/entity"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type Repo struct {
	StorageInterface contract.StorageInterface
}

func (r Repo) CreateAgency(agency entity.Agency) {
	// convert to json
	jsonAgency, err := json.Marshal(agency) // creates slice of bytes
	if err != nil {
		fmt.Println("there is an error in marshalling", err)
	}
	// store data in a file
	r.StorageInterface.StoreIntoFile(jsonAgency)
}

func (r Repo) ListRegionAgencies(region string) []entity.Agency {
	var agencyList []entity.Agency
	// read-from-file with region id
	agencies := r.StorageInterface.ReadFromFile()
	// return lists of data
	for _, u := range agencies {
		agencyStruct := entity.Agency{}
		json.Unmarshal([]byte(u), &agencyStruct)
		if agencyStruct.Region != region {
			continue
		}
		agencyList = append(agencyList, agencyStruct)
	}
	return agencyList
}

func (r Repo) GetAgency(id int) (entity.Agency, error) {
	// read-from-file with region id
	agencies := r.StorageInterface.ReadFromFile()
	agencyStruct := entity.Agency{}
	for _, u := range agencies {
		json.Unmarshal([]byte(u), &agencyStruct)
		if agencyStruct.ID == id {
			return agencyStruct, nil
		}
	}
	return entity.Agency{}, errors.New("there is no agency in this region")
}

func (r Repo) GetAgencyStatus(region string) (int, int) {
	var agencyList []entity.Agency
	count := 0
	// read-from-file with region id
	agencies := r.StorageInterface.ReadFromFile()
	// return lists of data
	for _, u := range agencies {
		agencyStruct := entity.Agency{}
		json.Unmarshal([]byte(u), &agencyStruct)
		if agencyStruct.Region != region {
			continue
		}
		agencyList = append(agencyList, agencyStruct)
	}
	for _, agency := range agencyList {
		employeeCount, _ := strconv.Atoi(agency.EmployeeCount)
		count += employeeCount
	}
	return len(agencyList), count
}
