package main

import "testing"

func TestIsValid(t *testing.T) {
	testCases := map[string]bool{
		"aa bb cc dd ee":  true,
		"aa bb cc dd aa":  false,
		"aa bb cc dd aaa": true,
	}

	for testCase, truthValue := range testCases {
		result := IsValid(testCase)
		if truthValue != result {
			t.Error("invalid result")
			t.Logf("test case: \"%v\"", testCase)
			t.Logf("expected=%v", truthValue)
			t.Logf("result=%v", result)
		}
	}
}
