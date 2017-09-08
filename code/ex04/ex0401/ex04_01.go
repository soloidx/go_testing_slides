package ex0401

import "net/http"

// GetGoogleStatusBad get the http status of google.com
// START OMIT
func GetGoogleStatusBad() (string, error) {
	resp, err := http.Get("http://google.com")
	if err != nil {
		return "", err
	}

	status := resp.Status

	return status, nil
}

// END OMIT
