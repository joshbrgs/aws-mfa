package utils

import (
	"testing"
)

func TestConfigureSession(t *testing.T) {
	type args struct {
		arn            string
		code           string
		verboseFL      bool
		defaultProfile string
		mfaProfile     string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConfigureSession(tt.args.arn, tt.args.code, tt.args.verboseFL, tt.args.defaultProfile, tt.args.mfaProfile)
		})
	}
}

func TestGetARN(t *testing.T) {
	type args struct {
		defaultProfile string
	}
	tests := []struct {
		name    string
		args    args
		wantArn string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArn := GetARN(tt.args.defaultProfile); gotArn != tt.wantArn {
				t.Errorf("GetARN() = %v, want %v", gotArn, tt.wantArn)
			}
		})
	}
}
