package service

import (
	"github.com/marksman-jobs/backend/entity"
	"github.com/marksman-jobs/backend/model"
	"github.com/marksman-jobs/backend/repository"
	"github.com/marksman-jobs/backend/validation"
	"go.mongodb.org/mongo-driver/bson"
)

type companyServiceImpl struct {
	companyRepository repository.CompanyRepository
}

func NewCompanyService(companyRepository *repository.CompanyRepository) CompanyService {
	return &companyServiceImpl{
		companyRepository: *companyRepository,
	}
}

func (service *companyServiceImpl) List() (responses []model.GetCompanyResponse, err error) {

	// companies, err := service.companyRepository.FindAll()
	_, err = service.companyRepository.FindAll()

	if err != nil {
		return []model.GetCompanyResponse{}, err
	}

	response := []model.GetCompanyResponse{}

	// TODO: Unmarshal company values into response

	// for _, company := range companies {
	// 	response = append(response, model.GetCompanyResponse{
	// 		company.id,
	// 	})
	// }

	return response, nil

}

func (service *companyServiceImpl) Get(companyId string) (response model.GetCompanyResponse, err error) {

	request := entity.Company{}
	// request := entity.Company{
	// 	company.Id: companyId
	// }
	// request, err := service.companyRepository.Get(request)

	_, err = service.companyRepository.Get(request)

	if err != nil {
		return model.GetCompanyResponse{}, nil
	}

	// TODO: Unmarshal request into model.GetCompanyResponse

	return model.GetCompanyResponse{}, nil

}

func (service *companyServiceImpl) Create(requestBody []byte) (response model.CreateCompanyResponse, err error) {

	request := &model.CreateCompanyRequest{}
	bson.Unmarshal(requestBody, request)

	if err := validation.CompanyCreateValidation(*request); err != nil {
		return model.CreateCompanyResponse{}, err
	}

	company := &entity.Company{}

	// TODO: Unmarshal request values into candidate

	service.companyRepository.Insert(*company)

	return model.CreateCompanyResponse{}, nil

}
