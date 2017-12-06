package main

import "testing"

func TestEscapeMaze(t *testing.T) {
	testList := []int{
		0,
		3,
		0,
		1,
		-3,
	}

	goldStepCount := 5

	stepCount := EscapeMaze(testList)

	if goldStepCount != stepCount {
		t.Error("invalid step count")
		t.Logf("expected %v", goldStepCount)
		t.Logf("got %v", stepCount)
	}
}

func TestEscapeMazeP2(t *testing.T) {
	testList := []int{
		0,
		3,
		0,
		1,
		-3,
	}

	goldStepCount := 10

	stepCount := EscapeMazeP2(testList)

	if goldStepCount != stepCount {
		t.Error("invalid step count")
		t.Logf("expected %v", goldStepCount)
		t.Logf("got %v", stepCount)
	}
}
