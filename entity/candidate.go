package entity

import "github.com/uptrace/bun"

type Candidate struct {
	bun.BaseModel `bun:"table:candidates"`

	CandidateId string `bun:"candidate_id,pk"`
}
