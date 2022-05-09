package service

import (
	"context"
	"encoding/json"

	"github.com/marksman-jobs/backend/db"
	"github.com/marksman-jobs/backend/validation"
)

func (service *ServiceImpl) JobList() (responses []db.Job, err error) {

	jobs, err := service.postgres.ListJobs(context.Background())

	if err != nil {
		return []db.Job{}, err
	}

	return jobs, nil
}

func (service *ServiceImpl) JobGet(candidateId string) (response db.Job, err error) {

	job, err := service.postgres.GetJob(context.Background(), candidateId)
	if err != nil {
		return db.Job{}, err
	}

	return job, nil

}

func (service *ServiceImpl) JobCreate(requestBody []byte) (response string, err error) {

	request := &db.CreateJobParams{}
	json.Unmarshal(requestBody, request)

	if err := validation.JobCreateValidation(*request); err != nil {
		return "", err
	}

	err = service.postgres.CreateJob(context.Background(), *request)
	if err != nil {
		return "", err
	}

	// TODO: Return id here
	return "", nil

}
