package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/marksman-jobs/backend/entity"
)

type candidateRepositoryImpl struct {
	pgconn *pgxpool.Pool
}

func NewCandidateRepository(connPool *pgxpool.Pool) CandidateRepository {
	return &candidateRepositoryImpl{
		pgconn: connPool,
	}
}

func (repository *candidateRepositoryImpl) Insert(candidate entity.Candidate) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// TODO: Write query to insert candidate here
	row := repository.pgconn.QueryRow(ctx, "INSERT INTO candidates ")

	inserted := entity.Candidate{}
	if err := row.Scan(inserted); err != nil {
		return "", err
	}

	return inserted.CandidateId, nil

}

func (repository *candidateRepositoryImpl) FindAll() ([]entity.Candidate, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response := []entity.Candidate{}
	rows, err := repository.pgconn.Query(ctx, "SELECT * FROM candidates")
	if err != nil {
		return []entity.Candidate{}, err
	}

	for rows.Next() {
		res := entity.Candidate{}
		err = rows.Scan(res)

		if err != nil {
			return []entity.Candidate{}, err
		}

		response = append(response, res)
	}

	return response, nil

}

func (repository *candidateRepositoryImpl) Get(candidate entity.Candidate) (entity.Candidate, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	row := repository.pgconn.QueryRow(ctx, "SELECT * FROM candidates WHERE candidate_id = $1", candidate.CandidateId)

	response := entity.Candidate{}
	if err := row.Scan(response); err != nil {
		return entity.Candidate{}, err
	}

	return response, nil

}
