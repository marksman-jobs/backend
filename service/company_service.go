package service

import (
	"context"
	"database/sql"

	"github.com/marksman-jobs/backend/db"
)

func (service *ServiceImpl) CompanyList() (responses []db.Company, err error) {

	companies, err := service.postgres.ListCompanies(context.Background())

	if err != nil {
		return []db.Company{}, err
	}

	return companies, nil
}

func (service *ServiceImpl) CompanyGet(candidateId string) (response db.Company, err error) {

	company, err := service.postgres.GetCompany(context.Background(), candidateId)
	if err != nil {
		return db.Company{}, err
	}

	return company, nil

}

func (service *ServiceImpl) CompanyCreate(requestBody []byte) (response string, err error) {

	err = service.postgres.CreateCompany(context.Background(), sql.NullString{String: string(requestBody), Valid: true})
	if err != nil {
		return "", err
	}

	// TODO: Return id here
	return "", nil

}
