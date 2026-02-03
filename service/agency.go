package service

import (
	"companyAgancies/contract"
	"companyAgancies/entity"
	"companyAgancies/repository"
)

type AgencyService struct {
	repo contract.AgencyInterface
}

type CreateAgencyRequest struct {
	Name          string
	Region        string
	Address       string
	Phone         string
	JoinedAt      string
	EmployeeCount string
}

type ListAgenciesRequest struct {
	Region string
}

type GetAgencyRequest struct {
	AgencyId int
}

type StatusAgencyRequest struct {
	Region string
}

func New(repo repository.Repo) AgencyService {
	return AgencyService{
		repo: repo,
	}
}

func (a AgencyService) CreateAgency(req CreateAgencyRequest) {
	var agency = entity.Agency{
		ID:            0,
		Name:          req.Name,
		Region:        req.Region,
		Address:       req.Address,
		Phone:         req.Phone,
		JoinedAt:      req.JoinedAt,
		EmployeeCount: req.EmployeeCount,
	}
	a.repo.CreateAgency(agency)
}

func (a AgencyService) ListRegionAgencies(region string) []entity.Agency {
	return a.repo.ListRegionAgencies(region)
}

func (a AgencyService) GetAgency(agencyId int) (entity.Agency, error) {
	return a.repo.GetAgency(agencyId)
}

func (a AgencyService) GetAgencyStatus(region string) (int, int) {
	return a.repo.GetAgencyStatus(region)
}
