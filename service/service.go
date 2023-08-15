package service

import (
	"CRM/model"
	"CRM/repository"
	"log"
)

type ServicePort interface {
	GetSer() (*model.GetResponse, error)
	Addser(req model.Addrequest) (*int64, error)
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

func (s serviceAdapter) Addser(req model.Addrequest) (*int64, error) {
	lastID, err := s.r.Add(req)
	if err != nil {
		log.Println(err.Error())
		return lastID, err
	}
	return lastID, nil
}
