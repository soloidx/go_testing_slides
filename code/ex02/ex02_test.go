package ex02_test // HL

import "testing"

import "../ex02" // HL

// START OMIT
func TestPassMethodReturnsFail(t *testing.T) { // HL
	ret := ex02.PassMethod(1)
	expected := "A simple string" // Will fail

	if ret != expected {
		t.Error( // HL
			"PassMethod: expected:", expected, // HL
			"returned:", ret) // HL
	}
}

// END OMIT
