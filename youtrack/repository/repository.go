package repository

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"encoding/json"
	. "github.com/cifren/ghyt-api/youtrack/core"
)

type RepositoryInterface interface {
	Find(id string) (interface {}, error)
	Flush(model interface{})
	getRepository() RepositoryHelper
}

type RepositoryHelper struct {
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
) error {
	request := Request{
		QueryParams: make(url.Values),
		Endpoint: endpoint,
	}
	q := request.QueryParams
	q.Add("fields", fields)

	res, err := client.Get(request)

	if err != nil {
	  return err
	}
	defer res.Body.Close()

	this.Load(res, &model)
	return nil
}

func(this RepositoryHelper) Flush(
	modelPointer interface{},
	endpoint string,
	client ClientInterface,
	fields string,
	customData interface{},
) {
	var body interface{}
	if customData == nil {
		body = modelPointer
	} else {
		body = customData
	}

	request := Request{
		QueryParams: make(url.Values),
		Endpoint: endpoint,
		Body: body,
	}
	q := request.QueryParams
	q.Add("fields", fields)

	res, _ := client.Post(request)
	defer res.Body.Close()

	this.Load(res, &modelPointer)
}

func(this RepositoryHelper) BuildUrl(baseUrl string, request Request) string {
	buildUrl, _ := url.Parse(baseUrl + "/" + request.Endpoint)
	buildUrl.RawQuery = request.QueryParams.Encode()

	return buildUrl.String()
}
