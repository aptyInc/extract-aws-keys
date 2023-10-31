package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/actions-go/toolkit/core"
)

func runMain() {

	fmt.Println("Starting extract aws keys process...")
	
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

	fmt.Println("Region: ", region)
	fmt.Println("secrets: ", secrets)
	fmt.Println("branch: ", branch)

	if branch == "development" || branch == "qa" || branch == "qa1" || branch == "staging" || branch == "hotfix" || branch == "demo" || branch == "automation" {
		fmt.Println("Using AWS_APTY_NON_PROD_ACCESS_KEY_ID")
		AWS_ACCESS_KEY = secretsMap["AWS_APTY_NON_PROD_ACCESS_KEY_ID"]
		AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_NON_PROD_SECRET_ACCESS_KEY"]
	} else if region == "us-east-1" || region == "ap-southeast-2" {
		fmt.Println("Using AWS_APTY_US_PROD_ACCESS_KEY_ID")
		AWS_ACCESS_KEY = secretsMap["AWS_APTY_US_PROD_ACCESS_KEY_ID"]
		AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_US_PROD_SECRET_ACCESS_KEY"]
	} else if region == "eu-central-1" {
		fmt.Println("Using AWS_APTY_EU1_PROD_ACCESS_KEY_ID")
		AWS_ACCESS_KEY = secretsMap["AWS_APTY_EU1_PROD_ACCESS_KEY_ID"]
		AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_EU1_PROD_SECRET_ACCESS_KEY"]
	} else {
		core.Error("No AWS keys used check branch name and region configuration")
		os.Exit(1)
	}

	core.SetOutput("AWS_ACCESS_KEY", AWS_ACCESS_KEY)
	core.SetOutput("AWS_SECRET_ACCESS_KEY", AWS_SECRET_ACCESS_KEY)

	fmt.Println("Done extracting aws keys process...")
}

func main() {
	runMain()
}
