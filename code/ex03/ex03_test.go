package ex03_test

import "testing"

import "github.com/stretchr/testify/assert"
import . "github.com/onsi/gomega"

import "../ex03"

// GO VERSION START OMIT
func TestCheckAgeIsOK(t *testing.T) {
	isOlder, err := ex03.CheckAge(18)

	if isOlder != true {
		t.Error("CheckAge: expected true but returned:", isOlder)
	}
	if err != nil {
		t.Error("CheckAge: expected error as nil")
	}
}

func TestCheckAgeIsErr(t *testing.T) {
	isOlder, err := ex03.CheckAge(-1)

	if isOlder != false {
		t.Error("CheckAge: expected false but returned:", isOlder)
	}

	if err == nil {
		t.Error("CheckAge: expected error")
	}
}

// GO VERSION END OMIT

// TESTIFY START OMIT
func TestTestifyCheckAgeIsOK(t *testing.T) {
	isOlder, err := ex03.CheckAge(18)
	assert.True(t, isOlder)
	assert.Nil(t, err)

	noOlder, err := ex03.CheckAge(17)
	assert.False(t, noOlder)
	assert.Nil(t, err)
}

func TestTestifyCheckAgeIsErr(t *testing.T) {
	isOlder, err := ex03.CheckAge(-1)
	assert.False(t, isOlder)
	assert.Error(t, err)
}

func TestTestifyCheckAgeIsFail(t *testing.T) {
	isOlder, err := ex03.CheckAge(18)
	assert.False(t, isOlder)
	assert.Error(t, err)
}

// TESTIFY  END OMIT

// GOMEGA START OMIT
func TestGomegaCheckAgeIsOK(t *testing.T) {
	RegisterTestingT(t) // HL

	isOlder, err := ex03.CheckAge(18) // HLex
	Ω(isOlder).Should(BeTrue())       // Expect(isOlder).To(BeTrue())
	Ω(err).Should(BeNil())            // Expect(err).To(BeNil())
}

func TestGomegaCheckAgeIsErr(t *testing.T) {
	RegisterTestingT(t)

	isOlder, err := ex03.CheckAge(-1)
	Ω(isOlder).Should(BeFalse())
	Ω(err).ShouldNot(BeNil())
}

func TestGomegaCheckAgeIsFail(t *testing.T) {
	RegisterTestingT(t)

	isOlder, err := ex03.CheckAge(18)
	Ω(isOlder).Should(BeFalse())
	Ω(err).ShouldNot(BeNil())
}

// GOMEGA  END OMIT
