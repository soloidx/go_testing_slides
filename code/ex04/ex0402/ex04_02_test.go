package ex0402_test

import "testing"
import "net/http"

import "github.com/stretchr/testify/assert"
import "github.com/stretchr/testify/mock"

import "../ex0402"

// START OMIT
type httpClientMock struct {
	mock.Mock
}

func (c *httpClientMock) Get(url string) (*http.Response, error) {
	args := c.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestGetGoogleStatusOK(t *testing.T) {
	//create the expected response
	expResp := &http.Response{}
	expResp.Status = "200 OK"
	clientMock := &httpClientMock{}
	clientMock.On("Get", "http://google.com").Return(expResp, nil)

	resp, err := ex0402.GetGoogleStatus(clientMock)
	assert.Equal(t, "200 OK", resp)
	assert.Nil(t, err)
	clientMock.AssertExpectations(t)
}

// END OMIT
