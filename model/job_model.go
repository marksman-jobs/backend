package model

type CreateJobRequest struct {
	JobTitle          string   `bun:"job_title"`
	JobType           string   `bun:"job_type,type:enum"`
	JobFunction       string   `bun:"job_function,type:enum"`
	Currency          string   `bun:"currency,type:enum"`
	Value             int      `bun:"value"`
	SkillsId          []string `bun:"skills_id,fk,array"`
	JobApplicantIds   []string `bun:"job_applicant_ids,fk,array"`
	JobApplicantCount int      `bun:"job_applicant_count"`
	CompanyId         string   `bun:"company_id,fk"`
	RecruiterId       string   `bun:"recruiter_id,fk"`
	LocationId        string   `bun:"location_id,fk"`
}

type CreateJobResponse struct {
	JobId string
}

type GetJobRequest struct {
	JobId string
}

type GetJobResponse struct {
	JobId             string   `bun:"job_id,pk"`
	JobTitle          string   `bun:"job_title"`
	JobType           string   `bun:"job_type,type:enum"`
	JobFunction       string   `bun:"job_function,type:enum"`
	Currency          string   `bun:"currency,type:enum"`
	Value             int      `bun:"value"`
	SkillsId          []string `bun:"skills_id,fk,array"`
	JobApplicantIds   []string `bun:"job_applicant_ids,fk,array"`
	JobApplicantCount int      `bun:"job_applicant_count"`
	CompanyId         string   `bun:"company_id,fk"`
	RecruiterId       string   `bun:"recruiter_id,fk"`
	LocationId        string   `bun:"location_id,fk"`
}
