package repository

import "github.com/marksman-jobs/backend/entity"

type CandidateRepository interface {
	Insert(candidate entity.Candidate) (string, error)
	FindAll() ([]entity.Candidate, error)
	Get(candidate entity.Candidate) (entity.Candidate, error)
}

type CompanyRepository interface {
	Insert(company entity.Company) (string, error)
	FindAll() ([]entity.Company, error)
	Get(company entity.Company) (entity.Company, error)
}
