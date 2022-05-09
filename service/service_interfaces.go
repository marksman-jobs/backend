package service

import (
	"github.com/marksman-jobs/backend/db"
)

type Service interface {
	CandidateList() (responses []db.Candidate, err error)
	CandidateGet(candidateId string) (response db.Candidate, err error)
	CandidateCreate(requestBody []byte) (response string, err error)

	CompanyList() (responses []db.Company, err error)
	CompanyGet(companyId string) (response db.Company, err error)
	CompanyCreate(requestBody []byte) (response string, err error)

	JobList() (responses []db.Job, err error)
	JobGet(jobId string) (response db.Job, err error)
	JobCreate(requestBody []byte) (response string, err error)
}

type ServiceImpl struct {
	postgres db.Queries
}

func NewService(postgres db.Queries) Service {
	return &ServiceImpl{postgres: postgres}
}
