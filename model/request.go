package model

type BusinessType struct {
	ID           int    `db:"id" json:"id"`
	BusinessName string `db:"businessName" json:"businessName"`
}

type OrganizationType struct {
	ID               int    `db:"id" json:"id"`
	OrganizationName string `db:"organizationName" json:"organizationName"`
}

type Relation struct {
	ID           int    `db:"id" json:"id"`
	RelationType string `db:"relationType" json:"relationType"`
}

type GetResponse struct {
	BusinessType     []BusinessType     `json:"businessType"`
	OrganizationType []OrganizationType `json:"organizationType"`
	Relation         []Relation         `json:"relation"`
}

type Addrequest struct {
	OrganizationType uint   `db:"oraganizationType" json:"oraganizationType"`
	AliasName        string `db:"aliasName" json:"aliasName"`
	CompanyName      string `db:"companyName" json:"companyName"`
	BusinessType     uint   `db:"businessType" json:"businessType"`
	Domain           string `db:"domain" json:"domain"`
	WebSite          string `db:"webSite" json:"webSite"`
	Contact          string `db:"contact" json:"contact"`
	ContactEmail     string `db:"contactEmail" json:"contactEmail"`
	ContactPhone     string `db:"contactPhone" json:"contactPhone"`
	Relation         uint   `db:"relation" json:"relation"`
}
