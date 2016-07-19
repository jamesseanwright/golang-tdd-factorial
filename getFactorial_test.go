package main

import (
	"testing"
)

const successMessage = "Expected value matches the actual value!"
const failureFormat = "Expected %d to equal %d"

func TestItShouldComputeTheCorrectFactorial(t *testing.T) {
	if actual, expected := GetFactorial(6), 720; actual != expected {
		t.Errorf(failureFormat, actual, expected)
	} else {
		t.Log(successMessage)
	}
}

func TestItShouldComputeOneForOne(t *testing.T) {
	if actual, expected := GetFactorial(1), 1; actual != expected {
		t.Errorf(failureFormat, actual, expected)
	} else {
		t.Log(successMessage)
	}
}

func TestItShouldComputeOneForZero(t *testing.T) {
	if actual, expected := GetFactorial(0), 1; actual != expected {
		t.Errorf(failureFormat, actual, expected)
	} else {
		t.Log(successMessage)
	}
}

func TestItShouldReturnAZeroValueForNegativeIntegers(t *testing.T) {
	if actual, expected := GetFactorial(-1), 0; actual != expected {
		t.Errorf(failureFormat, actual, expected)
	} else {
		t.Log(successMessage)
	}
}