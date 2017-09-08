package ex01_test

import "testing"

import "../ex01"

// START OMIT
func TestPassMethodReturnsOk(t *testing.T) {
	ret := ex01.PassMethod(1)
	expected := "A valid string 1"

	if ret != expected {
		t.Error(
			"PassMethod: expected:", expected,
			"returned:", ret)
	}
}

// END OMIT
