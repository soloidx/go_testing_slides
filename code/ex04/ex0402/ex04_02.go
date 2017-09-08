package ex0402

import "net/http"

//import "errors"
import "fmt"

// START OMIT

// HTTPClientInterface interface for http Client
type HTTPClientInterface interface {
	Get(url string) (resp *http.Response, err error)
}

// GetGoogleStatus get the http status of google.com
func GetGoogleStatus(httpClient HTTPClientInterface) (string, error) {
	resp, err := httpClient.Get("http://google.com")
	if err != nil {
		return "", err
	}

	status := resp.Status

	return status, nil
}

func main() {
	client := &http.Client{}
	resp, _ := GetGoogleStatus(client)
	fmt.Print(resp)
}

// END OMIT
