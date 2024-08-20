package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/joshbrgs/aws-mfa/aws"
	"github.com/joshbrgs/aws-mfa/color"
)

type CredentialsType struct {
	SecretAccessKey string `json:"SecretAccessKey"`
	SessionToken    string `json:"SessionToken"`
	Expiration      string `json:"Expiration"`
	AccessKeyId     string `json:"AccessKeyId"`
}

type AwsSession struct {
	Credentials CredentialsType `json:"Credentials"`
}

type MfaDevicesType struct {
	UserName     string `json:"UserName"`
	SerialNumber string `json:"SerialNumber"`
	EnableDate   string `json:"EnableDate"`
}

type AwsMfaDevices struct {
	MfaDevices []MfaDevicesType `json:"MFADevices"`
}

/* Configures the AWS SEssion utilizing the AWS package functions with the arguments of the arn, mfa code, and optional verbose flag*/
func ConfigureSession(arn string, code string, verboseFL bool, defaultProfile string, mfaProfile string) {
	// build aws mfa get session token by mfa command
	result := aws.AWSSessionCommand(arn, code, verboseFL, defaultProfile)

	// convert the result to the json structure for go
	var convertedResult AwsSession

	decodeErr := json.Unmarshal([]byte(result), &convertedResult)

	if decodeErr != nil {
		fmt.Println(color.Red + "JSON decode error!" + color.Reset)
		panic(decodeErr)
	}

	if verboseFL {
		fmt.Printf("Expires at: %s\n", convertedResult.Credentials.Expiration)
	}

	var varname string = "aws_access_key_id"
	value := convertedResult.Credentials.AccessKeyId
	aws.AWSSetCommand(varname, value, verboseFL, mfaProfile)

	varname = "aws_secret_access_key"
	value = convertedResult.Credentials.SecretAccessKey
	aws.AWSSetCommand(varname, value, verboseFL, mfaProfile)

	varname = "aws_session_token"
	value = convertedResult.Credentials.SessionToken
	aws.AWSSetCommand(varname, value, verboseFL, mfaProfile)

	varname = "expiration"
	value = convertedResult.Credentials.Expiration
	aws.AWSSetCommand(varname, value, verboseFL, mfaProfile)
}

/* Retrieves the Arn via username or an environment variable and returns it as a string */
func GetARN(defaultProfile string) (arn string) {
	app := "aws"

	// If arn device has not been set
	arn, ok := os.LookupEnv("MFA_ARN")
	name, alright := os.LookupEnv("AWS_USER")

	if !ok {
		if !alright {
			fmt.Println("What is your AWS username?:")
			fmt.Scanln(&name)
			fmt.Printf("For an easier Auth use this command on windows cmd: set AWS_USER=%s\n", name)
		}

		out, err := exec.Command(app, "iam", "list-mfa-devices", "--user-name", name, "--profile", defaultProfile, "--output", "json").CombinedOutput()
		if err != nil {
			fmt.Println(color.Red + string(out) + color.Reset)
			fmt.Println("Do you have MFA set up in AWS?")
			panic(err)
		}

		// convert the result to the json structure for go
		var result string = string(out)
		var convertedMFA AwsMfaDevices

		decodeErr := json.Unmarshal([]byte(result), &convertedMFA)

		if decodeErr != nil {
			fmt.Println(color.Red + "JSON decode error!" + color.Reset)
			panic(decodeErr)
		}

		arn = convertedMFA.MfaDevices[0].SerialNumber

		if err != nil {
			fmt.Println(color.Purple + "What is your mfa devices arn?:" + color.Reset)
			fmt.Println()

			fmt.Scanln(&arn)

			fmt.Printf("For an easier Auth use this command on windows cmd: set MFA_ARN=%s", arn)
		}
	}

	return arn
}
