package service

import (
	"encoding/json"

	"github.com/marksman-jobs/backend/entity"
	"github.com/marksman-jobs/backend/model"
	"github.com/marksman-jobs/backend/repository"
	"github.com/marksman-jobs/backend/validation"
)

type candidateServiceImpl struct {
	candidateRepository repository.CandidateRepository
}

func NewCandidateService(candidateRepository *repository.CandidateRepository) CandidateService {
	return &candidateServiceImpl{
		candidateRepository: *candidateRepository,
	}
}

func (service *candidateServiceImpl) List() (responses []model.GetCandidateResponse, err error) {

	candidates, err := service.candidateRepository.FindAll()

	if err != nil {
		return []model.GetCandidateResponse{}, err
	}

	response := []model.GetCandidateResponse{}

	for _, candidate := range candidates {
		response = append(response, model.GetCandidateResponse{
			FirstName:          candidate.FirstName,
			LastName:           candidate.LastName,
			CurrentCompanyId:   candidate.CurrentCompanyId,
			CurrentRoleId:      candidate.CurrentRoleId,
			CurrentLocationId:  candidate.CurrentLocationId,
			DesiredLocationIds: candidate.DesiredLocationIds,
			DesiredRoleId:      candidate.DesiredRoleId,
			Email:              candidate.Email,
			PasswordHash:       candidate.PasswordHash,
			IsEmailVerified:    candidate.IsEmailVerified,
			Pronouns:           candidate.Pronouns,
			EducationId:        candidate.EducationId,
			SkillsId:           candidate.SkillsId,
			ResumeId:           candidate.ResumeId,
		})
	}

	return response, nil
}

func (service *candidateServiceImpl) Get(candidateId string) (response model.GetCandidateResponse, err error) {

	request := entity.Candidate{CandidateId: candidateId}

	data, err := service.candidateRepository.Get(request)
	if err != nil {
		return model.GetCandidateResponse{}, err
	}

	response = model.GetCandidateResponse{
		FirstName:          data.FirstName,
		LastName:           data.LastName,
		CurrentCompanyId:   data.CurrentCompanyId,
		CurrentRoleId:      data.CurrentRoleId,
		CurrentLocationId:  data.CurrentLocationId,
		DesiredLocationIds: data.DesiredLocationIds,
		DesiredRoleId:      data.DesiredRoleId,
		Email:              data.Email,
		PasswordHash:       data.PasswordHash,
		IsEmailVerified:    data.IsEmailVerified,
		Pronouns:           data.Pronouns,
		EducationId:        data.EducationId,
		SkillsId:           data.SkillsId,
		ResumeId:           data.ResumeId,
	}
	return model.GetCandidateResponse{}, nil

}

func (service *candidateServiceImpl) Create(requestBody []byte) (response model.CreateCandidateResponse, err error) {

	request := &model.CreateCandidateRequest{}
	json.Unmarshal(requestBody, request)

	if err := validation.CandidateCreateValidation(*request); err != nil {
		return model.CreateCandidateResponse{}, err
	}

	candidate := entity.Candidate{
		FirstName:          request.FirstName,
		LastName:           request.LastName,
		CurrentCompanyId:   request.CurrentCompanyId,
		CurrentRoleId:      request.CurrentRoleId,
		CurrentLocationId:  request.CurrentLocationId,
		DesiredLocationIds: request.DesiredLocationIds,
		DesiredRoleId:      request.DesiredRoleId,
		Email:              request.Email,
		PasswordHash:       request.PasswordHash,
		Pronouns:           request.Pronouns,
		EducationId:        request.EducationId,
		SkillsId:           request.SkillsId,
		ResumeId:           request.ResumeId,
	}

	candidateId, err := service.candidateRepository.Insert(candidate)
	if err != nil {
		return model.CreateCandidateResponse{}, err
	}

	return model.CreateCandidateResponse{CandidateId: candidateId}, nil

}
