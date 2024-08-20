package aws

import (
	"testing"
)

func TestAWSSetCommand(t *testing.T) {
	type args struct {
		varname    string
		value      string
		verboseFl  bool
		mfaProfile string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AWSSetCommand(tt.args.varname, tt.args.value, tt.args.verboseFl, tt.args.mfaProfile)
		})
	}
}

func TestAWSSessionCommand(t *testing.T) {
	type args struct {
		arn            string
		code           string
		verboseFl      bool
		defaultProfile string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := AWSSessionCommand(tt.args.arn, tt.args.code, tt.args.verboseFl, tt.args.defaultProfile); gotResult != tt.wantResult {
				t.Errorf("AWSSessionCommand() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestCreateAwsMfaProfile(t *testing.T) {
	type args struct {
		mfaProfile string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateAwsMfaProfile(tt.args.mfaProfile)
		})
	}
}

func TestKubeConfig(t *testing.T) {
	type args struct {
		verboseFl  bool
		mfaProfile string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			KubeConfig(tt.args.verboseFl, tt.args.mfaProfile)
		})
	}
}

func TestCheckSessionExpiration(t *testing.T) {
	type args struct {
		mfaProfile string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSessionExpiration(tt.args.mfaProfile); got != tt.want {
				t.Errorf("CheckSessionExpiration() = %v, want %v", got, tt.want)
			}
		})
	}
}
