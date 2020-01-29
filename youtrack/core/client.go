package core

import (
	"io"
	"net/url"
	"net/http"
	"bytes"
	"fmt"
)

type Request struct {
	QueryParams url.Values
	Endpoint    string
	Method      string
	Headers     map[string]string
	Body        io.Reader
}
func NewRequest(endpoint string) *Request {
    request := Request{
		Endpoint: endpoint,
		QueryParams: make(url.Values),
	}

	return &request
}

type ClientInterface interface {
	Get(request Request) (http.Response, error)
	Post(request Request) (http.Response, error)
	Request(request Request) (http.Response, error)
}

type Client struct {
	Url string
	Token string
	http http.Client
}
func (this Client) Get(request Request) (http.Response, error){
	request.Method = "GET"
	return this.Request(request)
}
func (this Client) Post(request Request) (http.Response, error){
	request.Method = "POST"
	return this.Request(request)
}
func (this Client) Request(request Request) (http.Response, error){
	url, err := url.Parse(this.Url + "/" + request.Endpoint)
	url.RawQuery = request.QueryParams.Encode()

	var body io.Reader
	if request.Body != nil {
		body = request.Body.(io.Reader)
	}
	req, err := http.NewRequest(
		request.Method,
		url.String(),
		body,
	)

	if err != nil {
		panic(err)
		//return http.Response{}, err
	}

	// Add auth header
	req.Header.Add("Authorization", this.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	for k, v := range request.Headers {
		req.Header.Add(k, v)
	}
	if body != nil {
		fmt.Println(fmt.Sprintf(
			"Request '%s' : %s, headers => %+v, body => %s",
			request.Method,
			url.String(),
			req.Header,
			body.(*bytes.Buffer).String(),
		))
	} else {
		fmt.Println(fmt.Sprintf(
			"Request '%s' : %s, headers => %+v",
			request.Method,
			url.String(),
			req.Header,
		))
	}

	res, err2 := this.http.Do(req)

	if err2 != nil {
		panic(err2)
		//return http.Response{}, err
	}
	fmt.Println(fmt.Sprintf(
		"Request status '%s'",
		res.Status,
	))
	return *res, nil
}
