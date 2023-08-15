package model

type BusinessType struct {
	BusinessID   int    `db:"businessID" json:"businessID"`
	BusinessName string `db:"businessName" json:"businessName"`
}

type OrganizationType struct {
	OrganizationID   int    `db:"organizationID" json:"organizationID"`
	OrganizationName string `db:"organizationName" json:"organizationName"`
}

type Relation struct {
	RelationID   int    `db:"relationID" json:"relationID"`
	RelationType string `db:"relationType" json:"relationType"`
}

type GetResponse struct {
	BusinessType     []BusinessType     `json:"businessType"`
	OrganizationType []OrganizationType `json:"organizationType"`
	Relation         []Relation         `json:"relation"`
}

type GetResponsebyDomain struct {
	NewOrganizationID int    `db:"newOrganizationID" json:"newOrganizationID"`
	OrganizationID    string `db:"organizationID" json:"organizationName"`
	AliasName         string `db:"aliasName" json:"aliasName"`
	CompanyNameEN     string `db:"companyNameEN" json:"companyNameEN"`
	BusinessID        string `db:"businessID" json:"businessName"`
	Domain            string `db:"domain" json:"domain"`
	WebSite           string `db:"webSite" json:"webSite"`
	Contact           string `db:"contact" json:"contact"`
	ContactEmail      string `db:"contactEmail" json:"contactEmail"`
	ContactPhone      string `db:"contactPhone" json:"contactPhone"`
	RelationID        string `db:"relationID" json:"relationType"`
}

type Addrequest struct { // ไม่ต้องใส่แปลง db ก็ได้ แต่ใส่ให้เรารู้ว่า model ใช้งานที่ repository
	OrganizationID int    `db:"organizationID" json:"organizationID"`
	AliasName      string `db:"aliasName" json:"aliasName"`
	CompanyNameEN  string `db:"companyNameEN" json:"companyNameEN"`
	BusinessID     int    `db:"businessID" json:"businessID"`
	Domain         string `db:"domain" json:"domain"`
	WebSite        string `db:"webSite" json:"webSite"`
	Contact        string `db:"contact" json:"contact"`
	ContactEmail   string `db:"contactEmail" json:"contactEmail"`
	ContactPhone   string `db:"contactPhone" json:"contactPhone"`
	RelationID     int    `db:"relationID" json:"relationID"`
}
