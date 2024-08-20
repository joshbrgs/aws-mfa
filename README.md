# AWS MFA
## v1.0.1

### What this tool achieves
> Goal of this program in Golang is to make the life of the developer easier, by providing a quick command that sets the user's
> session token and temporary credentials to an aws cli profile [mfa] for AWS CLI Multifactor Authentication.
> This profile keeps track of the user's session determining if it is valid or not.
> The user may also opt in to having the command update their local Kube Config for a quicker MFA login and setup

### Prerequsites
The user will need to download the latest Golang version at [Go Dev](https://go.dev/doc/install).
By default, the installer will install Go to Program Files or Program Files (x86).
I suggest changing this during the install process to your User directory.
To create the executable, cd into this directory and run `go build .`
This will create an executable with the name of the directory.
You can then place it in your `C:\Users\<username>\go\bin` directory to place it in you path and use it like any command *ex) `aws-mfa --kubeConfig`*

### Usage
Base command: aws-mfa
Supported flags:
-kubeConfig ==> _This will prompt you for region code and project name to set up your kubectl config context_
-verbose ==> _This will turn on all error message and command preview_
-mfaProfile ==> ex) aws-mfa -mfaProfile=newProfile _The default is "mfa" but this flag will allow you to choose the name of the profile where the session creds are stored_
-defaultProfile ==> ex) aws-mfa -defaultProfile=profile _The default is "default" but this flag will allow you to choose the name of the profile where your default aws credentials are, useful if you have multiple accounts and projects_
-help or -h ==> _This will show you available flag usage_

### Behind the Curtain
This program is executing aws cli commands and parsing the json response to retrieve your mfa device arn,
utilize the mfa code, and set the session token and keys to a second aws cli profile called [mfa].
If the user does not have the profile it will create it.
If the user's token is still valid, it will let the user know.
If the kubeConfig flag is included it will then execute the aws cli command to configure the kubernetes context with the sessions credentials.

#### Dependencies
- AWS CLI:^2.7

#### Dev Dependancies
- Golang: ^1.17

_MD Author: Josh Burgess_
