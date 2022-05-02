package entity

type Company struct {
	CompanyId          string   `bun:"company_id,pk"`
	LocationIds        []string `bun:"location_ids,fk,array"`
	CompanyDescription string   `bun:"company_description,type:text"`
}
