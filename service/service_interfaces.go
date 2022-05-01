package service

import "github.com/marksman-jobs/backend/model"

type CandidateService interface {
	List() (responses []model.GetCandidateResponse, err error)
	Get(candidateId string) (response model.GetCandidateResponse, err error)
	Create(requestBody []byte) (response model.CreateCandidateResponse, err error)
}
