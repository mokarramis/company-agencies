package deliveryparam

import "companyAgancies/service"

type Request struct {
	Command             string
	Region              string
	CreateAgencyRequest service.CreateAgencyRequest
	ListAgenciesRequest service.ListAgenciesRequest
	GetAgencyRequest    service.GetAgencyRequest
	StatusAgencyRequest service.StatusAgencyRequest
}
