package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/actions-go/toolkit/core"
)

func runMain() {

	region := os.Getenv("AWS_REGION")
	secrets := os.Getenv("SECRETS")
	branch := os.Getenv("BRANCH")

	if region == "" || secrets == "" {
		core.Error("AWS_REGION or SECRETS is not set")
		os.Exit(1)
	}

	AWS_ACCESS_KEY := ""
	AWS_SECRET_ACCESS_KEY := ""

	var secretsMap map[string]string
	if err := json.Unmarshal([]byte(secrets), &secretsMap); err != nil {
		core.Error(fmt.Sprintf("error reading in secrets map %s", err.Error()))
		return
	}

	if branch == "development" || branch == "qa" || branch == "qa1" || branch == "staging" || branch == "hotfix" || branch == "demo" {
		AWS_ACCESS_KEY = secretsMap["AWS_APTY_NON_PROD_ACCESS_KEY_ID"]
		AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_NON_PROD_SECRET_ACCESS_KEY"]
	} else if region == "us-east-1" || region == "ap-southeast-2" {
		AWS_ACCESS_KEY = secretsMap["AWS_APTY_US_PROD_ACCESS_KEY_ID"]
		AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_US_PROD_SECRET_ACCESS_KEY"]
	} else if region == "eu-central-1" {
		AWS_ACCESS_KEY = secretsMap["AWS_APTY_EU1_PROD_ACCESS_KEY_ID"]
		AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_EU1_PROD_SECRET_ACCESS_KEY"]
	} else {
		// Do nothing
	}

	core.SetOutput("AWS_ACCESS_KEY", AWS_ACCESS_KEY)
	core.SetOutput("AWS_SECRET_ACCESS_KEY", AWS_SECRET_ACCESS_KEY)

}

func main() {
	runMain()
}
