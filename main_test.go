package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_retrieveMFA(t *testing.T) {
	tests := []struct {
		name     string
		wantCode string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCode := retrieveMFA(); gotCode != tt.wantCode {
				t.Errorf("retrieveMFA() = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}

func Test_updateKube(t *testing.T) {
	type args struct {
		updateFL   bool
		verboseFL  bool
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
			updateKube(tt.args.updateFL, tt.args.verboseFL, tt.args.mfaProfile)
		})
	}
}
