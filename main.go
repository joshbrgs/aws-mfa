package main

import (
	"flag"
	"fmt"

	"github.com/joshbrgs/aws-mfa/aws"
	"github.com/joshbrgs/aws-mfa/color"
	"github.com/joshbrgs/aws-mfa/utils"
)

// Option for flag arguments : --verbose, --kubeConfig, --help
func main() {
	// Flags
	verboseFL := flag.Bool("verbose", false, "a bool that turns on output")
	updateFL := flag.Bool("kubeConfig", false, "a bool to activate updating your local kube config with your session keys")
	defaultProfileFL := flag.String("defaultProfile", "default", "declare the specific profile that holds your keys for your aws account, default value is --profile default")
	mfaProfileFL := flag.String("mfaProfile", "mfa", "declare the specific profile you want your mfa credentials and keys saved to")
	flag.Parse()

	// Get the user's arn
	var arn string = utils.GetARN(*defaultProfileFL)
	// check if they have mfa profile, if not create it
	var live bool = aws.MfaProfileCheck(*mfaProfileFL)
	// If the session token is still live skip MFA
	if live {
		updateKube(*updateFL, *verboseFL, *mfaProfileFL)
		return
	}

	var code string = retrieveMFA()
	// set results of the mfa as cred variables under mfa profile
	utils.ConfigureSession(arn, code, *verboseFL, *defaultProfileFL, *mfaProfileFL)
	fmt.Println(color.Cyan + "Your Session token is set ✔" + color.Reset)
	updateKube(*updateFL, *verboseFL, *mfaProfileFL)
}

/***************************** Helper Functions *********************************/
func retrieveMFA() (code string) {
	// ask for the mfa code if one is not in args
	fmt.Println(color.Purple + "What is your mfa code?:" + color.Reset)
	fmt.Println()

	fmt.Scanln(&code)

	return code
}

func updateKube(updateFL bool, verboseFL bool, mfaProfile string) {
	// Update kubernetes config if flagged
	if updateFL {
		aws.KubeConfig(verboseFL, mfaProfile)
	}
}
