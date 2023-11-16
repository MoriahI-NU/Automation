package main

import (
	"reflect"
	"testing"
)

// testing ExerimentPython
func TestExperimentPython(t *testing.T) {
	tests := []struct {
		name string
		set  string
		want Results
	}{
		{
			//Only testing Anscombe set One
			name: "Test Case Python",
			set:  "One",
			//rounding wants because time may vary with external factors - test will fail because of this
			//but for this experiment, it is doable to check the got and want values manually
			want: Results{Coefficients: "[3, 0.5]", Time: "0.0002"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExperimentPython(tt.set); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExperimentPython() = %v, want %v", got, tt.want)
			}
		})
	}
}

// testing ExerimentR
func TestExperimentR(t *testing.T) {
	tests := []struct {
		name string
		set  string
		want Results
	}{
		{
			//Only testing Anscombe set One
			name: "Test Case R",
			set:  "One",
			//rounding due to variability - check the got and want values manually
			want: Results{Coefficients: "[3, 0.5]", Time: "0.0006"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExperimentR(tt.set); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExperimentR() = %v, want %v", got, tt.want)
			}
		})
	}
}

// testing ExerimentGo
func TestExperimentGo(t *testing.T) {
	tests := []struct {
		name string
		set  string
		want Results
	}{
		{
			//only testing Anscombe set One
			name: "Test Case Go",
			set:  "One",
			//rounding due to variability - check the got and want values manually

			//ExperimentGo uses different code than that in LinearRegressionComp repo, so the timing will be different.
			//However, it should still run the fastest out of all 3 languages, which is why I've set Time: 0.0001
			want: Results{Coefficients: "[3, 0.5]", Time: "0.0001"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExperimentGo(tt.set); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExperimentGo() = %v, want %v", got, tt.want)
			}
		})
	}
}
