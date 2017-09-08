package ex03

import "errors"

// CheckAge checks if you are old enough
// START OMIT
func CheckAge(age int) (bool, error) {
	if age <= 0 {
		return false, errors.New("the age is invalid")
	}

	if age >= 18 {
		return true, nil
	}

	return false, nil
}

// END OMIT
