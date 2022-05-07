package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/marksman-jobs/backend/entity"
)

type jobRepositoryImpl struct {
	pgconn *pgxpool.Pool
}

func NewJobRepository(connPool *pgxpool.Pool) JobRepository {
	return &jobRepositoryImpl{
		pgconn: connPool,
	}
}

func (repository *jobRepositoryImpl) Insert(job entity.Job) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// TODO: Add query to insert job here
	row := repository.pgconn.QueryRow(ctx, "")
	err := row.Scan(job)
	if err != nil {
		return "", err
	}

	return job.JobId, nil

}

func (repository *jobRepositoryImpl) FindAll() ([]entity.Job, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response := []entity.Job{}
	rows, err := repository.pgconn.Query(ctx, "SELECT * FROM jobs")
	if err != nil {
		return response, err
	}

	for rows.Next() {
		res := entity.Job{}
		err := rows.Scan(res)

		if err != nil {
			return []entity.Job{}, err
		}

		response = append(response, res)
	}

	return response, nil

}

func (repository *jobRepositoryImpl) Get(job entity.Job) (entity.Job, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	row := repository.pgconn.QueryRow(ctx, "SELECT * FROM jobs WHERE job_id = $1", job.JobId)

	response := entity.Job{}
	err := row.Scan(response)

	if err != nil {
		return entity.Job{}, err
	}

	return response, nil

}
