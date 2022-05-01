package service

import (
	"github.com/marksman-jobs/backend/entity"
	"github.com/marksman-jobs/backend/model"
	"github.com/marksman-jobs/backend/repository"
	"github.com/marksman-jobs/backend/validation"
	"go.mongodb.org/mongo-driver/bson"
)

type candidateServiceImpl struct {
	CandidateRespository repository.CandidateRepository
}

func NewCandidateService(candidateRepository *repository.CandidateRepository) CandidateService {
	return &candidateServiceImpl{
		CandidateRespository: *candidateRepository,
	}
}

func (service *candidateServiceImpl) List() (responses []model.GetCandidateResponse, err error) {

	// candidates, err := service.CandidateRespository.FindAll()
	_, err = service.CandidateRespository.FindAll()

	if err != nil {
		return []model.GetCandidateResponse{}, err
	}

	response := []model.GetCandidateResponse{}

	// TODO: Unmarshal candidate values into response

	// for _, candidate := range candidates {
	// 	response = append(response, model.GetCandidateResponse{
	// 		candidate.id,
	// 	})
	// }

	return response, nil
}

func (service *candidateServiceImpl) Get(candidateId string) (response model.GetCandidateResponse, err error) {

	request := entity.Candidate{}
	// request, err = service.CandidateRespository.Get(request)
	_, err = service.CandidateRespository.Get(request)

	if err != nil {
		return model.GetCandidateResponse{}, err
	}

	// TODO: Unmarshal request into model.GetCandidateResponse

	return model.GetCandidateResponse{}, nil

}

func (service *candidateServiceImpl) Create(requestBody []byte) (response model.CreateCandidateResponse, err error) {

	request := &model.CreateCandidateRequest{}
	bson.Unmarshal(requestBody, request)

	if err := validation.CandidateCreateValidation(*request); err != nil {
		return model.CreateCandidateResponse{}, err
	}

	candidate := &entity.Candidate{}

	// TODO: Unmarshal request values into candidate

	service.CandidateRespository.Insert(*candidate)

	return model.CreateCandidateResponse{}, nil

}
