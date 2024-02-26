package profile

import "app_blade/buildinfo"

type Profile string

const (
	DEV   Profile = "dev"
	TEST  Profile = "test"
	UAT   Profile = "uat"
	STAGE Profile = "stage"
	PROD  Profile = "prod"
)

var (
	// Env is the current environment
	Env Profile
)

func init() {
	Set(buildinfo.Env)
}

func Set(env string) {
	switch env {
	case "dev":
		Env = DEV
	case "test":
		Env = TEST
	case "uat":
		Env = UAT
	case "stage":
		Env = STAGE
	case "prod":
		Env = PROD
	}
}

func Get() string {
	switch Env {
	case DEV:
		return "dev"
	case TEST:
		return "test"
	case UAT:
		return "uat"
	case STAGE:
		return "stage"
	case PROD:
		return "prod"
	}
	return ""
}

func IsDev() bool {
	return Env == DEV
}

func IsTest() bool {
	return Env == TEST
}

func IsUat() bool {
	return Env == UAT
}

func IsStage() bool {
	return Env == STAGE
}

func IsProd() bool {
	return Env == PROD
}
