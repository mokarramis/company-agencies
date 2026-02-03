package contract

import "companyAgancies/entity"

type AgencyInterface interface {
	CreateAgency(agency entity.Agency)
	GetAgency(agencyId int) (entity.Agency, error)
	ListRegionAgencies(region string) []entity.Agency
	GetAgencyStatus(region string) (int, int)
}
