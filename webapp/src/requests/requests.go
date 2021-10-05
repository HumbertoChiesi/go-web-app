package requests

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

//MakeRequestWithAuthentication is used to put the token in the request
func MakeRequestWithAuthentication(r *http.Request, method string, url string, data io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	cookie, _ := cookies.Read(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
