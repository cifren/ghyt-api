package repository

import (
	"io"
	"io/ioutil"
	"net/url"
	"net/http"
	"encoding/json"
	"bytes"
	. "github.com/cifren/ghyt-api/youtrack/core"
)

type RepositoryInterface interface {
	Find(id string) interface {}
	Flush(model interface{})
	getRepository() RepositoryHelper
}

type RepositoryHelper struct {
}

func (this RepositoryHelper) GetJson (model interface{}) io.Reader {
	s, _ := json.Marshal(model)
	b := bytes.NewBuffer(s)
	return b
}

func(this RepositoryHelper) Load(res http.Response, model interface{}) {
	// That code should be better but it crashes, let see that later
	//err3 := json.NewDecoder(res.Body).Decode(&model)

	body, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		panic(errRead)
	}

	json.Unmarshal([]byte(body), &model)
}

func(this RepositoryHelper) Find(
	model interface{},
	endpoint string,
	client ClientInterface,
	fields string,
) {
	request := Request{
		QueryParams: make(url.Values),
		Endpoint: endpoint,
	}
	q := request.QueryParams
	q.Add("fields", fields)

	res, _ := client.Get(request)
	defer res.Body.Close()

	this.Load(res, &model)
}

func(this RepositoryHelper) Flush(
	modelPointer interface{},
	endpoint string,
	client Client,
) {
	body := this.GetJson(modelPointer)

	request := Request{
		Endpoint: endpoint,
		Body: body,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	client.Post(request)
}

func(this RepositoryHelper) BuildUrl(baseUrl string, request Request) string {
	buildUrl, _ := url.Parse(baseUrl + "/" + request.Endpoint)
	buildUrl.RawQuery = request.QueryParams.Encode()

	return buildUrl.String()
}
