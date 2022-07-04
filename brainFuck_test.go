package main

import (
	"errors"
	"testing"
)

type TestCase struct {
	value          string
	expectedResult string
	expectedErr    error
	actualResult   string
	actualErr      error
}

func TestHelloWorldRunBf(t *testing.T) {
	testCase := TestCase{
		value:          ">++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<++.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-]<+.",
		expectedResult: "Hello, World!",
	}
	_, testCase.actualResult = RunBf(testCase.value)

	if testCase.actualResult != testCase.expectedResult {
		t.Fail()
	}
}

func TestCloseLoopMissingRunBf(t *testing.T) {
	testCase := TestCase{
		value:       "+++++[-",
		expectedErr: errors.New("Loop brackets dont match! Could not found matching ']'"),
	}
	testCase.actualErr, _ = RunBf(testCase.value)
	if testCase.actualErr.Error() != testCase.expectedErr.Error() {
		t.Fail()
	}
}

func TestStartLoopMissingRunBf(t *testing.T) {
	testCase := TestCase{
		value:       "+++++-]",
		expectedErr: errors.New("Loop brackets dont match! Could not found matching '['"),
	}
	testCase.actualErr, _ = RunBf(testCase.value)
	if testCase.actualErr.Error() != testCase.expectedErr.Error() {
		t.Fail()
	}
}
