package main

import (
	"reflect"
	"testing"
)

func TestNoDuplicateWords(t *testing.T) {
	testCases := map[string]bool{
		"aa bb cc dd ee":  true,
		"aa bb cc dd aa":  false,
		"aa bb cc dd aaa": true,
	}

	for testCase, truthValue := range testCases {
		result := NoDuplicateWords(testCase)
		if truthValue != result {
			t.Error("invalid result")
			t.Logf("test case: \"%v\"", testCase)
			t.Logf("expected=%v", truthValue)
			t.Logf("result=%v", result)
		}
	}
}

func TestHistogram(t *testing.T) {
	testCases := map[string]map[rune]int{
		"ab": map[rune]int{
			'a': 1,
			'b': 1,
		},
		"aab": map[rune]int{
			'a': 2,
			'b': 1,
		},
		"zzz": map[rune]int{
			'z': 3,
		},
		"zaz": map[rune]int{
			'z': 2,
			'a': 1,
		},
	}

	for testCase, truthValue := range testCases {
		result := Histogram(testCase)
		if !reflect.DeepEqual(result, truthValue) {
			t.Error("invalid result")
			t.Logf("test case: \"%v\"", testCase)
			t.Logf("expected=%v", truthValue)
			t.Logf("result=%v", result)
		}
	}
}

func TestNoAnagrams(t *testing.T) {
	testCases := map[string]bool{
		"abcde fghij":              true,
		"abcde xyz ecdab":          false,
		"a ab abc abd abf abj":     true,
		"iiii oiii ooii oooi oooo": true,
		"oiii ioii iioi iiio":      false,
	}

	for testCase, truthValue := range testCases {
		result := NoAnagrams(testCase)
		if truthValue != result {
			t.Error("invalid result")
			t.Logf("test case: \"%v\"", testCase)
			t.Logf("expected=%v", truthValue)
			t.Logf("result=%v", result)
		}
	}

}
