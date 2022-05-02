package entity

type Candidate struct {
	CandidateId        string   `bun:"candidate_id,pk,type:uuid"`
	FirstName          string   `bun:"first_name"`
	LastName           string   `bun:"last_name"`
	CurrentCompanyId   string   `bun:"current_company_id,fk"`
	CurrentRoleId      string   `bun:"current_role_id,fk"`
	CurrentLocationId  string   `bun:"current_location_id,fk"`
	DesiredLocationIds []string `bun:"desired_location_ids,fk,array"`
	DesiredRoleId      string   `bun:"desired_role_id,fk"`
	Email              string   `bun:"email"`
	PasswordHash       string   `bun:"password_hash"`
	IsEmailVerified    bool     `bun:"is_email_verified"`
	Pronouns           int      `bun:"pronouns,type:enum"`
	EducationId        string   `bun:"education_id,fk"`
	SkillsId           string   `bun:"skills_id,fk,array"`
	ResumeId           string   `bun:"resume_id,fk"`
}
