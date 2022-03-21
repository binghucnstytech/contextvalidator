package util

import (
	"testing"
)

func TestValidCases(t *testing.T) {
	validContexts := []string{
		"([])",
		"([]{}())",
		"(){}()",
	}

	for _, context := range validContexts {
		if !IsContatinersDefinitionValid(context) {
			t.Errorf(context + " is valid, but result got false")
		}
	}
}

func TestInvalidCases(t *testing.T) {
	invalidContexts := []string{
		"(])",
		"([)]",
		"{[}",
	}

	for _, context := range invalidContexts {
		if IsContatinersDefinitionValid(context) {
			t.Errorf(context + " is invalid, but result got true")
		}
	}
}
