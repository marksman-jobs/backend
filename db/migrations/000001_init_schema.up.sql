CREATE TYPE "job_function" AS ENUM (
  'Engineering',
  'Software Engineering',
  'Legal',
  'Paralegal',
  'Medical',
  'Paramedical',
  'Management',
  'Finance',
  'Accounting',
  'Product',
  'Human Resources',
  'Other'
);

CREATE TYPE "pronouns" AS ENUM (
  'he/him',
  'she/her',
  'they/them',
  'other'
);

CREATE TYPE "job_type" AS ENUM (
  'Full Time',
  'Part Time',
  'Contract',
  'Internship',
  'Consultant'
);

CREATE TYPE "currency" AS ENUM (
  'USD',
  'GBP',
  'INR',
  'AED',
  'Other'
);

CREATE TYPE "proficiency" AS ENUM (
  'Beginner',
  'Intermediate',
  'Advanced',
  'Expert',
  'Legend'
);

CREATE TYPE "education_level" AS ENUM (
  'High School',
  'Bachelors',
  'Masters',
  'Doctorate'
);

CREATE TYPE "job_status" AS ENUM (
  'Current',
  'Desired',
  'Applied'
);

CREATE TABLE "skill" (
  "id" uuid PRIMARY KEY,
  "title" varchar(64) NOT NULL,
  "proficiency" proficiency NOT NULL
);

CREATE TABLE "education" (
  "id" uuid PRIMARY KEY,
  "school" varchar(96) NOT NULL,
  "level" education_level NOT NULL
);

CREATE TABLE "company" (
  "id" uuid PRIMARY KEY,
  "description" varchar(1024)
);

CREATE TABLE "location" (
  "id" uuid PRIMARY KEY,
  "remote" boolean NOT NULL,
  "city" varchar(64) NOT NULL,
  "state" varchar(64) NOT NULL,
  "country" varchar(64) NOT NULL
);

CREATE TABLE "company_location" (
  "id" uuid PRIMARY KEY,
  "company_id" uuid NOT NULL REFERENCES "company",
  "location_id" uuid NOT NULL REFERENCES "location"
);

CREATE TABLE "user" (
  "id" uuid PRIMARY KEY,
  "first_name" varchar(32) NOT NULL,
  "last_name" varchar(32) NOT NULL,
  "email" varchar(96) NOT NULL,
  "is_email_verified" boolean NOT NULL,
  "password_hash" char(72) NOT NULL,
  "pronouns" pronouns NOT NULL,
  "is_candidate" boolean NOT NULL,
  "is_recruiter" boolean NOT NULL
);

CREATE TABLE "recruiter" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL REFERENCES "user",
  "company_id" uuid NOT NULL REFERENCES "company"
);

CREATE TABLE "candidate" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL REFERENCES "user",
  "resume_id" char(64)
);

CREATE TABLE "job" (
  "id" uuid PRIMARY KEY,
  "title" varchar(96) NOT NULL,
  "type" job_type NOT NULL,
  "function" job_function NOT NULL,
  "currency" currency NOT NULL,
  "value" int NOT NULL,
  "number_of_applicants" int,
  "company_id" uuid REFERENCES "company",
  "recruiter_id" uuid REFERENCES "recruiter",
  "location_id" uuid NOT NULL REFERENCES "location"
);

CREATE TABLE "candidate_location" (
  "id" uuid PRIMARY KEY,
  "candidate_id" uuid NOT NULL REFERENCES "candidate",
  "location_id" uuid NOT NULL REFERENCES "location",
  "current_location" boolean NOT NULL
);

CREATE TABLE "candidate_job" (
  "id" uuid PRIMARY KEY,
  "candidate_id" uuid NOT NULL REFERENCES "candidate",
  "job_id" uuid NOT NULL REFERENCES "job",
  "status" job_status NOT NULL
);

CREATE TABLE "candidate_company" (
  "id" uuid PRIMARY KEY,
  "candidate_id" uuid NOT NULL REFERENCES "candidate",
  "company_id" uuid NOT NULL REFERENCES "company",
  "current_company" boolean NOT NULL,
  "start" timestamp NOT NULL,
  "end" timestamp
);

CREATE TABLE "candidate_skill" (
  "id" uuid PRIMARY KEY,
  "candidate_id" uuid NOT NULL REFERENCES "candidate",
  "skill_id" uuid NOT NULL REFERENCES "skill"
);
