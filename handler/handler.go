package handler

import (
	"companyAgancies/entity"
	"companyAgancies/service"
	"flag"
	"fmt"
)

func (h Handler) GetInputFlags() (string, string) {
	region := flag.String("region", "there is no flag", "region flag")
	command := flag.String("command", "there is no command", "command flag to work on")
	flag.Parse()

	return *region, *command
}

type Handler struct {
	Service service.AgencyService
}

func (h Handler) Create(request service.CreateAgencyRequest) {

	var createAgencyRequest = service.CreateAgencyRequest{
		Name:          request.Name,
		Region:        request.Region,
		Address:       request.Address,
		Phone:         request.Phone,
		JoinedAt:      request.JoinedAt,
		EmployeeCount: request.EmployeeCount,
	}
	h.Service.CreateAgency(createAgencyRequest)

	fmt.Println("data created successfully")
}

func (h Handler) List(req service.ListAgenciesRequest) []entity.Agency {
	region := req.Region

	return h.Service.ListRegionAgencies(region)
}

func (h Handler) Get(req service.GetAgencyRequest) (entity.Agency, error) {
	AgencyId := req.AgencyId
	return h.Service.GetAgency(AgencyId)
}

func (h Handler) Status(req service.StatusAgencyRequest) (int, int) {
	region := req.Region

	return h.Service.GetAgencyStatus(region)
}
