package model

type CreateCandidateRequest struct {
	FirstName          string
	LastName           string
	CurrentCompanyId   string
	CurrentRoleId      string
	CurrentLocationId  string
	DesiredLocationIds []string
	DesiredRoleId      string
	Email              string
	PasswordHash       string
	Pronouns           int
	EducationId        string
	SkillsId           string
	ResumeId           string
}

type CreateCandidateResponse struct {
	CandidateId string
}

type GetCandidateRequest struct {
	CandidateId string
}

type GetCandidateResponse struct {
	FirstName          string
	LastName           string
	CurrentCompanyId   string
	CurrentRoleId      string
	CurrentLocationId  string
	DesiredLocationIds []string
	DesiredRoleId      string
	Email              string
	PasswordHash       string
	IsEmailVerified    bool
	Pronouns           int
	EducationId        string
	SkillsId           string
	ResumeId           string
}
