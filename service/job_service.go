package service

import (
	"encoding/json"

	"github.com/marksman-jobs/backend/entity"
	"github.com/marksman-jobs/backend/model"
	"github.com/marksman-jobs/backend/repository"
	"github.com/marksman-jobs/backend/validation"
)

type jobServiceImpl struct {
	jobRepository repository.JobRepository
}

func NewJobService(jobRepository *repository.JobRepository) JobService {
	return &jobServiceImpl{
		jobRepository: *jobRepository,
	}
}

func (service *jobServiceImpl) List() (responses []model.GetJobResponse, err error) {

	jobs, err := service.jobRepository.FindAll()
	if err != nil {
		return []model.GetJobResponse{}, err
	}

	response := []model.GetJobResponse{}

	for _, job := range jobs {
		response = append(response, model.GetJobResponse{
			JobId:             job.JobId,
			JobTitle:          job.JobTitle,
			JobType:           job.JobType,
			JobFunction:       job.JobFunction,
			Currency:          job.Currency,
			Value:             job.Value,
			SkillsId:          job.SkillsId,
			JobApplicantIds:   job.JobApplicantIds,
			JobApplicantCount: job.JobApplicantCount,
			CompanyId:         job.CompanyId,
			RecruiterId:       job.RecruiterId,
			LocationId:        job.LocationId,
		})
	}

	return response, nil

}

func (service *jobServiceImpl) Get(jobId string) (response model.GetJobResponse, err error) {

	request := entity.Job{
		JobId: jobId,
	}

	job, err := service.jobRepository.Get(request)
	if err != nil {
		return model.GetJobResponse{}, err
	}

	response = model.GetJobResponse{
		JobId:             job.JobId,
		JobTitle:          job.JobTitle,
		JobType:           job.JobType,
		JobFunction:       job.JobFunction,
		Currency:          job.Currency,
		Value:             job.Value,
		SkillsId:          job.SkillsId,
		JobApplicantIds:   job.JobApplicantIds,
		JobApplicantCount: job.JobApplicantCount,
		CompanyId:         job.CompanyId,
		RecruiterId:       job.RecruiterId,
		LocationId:        job.LocationId,
	}

	return response, nil

}

func (service *jobServiceImpl) Create(requestBody []byte) (response model.CreateJobResponse, err error) {

	request := &model.CreateJobRequest{}
	json.Unmarshal(requestBody, request)

	if err := validation.JobCreateValidation(*request); err != nil {
		return model.CreateJobResponse{}, err
	}

	job := entity.Job{
		JobTitle:          request.JobTitle,
		JobType:           request.JobType,
		JobFunction:       request.JobFunction,
		Currency:          request.Currency,
		Value:             request.Value,
		SkillsId:          request.SkillsId,
		JobApplicantIds:   request.JobApplicantIds,
		JobApplicantCount: request.JobApplicantCount,
		CompanyId:         request.CompanyId,
		RecruiterId:       request.RecruiterId,
		LocationId:        request.LocationId,
	}

	jobId, err := service.jobRepository.Insert(job)
	if err != nil {
		return model.CreateJobResponse{}, err
	}

	return model.CreateJobResponse{JobId: jobId}, nil

}
