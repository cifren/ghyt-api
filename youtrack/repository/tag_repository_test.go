package repository

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io"

	"github.com/stretchr/testify/require"
	"github.com/cifren/youtrack/core"

	"fmt"
)

func TestFindTagsByName(t *testing.T) {
	assert := require.New(t)

	client := TestClient{}

	repo := TagRepository{
		Client: client,
	}

	tags := repo.FindTagsByName("tag1")
	assert.Equal(len(tags), 0)
}

// ClientInterface
type TestClient struct {}
func(this TestClient) Get(request core.Request)(http.Response, error){
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	if request.QueryParams.Get("$skip") > fmt.Sprintf("%d", 400) {
		io.WriteString(w, `[]`)
		fmt.Printf("%v\n", request.QueryParams.Get("$skip"))
	} else {
		io.WriteString(w, `[{"name":"Backend","id":"5-17","$type":"IssueTag"},{"name":"Frontend","id":"5-18","$type":"IssueTag"},{"name":"Infrastructure","id":"5-20","$type":"IssueTag"}]`)
		fmt.Printf("%v\n", request.QueryParams.Get("$skip"))
	}
	resp := w.Result()

	return *resp, nil
}
func(this TestClient) Post(request core.Request)(http.Response, error){
	return http.Response{}, nil
}
func(this TestClient) Request(request core.Request)(http.Response, error){
	return http.Response{}, nil
}
