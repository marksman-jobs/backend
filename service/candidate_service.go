package service

import (
	"context"
	"encoding/json"

	"github.com/marksman-jobs/backend/db"
	"github.com/marksman-jobs/backend/validation"
)

func (service *ServiceImpl) CandidateList() (responses []db.Candidate, err error) {

	candidates, err := service.postgres.ListCandidates(context.Background())

	if err != nil {
		return []db.Candidate{}, err
	}

	return candidates, nil
}

func (service *ServiceImpl) CandidateGet(candidateId string) (response db.Candidate, err error) {

	candidate, err := service.postgres.GetCandidate(context.Background(), candidateId)
	if err != nil {
		return db.Candidate{}, err
	}

	return candidate, nil

}

func (service *ServiceImpl) CandidateCreate(requestBody []byte) (response string, err error) {

	request := &db.CreateCandidateParams{}
	json.Unmarshal(requestBody, request)

	if err := validation.CandidateCreateValidation(*request); err != nil {
		return "", err
	}

	err = service.postgres.CreateCandidate(context.Background(), *request)
	if err != nil {
		return "", err
	}

	// TODO: Return id here
	return "", nil

}
