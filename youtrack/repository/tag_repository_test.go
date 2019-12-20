package repository

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io"
	"strconv"

	"github.com/stretchr/testify/require"
	"github.com/cifren/youtrack/core"
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
		{
			name: "search tag2",
			tagSearch: "tag2",
			dataPages: []string{
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag2","id":"5-20","$type":"IssueTag"}]`,
			},
			paginationSize: 3,
			newExpectedValue: 2,
			errExpect: false,
		},
		{
			name: "search tag3",
			tagSearch: "tag3",
			dataPages: []string{
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag2","id":"5-20","$type":"IssueTag"}]`,
			},
			paginationSize: 3,
			newExpectedValue: 0,
			errExpect: false,
		},
		// tes4 : 3 pages, result on page 3
		// test : 3 pages, result on page 2
		// test : 3 pages, result not found
		// test : 4 pages, result on page 1
		// test : 4 pages, result not found
		// test : 1 page, http return 500
		// test : 1 page, http return 401
		// test : 2 page, http return timeout on page 2
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
			tags := repo.FindTagsByName(tc.tagSearch)
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
