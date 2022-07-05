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
	testCase.actualResult, _ = RunBf(testCase.value, nil)

	if testCase.actualResult != testCase.expectedResult {
		t.Fail()
	}
}

func TestCloseLoopMissingRunBf(t *testing.T) {
	testCase := TestCase{
		value:       "+++++[-",
		expectedErr: errors.New("Loop brackets dont match! Could not found matching ']'"),
	}
	_, testCase.actualErr = RunBf(testCase.value, nil)
	if testCase.actualErr.Error() != testCase.expectedErr.Error() {
		t.Fail()
	}
}

func TestStartLoopMissingRunBf(t *testing.T) {
	testCase := TestCase{
		value:       "+++++-]",
		expectedErr: errors.New("Loop brackets dont match! Could not found matching '['"),
	}
	_, testCase.actualErr = RunBf(testCase.value, nil)
	if testCase.actualErr.Error() != testCase.expectedErr.Error() {
		t.Fail()
	}
}

func fib(n int) int {
	if n == 1 {
		return 1
	} else if n == 0 {
		return 0
	}
	return fib(n-1) + fib(n-2)
}

func TestCustomCommandRunBf(t *testing.T) {
	testCase := TestCase{
		value:          "+++++++++++f--.",
		expectedResult: "W",
		expectedErr:    nil,
	}
	//Creates a new command List
	customCommands := newCommandList()
	cmd := Command{
		name: 'f',
		calc: fib,
	}
	//We are adding new commands
	customCommands.Add(cmd)

	testCase.actualResult, _ = RunBf(testCase.value, customCommands)
	if testCase.actualResult != testCase.expectedResult {
		t.Fail()
	}
}
