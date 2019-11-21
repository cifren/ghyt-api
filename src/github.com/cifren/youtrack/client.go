package youtrack

import (
	"io"
	"net/url"
	"net/http"
)

type Request struct {
	QueryParams *url.Values
	Endpoint    string
	Method      string
	Headers     map[string]string
	Body        io.Reader
}

type ClientInterface interface {
	Get(request *Request) (*http.Response, error)
	Post(request *Request) (*http.Response, error)
	Request(request *Request) (*http.Response, error)
}

type Client struct {
	Url string
	Token string
	http *http.Client
}
func (this *Client) Get(request *Request) (*http.Response, error){
	request.Method = "get"
	return this.Request(request)
}
func (this *Client) Post(request *Request) (*http.Response, error){
	request.Method = "post"
	return this.Request(request)
}
func (this *Client) Request(request *Request) (*http.Response, error){
	req, err := http.NewRequest(
		request.Method,
		this.Url + request.Endpoint + request.QueryParams.Encode(),
		request.Body,
	)

	if err != nil {
		return nil, err
	}

	// Add auth header
	req.Header.Add("Authorization", this.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	for k, v := range request.Headers {
		req.Header.Add(k, v)
	}

	res, err := this.http.Do(req)

	if err != nil {
		return nil, err
	}

	return res, err
}

