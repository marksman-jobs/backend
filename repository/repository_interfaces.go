package repository

import "github.com/marksman-jobs/backend/entity"

type CandidateRepository interface {
	Insert(candidate entity.Candidate) error
	FindAll() ([]entity.Candidate, error)
	Get(candidate entity.Candidate) (entity.Candidate, error)
}
