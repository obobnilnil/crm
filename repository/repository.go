package repository

import (
	"CRM/model"
	"database/sql"
	"log"
)

type RepositoryPort interface {
	Get() ([]model.BusinessType, []model.OrganizationType, []model.Relation, error)
	GetbyDomain(domainUrl string) (*model.GetResponsebyDomain, error)
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
	businessQuery := "SELECT businessID, businessName FROM businessType"
	businessRows, err := r.db.Query(businessQuery)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}
	defer businessRows.Close()

	for businessRows.Next() {
		var businessType model.BusinessType
		err = businessRows.Scan(&businessType.BusinessID, &businessType.BusinessName)
		if err != nil {
			log.Println(err)
			return nil, nil, nil, err
		}
		businessTypes = append(businessTypes, businessType)
	}

	// Query for organization types
	organizationQuery := "SELECT organizationID, organizationName FROM organizationType"
	organizationRows, err := r.db.Query(organizationQuery)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}
	defer organizationRows.Close()

	for organizationRows.Next() {
		var organizationType model.OrganizationType
		err = organizationRows.Scan(&organizationType.OrganizationID, &organizationType.OrganizationName)
		if err != nil {
			log.Println(err)
			return nil, nil, nil, err
		}
		organizationTypes = append(organizationTypes, organizationType)
	}

	// Query for relations
	relationQuery := "SELECT id, relationID FROM relation"
	relationRows, err := r.db.Query(relationQuery)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}
	defer relationRows.Close()

	for relationRows.Next() {
		var relation model.Relation
		err = relationRows.Scan(&relation.RelationID, &relation.RelationType)
		if err != nil {
			log.Println(err)
			return nil, nil, nil, err
		}
		relations = append(relations, relation)
	}

	return businessTypes, organizationTypes, relations, nil
}

func (r repositoryAdapter) GetbyDomain(domainUrl string) (*model.GetResponsebyDomain, error) {
	query := `
	SELECT
		newOrganization.newOrganizationID,
		organizationType.organizationName,
		newOrganization.aliasName,
		newOrganization.companyNameEN,
		businessType.businessName,
		newOrganization.domain,
		newOrganization.webSite,
		newOrganization.contact,
		newOrganization.contactEmail,
		newOrganization.contactPhone,
		relation.relationType
	FROM newOrganization
	JOIN businessType ON newOrganization.businessID = businessType.businessID
	JOIN organizationType ON newOrganization.organizationID = organizationType.organizationID
	JOIN relation ON newOrganization.relationID = relation.relationID
	WHERE newOrganization.domain = ?
`
	var responseDomain model.GetResponsebyDomain
	err := r.db.QueryRow(query, domainUrl).Scan(&responseDomain.NewOrganizationID, &responseDomain.OrganizationID, &responseDomain.AliasName, &responseDomain.CompanyNameEN, &responseDomain.BusinessID, &responseDomain.Domain, &responseDomain.WebSite, &responseDomain.Contact, &responseDomain.ContactEmail, &responseDomain.ContactPhone, &responseDomain.RelationID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//log.Println("respository", responseDomain)
	return &responseDomain, nil
}

func (r repositoryAdapter) Add(req model.Addrequest) (*int64, error) {
	query := "INSERT INTO newOrganization (organizationID, aliasName, companyNameEN, businessID, domain, webSite, contact, contactEmail, contactPhone, relationID) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, req.OrganizationID, req.AliasName, req.CompanyNameEN, req.BusinessID, req.Domain, req.WebSite, req.Contact, req.ContactEmail, req.ContactPhone, req.RelationID)
	if err != nil {
		return nil, err
	}
	lastID, _ := result.LastInsertId()
	return &lastID, nil
}
