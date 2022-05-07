package service

import "github.com/marksman-jobs/backend/model"

type CandidateService interface {
	List() (responses []model.GetCandidateResponse, err error)
	Get(candidateId string) (response model.GetCandidateResponse, err error)
	Create(requestBody []byte) (response model.CreateCandidateResponse, err error)
}

type CompanyService interface {
	List() (responses []model.GetCompanyResponse, err error)
	Get(companyId string) (response model.GetCompanyResponse, err error)
	Create(requestBody []byte) (response model.CreateCompanyResponse, err error)
}

type JobService interface {
	List() (responses []model.GetJobResponse, err error)
	Get(jobId string) (response model.GetJobResponse, err error)
	Create(requestBody []byte) (response model.CreateJobResponse, err error)
}
