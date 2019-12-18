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
	
	testCases := []struct {
		name             string
		newExpectedValue int
		errExpect        bool
	}{
		{
			name: "tag1",
			newExpectedValue: 1,
			errExpect: true,
		},
	}

	assert := require.New(t)
	client := TestClient{}
	var repo TagRepository

	for _, tc := range testCases{
		repo = TagRepository{
			Client: client,
		}
		tags := repo.FindTagsByName(tc.name)
		assert.Equal(tc.newExpectedValue, len(tags))
	}
}

// ClientInterface
type TestClient struct {}
func(this TestClient) Get(request core.Request)(http.Response, error){
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	if request.QueryParams.Get("$skip") > fmt.Sprintf("%d", 400) {
		io.WriteString(w, `[]`)
	} else {
		io.WriteString(w, `[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"Infrastructure","id":"5-20","$type":"IssueTag"}]`)
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
