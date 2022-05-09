package dto

type CreateCompanyRequest struct {
	LocationIds        []string `bun:"location_ids,fk,array"`
	CompanyDescription string   `bun:"company_description,type:text"`
}

type CreateCompanyResponse struct {
	CompanyId string `bun:"company_id,pk"`
}

type GetCompanyRequest struct {
	CompanyId string `bun:"company_id,pk"`
}

type GetCompanyResponse struct {
	CompanyId          string   `bun:"company_id,pk"`
	LocationIds        []string `bun:"location_ids,fk,array"`
	CompanyDescription string   `bun:"company_description,type:text"`
}
