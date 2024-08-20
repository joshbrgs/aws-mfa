package aws

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/joshbrgs/aws-mfa/color"
)

func AWSSetCommand(varname string, value string, verboseFl bool, mfaProfile string) {
	var app string = "aws"
	var configureCommand string = "configure"
	var setCommand string = "set"
	var arg1 string = "--profile"
	var arg1Val string = mfaProfile
	var arg2 string = "--output"
	var arg2Val string = "json"
	var exampleCommand string = app + " " + configureCommand + " " + setCommand + " " + varname + " " + value + " " + arg1 + " " + arg1Val + " " + arg2 + " " + arg2Val

	if verboseFl {
		fmt.Println(exampleCommand)
	}

	out, err := exec.Command(app, configureCommand, setCommand, varname, value, arg1, arg1Val, arg2, arg2Val).CombinedOutput()
	if err != nil {
		fmt.Println(color.Red + string(out) + color.Reset)
		panic(err)
	}

	if out != nil && verboseFl {
		fmt.Printf("Set AWS Cred: %s ✔\n", varname)
	}
}

func AWSSessionCommand(arn string, code string, verboseFl bool, defaultProfile string) (result string) {
	var app string = "aws"
	command := "sts"
	command1 := "get-session-token"
	arg1 := "--serial-number"
	arg1Val := arn
	arg2 := "--token-code"
	arg2Val := code
	arg3 := "--profile"
	arg3Val := defaultProfile
	arg4 := "--output"
	arg4Val := "json"
	execExample := app + " " + command + " " + command1 + " " + arg1 + " " + arg1Val + " " + arg2 + " " + arg2Val + " " + arg3 + " " + arg3Val + " " + arg4 + " " + arg4Val

	if verboseFl {
		fmt.Println(execExample)
	}
	out, err := exec.Command(app, command, command1, arg1, arg1Val, arg2, arg2Val, arg3, arg3Val, arg4, arg4Val).CombinedOutput()
	if err != nil {
		fmt.Println(color.Red + string(out) + color.Reset)
		panic(err)
	}

	result = string(out)

	return result
}

func CreateAwsMfaProfile(mfaProfile string) {
	fmt.Printf("Creating %s profile...\n", mfaProfile)
	profile := "profile." + mfaProfile + ".region"
	out, err := exec.Command("aws", "configure", "set", profile, "us-east-1").CombinedOutput()
	if err != nil {
		fmt.Println(color.Red + string(out) + color.Reset)
		panic(err)
	}

	profile = "profile." + mfaProfile + ".output"
	out, err = exec.Command("aws", "configure", "set", profile, "json").CombinedOutput()
	if err != nil {
		fmt.Println(color.Red + string(out) + color.Reset)
		panic(err)
	}

	if out != nil {
		fmt.Printf("%s profile created ✔\n", mfaProfile)
	}
}

func MfaProfileCheck(mfaProfile string) bool {
	out, err := exec.Command("aws", "configure", "list", "--profile", mfaProfile).CombinedOutput()
	if err != nil {
		fmt.Println("We will create a AWS CLI profile for your MFA Creds")
	}

	profileCheck := string(out)

	if strings.Contains(profileCheck, "could not be found") {
		CreateAwsMfaProfile(mfaProfile)
	}

	// get expiration, if expired, return and say sesssion is still active
	var live bool = CheckSessionExpiration(mfaProfile)

	if live {
		return true
	}

	return false
}

func KubeConfig(verboseFl bool, mfaProfile string) {
	var app string = "aws"

	fmt.Println("Region Code:")
	var region string
	fmt.Scanln(&region)

	fmt.Println("What is your Cluster Name:")
	var cluster string
	fmt.Scanln(&cluster)

	if verboseFl {
		fmt.Println(app + " " + "eks" + " " + "update-kubeconfig" + " " + "--region" + " " + region + " " + "--name" + " " + cluster + " " + "--profile" + " " + mfaProfile)
	}
	out, err := exec.Command(app, "eks", "update-kubeconfig", "--region", region, "--name", cluster, "--profile", mfaProfile).CombinedOutput()
	if err != nil {
		fmt.Println(color.Red + string(out) + color.Reset)
		panic(err)
	}

	if out != nil {
		fmt.Println("Kube Config Updated ✔")
	}
}

func CheckSessionExpiration(mfaProfile string) bool {
	profile := "profile." + mfaProfile + ".expiration"
	out, err := exec.Command("aws", "configure", "get", profile).CombinedOutput()
	if err != nil {
		fmt.Println("**Ignore if this is your first time**")
		fmt.Println("Error in reading token expiration date")
	}

	var expiresAt string = string(out)
	expiresAt = strings.ReplaceAll(expiresAt, "\x0d\x0a", "")
	var now time.Time = time.Now()
	t, err := time.Parse(time.RFC3339, expiresAt)
	if err != nil {
		fmt.Println(color.Red + err.Error() + color.Reset)
	}

	if now.Before(t) {
		fmt.Printf("Time is %s\n", now)
		fmt.Println("Session Token is still valid")
		return true
	}

	return false
}
