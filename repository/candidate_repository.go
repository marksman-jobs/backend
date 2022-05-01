package repository

import (
	"context"
	"time"

	"github.com/marksman-jobs/backend/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type candidateRepositoryImpl struct {
	collection *mongo.Collection
}

func NewCandidateRepository(database *mongo.Database) CandidateRepository {
	return &candidateRepositoryImpl{
		collection: database.Collection("candidates"),
	}
}

func (repository *candidateRepositoryImpl) Insert(candidate entity.Candidate) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := repository.collection.InsertOne(ctx, bson.M{})

	if err != nil {
		return err
	}

	return nil

}

func (repository *candidateRepositoryImpl) FindAll() ([]entity.Candidate, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response := []entity.Candidate{}
	cursor, err := repository.collection.Find(ctx, bson.M{})
	if err != nil {
		return []entity.Candidate{}, err
	}

	err = cursor.All(ctx, response)
	if err != nil {
		return []entity.Candidate{}, nil
	}

	return response, nil

}

func (repository *candidateRepositoryImpl) Get(candidate entity.Candidate) (entity.Candidate, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := repository.collection.FindOne(ctx, &candidate)

	response := entity.Candidate{}
	err := result.Decode(&response)

	if err != nil {
		return entity.Candidate{}, err
	}

	return response, nil

}
