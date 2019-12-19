package repository

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io"
	"strconv"

	"github.com/stretchr/testify/require"
	"github.com/cifren/youtrack/core"

	"fmt"
)

func TestFindTagsByName(t *testing.T) {

	testCases := []struct {
		name             string
		tagSearch        string
		dataPages		 []string
		paginationSize	 int
		newExpectedValue int
		errExpect        bool
	}{
		{
			name: "search tag1",
			tagSearch: "tag1",
			dataPages: []string{
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag2","id":"5-20","$type":"IssueTag"}]`,
			},
			paginationSize: 3,
			newExpectedValue: 1,
			errExpect: false,
		},
// 		{
// 			name: "search tag2",
// 			tagSearch: "tag2",
// 			dataPages: []string{
// 				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag2","id":"5-20","$type":"IssueTag"}]`,
// 			},
// 			paginationSize: 3,
// 			newExpectedValue: 2,
// 			errExpect: false,
// 		},
// 		{
// 			name: "search tag3",
// 			tagSearch: "tag3",
// 			dataPages: []string{
// 				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag2","id":"5-20","$type":"IssueTag"}]`,
// 			},
// 			paginationSize: 3,
// 			newExpectedValue: 0,
// 			errExpect: false,
// 		},
	}

	assert := require.New(t)

	for _, tt := range testCases{
		tc := tt
		client := TestClient{DataPages: tc.dataPages}

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			repo := TagRepository{
				Client: client,
				PaginationSize: tc.paginationSize,
			}
			tags := repo.FindTagsByName(tc.name)
			if !tc.errExpect {
				assert.Equal(tc.newExpectedValue, len(tags))
			}
		})
	}
}

// ClientInterface
type TestClient struct {
	DataPages []string
}
func(this TestClient) Get(request core.Request)(http.Response, error){
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	paginationSize, _ := strconv.Atoi(request.QueryParams.Get("$top"))
	skipSize, _ := strconv.Atoi(request.QueryParams.Get("$skip"))
	pageNumber := (skipSize + paginationSize) / paginationSize

	fmt.Printf("pageNumber %v\n", pageNumber)
	fmt.Printf("len(this.DataPages) %v\n", len(this.DataPages))
	if len(this.DataPages) >= pageNumber {
		io.WriteString(w, this.DataPages[pageNumber-1])
	} else {
		io.WriteString(w, `[]`)
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
