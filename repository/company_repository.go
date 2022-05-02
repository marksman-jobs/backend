package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/marksman-jobs/backend/entity"
)

type companyRepositoryImpl struct {
	pgconn *pgxpool.Pool
}

func NewCompanyRepository(connPool *pgxpool.Pool) CompanyRepository {
	return &companyRepositoryImpl{
		pgconn: connPool,
	}
}

func (repository *companyRepositoryImpl) Insert(company entity.Company) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// TODO: Add query to insert company here
	_, err := repository.pgconn.Query(ctx, "")

	if err != nil {
		return err
	}

	return nil

}

func (repository *companyRepositoryImpl) FindAll() ([]entity.Company, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response := []entity.Company{}
	rows, err := repository.pgconn.Query(ctx, "SELECT * FROM companies")
	if err != nil {
		return []entity.Company{}, err
	}

	for rows.Next() {
		res := entity.Company{}
		err = rows.Scan(res)

		if err != nil {
			return []entity.Company{}, err
		}

		response = append(response, res)
	}

	return response, nil

}

func (repository *companyRepositoryImpl) Get(company entity.Company) (entity.Company, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	row := repository.pgconn.QueryRow(ctx, "SELECT * FROM companies WHERE company_id = $1", company.CompanyId)

	response := entity.Company{}
	err := row.Scan(response)

	if err != nil {
		return entity.Company{}, err
	}

	return response, nil

}
