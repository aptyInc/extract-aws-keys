package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/actions-go/toolkit/core"
)

func runMain() {

	fmt.Println("Starting extract aws keys process v1.1...")

	region := os.Getenv("AWS_REGION")
	secrets := os.Getenv("SECRETS")
	environment := os.Getenv("ENVIRONMENT")

	if region == "" || secrets == "" {
		core.Error("AWS_REGION or SECRETS is not set")
		os.Exit(1)
	}

	AWS_ACCESS_KEY := ""
	AWS_SECRET_ACCESS_KEY := ""
	DB_PASSWORD := ""
	NO_REPLY_EMAIL_PASSWORD := ""
	INTERNAL_API_ACCESS_KEY := ""
	FEATURE_FLAG_API_KEY := ""
	EXTERNAL_API_ACCESS_KEY := ""

	var secretsMap map[string]string
	if err := json.Unmarshal([]byte(secrets), &secretsMap); err != nil {
		core.Error(fmt.Sprintf("error reading in secrets map %s", err.Error()))
		return
	}

	fmt.Println("Region: ", region)
	fmt.Println("secrets: ", secrets)
	fmt.Println("environment: ", environment)
	fmt.Println("external-key: ", secretsMap["K8S_NON_PROD_EXTERNAL_API_ACCESS_KEY"])

	if environment == "development" || environment == "qa" || environment == "staging" || environment == "hotfix" || environment == "automation" || environment == "onprem" {
		fmt.Println("Using Non Prod Keys")
		AWS_ACCESS_KEY = secretsMap["AWS_APTY_NON_PROD_ACCESS_KEY_ID"]
		AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_NON_PROD_SECRET_ACCESS_KEY"]
		DB_PASSWORD = secretsMap["K8S_NON_PROD_DB_PASSWORD"]
		INTERNAL_API_ACCESS_KEY = secretsMap["K8S_NON_PROD_INTERNAL_API_ACCESS_KEY"]
		EXTERNAL_API_ACCESS_KEY = secretsMap["K8S_NON_PROD_EXTERNAL_API_ACCESS_KEY"]
		FEATURE_FLAG_API_KEY = secretsMap["K8S_QA_FEATURE_FLAG_API_KEY"] // QA env used for QA, hotfix, automation envs
		if environment == "development" {
			FEATURE_FLAG_API_KEY = secretsMap["K8S_DEVELOPMENT_FEATURE_FLAG_API_KEY"]
		} else if environment == "staging" {
			FEATURE_FLAG_API_KEY = secretsMap["K8S_STAGING_FEATURE_FLAG_API_KEY"]
		}
	} else if environment == "demo" {
		fmt.Println("Using Demo Keys")
		AWS_ACCESS_KEY = secretsMap["AWS_APTY_NON_PROD_ACCESS_KEY_ID"]
		AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_NON_PROD_SECRET_ACCESS_KEY"]
		DB_PASSWORD = secretsMap["K8S_DEMO_DB_PASSWORD"]
		INTERNAL_API_ACCESS_KEY = secretsMap["K8S_NON_PROD_INTERNAL_API_ACCESS_KEY"]
		EXTERNAL_API_ACCESS_KEY = secretsMap["K8S_NON_PROD_EXTERNAL_API_ACCESS_KEY"]
		FEATURE_FLAG_API_KEY = secretsMap["K8S_DEMO_FEATURE_FLAG_API_KEY"]
	} else if region == "us-east-1" {
		fmt.Println("Using US prod keys")
		AWS_ACCESS_KEY = secretsMap["AWS_APTY_US_PROD_ACCESS_KEY_ID"]
		AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_US_PROD_SECRET_ACCESS_KEY"]
		DB_PASSWORD = secretsMap["K8S_US_PROD_DB_PASSWORD"]
		INTERNAL_API_ACCESS_KEY = secretsMap["K8S_PROD_INTERNAL_API_ACCESS_KEY"]
		EXTERNAL_API_ACCESS_KEY = secretsMap["K8S_PROD_EXTERNAL_API_ACCESS_KEY"]
		FEATURE_FLAG_API_KEY = secretsMap["K8S_US_PROD_FEATURE_FLAG_API_KEY"]
	} else if region == "ap-southeast-2" {
		fmt.Println("Using AU prod Keys")
		AWS_ACCESS_KEY = secretsMap["AWS_APTY_US_PROD_ACCESS_KEY_ID"]
		AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_US_PROD_SECRET_ACCESS_KEY"]
		DB_PASSWORD = secretsMap["K8S_AU_PROD_DB_PASSWORD"]
		INTERNAL_API_ACCESS_KEY = secretsMap["K8S_PROD_INTERNAL_API_ACCESS_KEY"]
		EXTERNAL_API_ACCESS_KEY = secretsMap["K8S_PROD_EXTERNAL_API_ACCESS_KEY"]
		FEATURE_FLAG_API_KEY = secretsMap["K8S_US_PROD_FEATURE_FLAG_API_KEY"]
	} else if region == "eu-central-1" {
		fmt.Println("Using EU1 prod keys")
		AWS_ACCESS_KEY = secretsMap["AWS_APTY_EU1_PROD_ACCESS_KEY_ID"]
		AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_EU1_PROD_SECRET_ACCESS_KEY"]
		DB_PASSWORD = secretsMap["K8S_EU1_PROD_DB_PASSWORD"]
		INTERNAL_API_ACCESS_KEY = secretsMap["K8S_PROD_INTERNAL_API_ACCESS_KEY"]
		EXTERNAL_API_ACCESS_KEY = secretsMap["K8S_PROD_EXTERNAL_API_ACCESS_KEY"]
		FEATURE_FLAG_API_KEY = secretsMap["K8S_EU1_PROD_FEATURE_FLAG_API_KEY"]
	} else {
		core.Error("No AWS keys used check environment name and region configuration")
		os.Exit(1)
	}
	NO_REPLY_EMAIL_PASSWORD = secretsMap["K8S_NO_REPLY_EMAIL_PASSWORD"]

	core.SetOutput("AWS_ACCESS_KEY", AWS_ACCESS_KEY)
	core.SetOutput("AWS_SECRET_ACCESS_KEY", AWS_SECRET_ACCESS_KEY)
	core.SetOutput("DB_PASSWORD", DB_PASSWORD)
	core.SetOutput("INTERNAL_API_ACCESS_KEY", INTERNAL_API_ACCESS_KEY)
	core.SetOutput("EXTERNAL_API_ACCESS_KEY", EXTERNAL_API_ACCESS_KEY)
	core.SetOutput("NO_REPLY_EMAIL_PASSWORD", NO_REPLY_EMAIL_PASSWORD)
	core.SetOutput("FEATURE_FLAG_API_KEY", FEATURE_FLAG_API_KEY)

	fmt.Println("Done extracting aws keys process...")

	// Save secretsMap to a JSON file
	file, err := os.Create("secrets.json")
	if err != nil {
		core.Error(fmt.Sprintf("Error creating file: %s", err.Error()))
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print JSON
	if err := encoder.Encode(secretsMap); err != nil {
		core.Error(fmt.Sprintf("Error writing JSON to file: %s", err.Error()))
		os.Exit(1)
	}

	fmt.Println("Secrets have been saved to secrets.json")
}

func main() {
	runMain()
}
