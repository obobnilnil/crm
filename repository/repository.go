package repository

import (
	"CRM/model"
	"database/sql"
	"log"
)

type RepositoryPort interface {
	Get() ([]model.BusinessType, []model.OrganizationType, []model.Relation, error)
	Add(req model.Addrequest) (*int64, error)
}

type repositoryAdapter struct {
	db *sql.DB
}

func NewRepositoryAdapter(db *sql.DB) RepositoryPort {
	return repositoryAdapter{db: db}
}

func (r repositoryAdapter) Get() ([]model.BusinessType, []model.OrganizationType, []model.Relation, error) {
	var businessTypes []model.BusinessType
	var organizationTypes []model.OrganizationType
	var relations []model.Relation

	// Query for business types
	businessQuery := "SELECT id, businessName FROM businessType"
	businessRows, err := r.db.Query(businessQuery)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}
	defer businessRows.Close()

	for businessRows.Next() {
		var businessType model.BusinessType
		err = businessRows.Scan(&businessType.ID, &businessType.BusinessName)
		if err != nil {
			log.Println(err)
			return nil, nil, nil, err
		}
		businessTypes = append(businessTypes, businessType)
	}

	// Query for organization types
	organizationQuery := "SELECT id, organizationName FROM organizationType"
	organizationRows, err := r.db.Query(organizationQuery)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}
	defer organizationRows.Close()

	for organizationRows.Next() {
		var organizationType model.OrganizationType
		err = organizationRows.Scan(&organizationType.ID, &organizationType.OrganizationName)
		if err != nil {
			log.Println(err)
			return nil, nil, nil, err
		}
		organizationTypes = append(organizationTypes, organizationType)
	}

	// Query for relations
	relationQuery := "SELECT id, relationType FROM relation"
	relationRows, err := r.db.Query(relationQuery)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}
	defer relationRows.Close()

	for relationRows.Next() {
		var relation model.Relation
		err = relationRows.Scan(&relation.ID, &relation.RelationType)
		if err != nil {
			log.Println(err)
			return nil, nil, nil, err
		}
		relations = append(relations, relation)
	}

	return businessTypes, organizationTypes, relations, nil
}

func (r repositoryAdapter) Add(req model.Addrequest) (*int64, error) {
	query := "INSERT INTO newDeal (organizationType, aliasName, companyName, businessType, domain, webSite, contact, contactEmail, contactPhone, relation) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, req.OrganizationType, req.AliasName, req.CompanyName, req.BusinessType, req.Domain, req.WebSite, req.Contact, req.ContactEmail, req.ContactPhone, req.Relation)
	if err != nil {
		return nil, err
	}
	lastID, _ := result.LastInsertId()
	return &lastID, nil
}
