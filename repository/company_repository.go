package repository

import (
	"context"
	"time"

	"github.com/marksman-jobs/backend/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type companyRepositoryImpl struct {
	collection *mongo.Collection
}

func NewCompanyRepository(database *mongo.Database) CompanyRepository {
	return &companyRepositoryImpl{
		collection: database.Collection("companies"),
	}
}

func (repository *companyRepositoryImpl) Insert(company entity.Company) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := repository.collection.InsertOne(ctx, bson.M{})

	if err != nil {
		return err
	}

	return nil

}

func (repository *companyRepositoryImpl) FindAll() ([]entity.Company, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response := []entity.Company{}
	cursor, err := repository.collection.Find(ctx, bson.M{})
	if err != nil {
		return []entity.Company{}, err
	}

	err = cursor.All(ctx, response)
	if err != nil {
		return []entity.Company{}, err
	}

	return response, nil

}

func (repository *companyRepositoryImpl) Get(company entity.Company) (entity.Company, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := repository.collection.FindOne(ctx, &company)

	response := entity.Company{}
	err := result.Decode(&response)

	if err != nil {
		return entity.Company{}, err
	}

	return response, nil

}
