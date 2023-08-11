package model

type BusinessType struct {
	ID           int    `db:"id" json:"id"`
	BusinessName string `db:"businessName" json:"businessName"`
}

type OrganizationType struct {
	ID               int    `json:"id"`
	OrganizationName string `json:"name"`
}

type Relation struct {
	ID           int    `json:"id"`
	RelationType string `json:"relationType"`
}

type GetResponse struct {
	BusinessType     []BusinessType     `json:"businessType"`
	OrganizationType []OrganizationType `json:"organizationType"`
	Relation         []Relation         `json:"relation"`
}
