const core = require('@actions/core');
const process = require('process');

function runMain() {
    console.log("Starting extract AWS keys process v1.1...");

    const region = process.env.AWS_REGION;
    const secrets = process.env.SECRETS;
    const environment = process.env.ENVIRONMENT;

    if (!region || !secrets) {
        core.setFailed("AWS_REGION or SECRETS is not set");
        return;
    }

    let AWS_ACCESS_KEY = "";
    let AWS_SECRET_ACCESS_KEY = "";
    let DB_PASSWORD = "";
    let NO_REPLY_EMAIL_PASSWORD = "";
    let INTERNAL_API_ACCESS_KEY = "";
    let FEATURE_FLAG_API_KEY = "";
    let EXTERNAL_API_ACCESS_KEY = "";

    let secretsMap;
    try {
        secretsMap = JSON.parse(secrets);
    } catch (error) {
        core.setFailed(`Error reading in secrets map: ${error.message}`);
        return;
    }

    console.log("Region: ", region);
    console.log("Secrets: ", secrets);
    console.log("Environment: ", environment);
    console.log("External Key: ", secretsMap["K8S_NON_PROD_EXTERNAL_API_ACCESS_KEY"]);

    if (["development", "qa", "staging", "hotfix", "automation", "onprem"].includes(environment)) {
        console.log("Using Non Prod Keys");
        AWS_ACCESS_KEY = secretsMap["AWS_APTY_NON_PROD_ACCESS_KEY_ID"];
        AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_NON_PROD_SECRET_ACCESS_KEY"];
        DB_PASSWORD = secretsMap["K8S_NON_PROD_DB_PASSWORD"];
        INTERNAL_API_ACCESS_KEY = secretsMap["K8S_NON_PROD_INTERNAL_API_ACCESS_KEY"];
        EXTERNAL_API_ACCESS_KEY = secretsMap["K8S_NON_PROD_EXTERNAL_API_ACCESS_KEY"];
        FEATURE_FLAG_API_KEY = secretsMap["K8S_QA_FEATURE_FLAG_API_KEY"];

        if (environment === "development") {
            FEATURE_FLAG_API_KEY = secretsMap["K8S_DEVELOPMENT_FEATURE_FLAG_API_KEY"];
        } else if (environment === "staging") {
            FEATURE_FLAG_API_KEY = secretsMap["K8S_STAGING_FEATURE_FLAG_API_KEY"];
        }
    } else if (environment === "demo") {
        console.log("Using Demo Keys");
        AWS_ACCESS_KEY = secretsMap["AWS_APTY_NON_PROD_ACCESS_KEY_ID"];
        AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_NON_PROD_SECRET_ACCESS_KEY"];
        DB_PASSWORD = secretsMap["K8S_DEMO_DB_PASSWORD"];
        INTERNAL_API_ACCESS_KEY = secretsMap["K8S_NON_PROD_INTERNAL_API_ACCESS_KEY"];
        EXTERNAL_API_ACCESS_KEY = secretsMap["K8S_NON_PROD_EXTERNAL_API_ACCESS_KEY"];
        FEATURE_FLAG_API_KEY = secretsMap["K8S_DEMO_FEATURE_FLAG_API_KEY"];
    } else if (region === "us-east-1") {
        console.log("Using US prod keys");
        AWS_ACCESS_KEY = secretsMap["AWS_APTY_US_PROD_ACCESS_KEY_ID"];
        AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_US_PROD_SECRET_ACCESS_KEY"];
        DB_PASSWORD = secretsMap["K8S_US_PROD_DB_PASSWORD"];
        INTERNAL_API_ACCESS_KEY = secretsMap["K8S_PROD_INTERNAL_API_ACCESS_KEY"];
        EXTERNAL_API_ACCESS_KEY = secretsMap["K8S_PROD_EXTERNAL_API_ACCESS_KEY"];
        FEATURE_FLAG_API_KEY = secretsMap["K8S_US_PROD_FEATURE_FLAG_API_KEY"];
    } else if (region === "ap-southeast-2") {
        console.log("Using AU prod Keys");
        AWS_ACCESS_KEY = secretsMap["AWS_APTY_US_PROD_ACCESS_KEY_ID"];
        AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_US_PROD_SECRET_ACCESS_KEY"];
        DB_PASSWORD = secretsMap["K8S_AU_PROD_DB_PASSWORD"];
        INTERNAL_API_ACCESS_KEY = secretsMap["K8S_PROD_INTERNAL_API_ACCESS_KEY"];
        EXTERNAL_API_ACCESS_KEY = secretsMap["K8S_PROD_EXTERNAL_API_ACCESS_KEY"];
        FEATURE_FLAG_API_KEY = secretsMap["K8S_US_PROD_FEATURE_FLAG_API_KEY"];
    } else if (region === "eu-central-1") {
        console.log("Using EU1 prod keys");
        AWS_ACCESS_KEY = secretsMap["AWS_APTY_EU1_PROD_ACCESS_KEY_ID"];
        AWS_SECRET_ACCESS_KEY = secretsMap["AWS_APTY_EU1_PROD_SECRET_ACCESS_KEY"];
        DB_PASSWORD = secretsMap["K8S_EU1_PROD_DB_PASSWORD"];
        INTERNAL_API_ACCESS_KEY = secretsMap["K8S_PROD_INTERNAL_API_ACCESS_KEY"];
        EXTERNAL_API_ACCESS_KEY = secretsMap["K8S_PROD_EXTERNAL_API_ACCESS_KEY"];
        FEATURE_FLAG_API_KEY = secretsMap["K8S_EU1_PROD_FEATURE_FLAG_API_KEY"];
    } else {
        core.setFailed("No AWS keys used, check environment name and region configuration");
        return;
    }

    NO_REPLY_EMAIL_PASSWORD = secretsMap["K8S_NO_REPLY_EMAIL_PASSWORD"];

    core.setOutput("AWS_ACCESS_KEY", AWS_ACCESS_KEY);
    core.setOutput("AWS_SECRET_ACCESS_KEY", AWS_SECRET_ACCESS_KEY);
    core.setOutput("DB_PASSWORD", DB_PASSWORD);
    core.setOutput("INTERNAL_API_ACCESS_KEY", INTERNAL_API_ACCESS_KEY);
    core.setOutput("EXTERNAL_API_ACCESS_KEY", EXTERNAL_API_ACCESS_KEY);
    core.setOutput("NO_REPLY_EMAIL_PASSWORD", NO_REPLY_EMAIL_PASSWORD);
    core.setOutput("FEATURE_FLAG_API_KEY", FEATURE_FLAG_API_KEY);

    console.log("Done extracting AWS keys process...");
}

runMain();
