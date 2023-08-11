package service

import (
	"followPtong/model"
	"followPtong/repository"
)

type ServicePort interface {
	GetSer() (*model.GetResponse, error)
}

type serviceAdapter struct {
	r repository.RepositoryPort
}

func NewServiceAdapter(r repository.RepositoryPort) ServicePort {
	return serviceAdapter{r: r}
}

func (s serviceAdapter) GetSer() (*model.GetResponse, error) {
	businessTypes, organizationTypes, relations, err := s.r.Get()
	if err != nil {
		return nil, err
	}

	responses := model.GetResponse{ // assign values of variables
		BusinessType:     businessTypes,
		OrganizationType: organizationTypes,
		Relation:         relations,
	}

	return &responses, nil
}
