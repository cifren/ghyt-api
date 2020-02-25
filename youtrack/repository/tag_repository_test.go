package repository

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io"
	"strconv"

	"github.com/stretchr/testify/require"
	"github.com/cifren/ghyt-api/youtrack/core"

	"fmt"
)

func TestFindTagsByName(t *testing.T) {

	testCases := []struct {
		name             string
		// which tags is we are looking for
		tagSearch        string
		// Each row is a page, should contain JSON and as many items as itemsPerPage in it
		dataPages		 []string
		// Number of items per page
		itemsPerPage	 int
		// Number of items returned by TagRepository
		newExpectedValue int
		errExpect        bool
	}{
		{
			name: "search tag1",
			tagSearch: "tag1",
			dataPages: []string{
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag2","id":"5-20","$type":"IssueTag"}]`,
			},
			itemsPerPage: 3,
			newExpectedValue: 1,
			errExpect: false,
		},
		{
			name: "search tag2",
			tagSearch: "tag2",
			dataPages: []string{
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag2","id":"5-20","$type":"IssueTag"}]`,
			},
			itemsPerPage: 3,
			newExpectedValue: 2,
			errExpect: false,
		},
		{
			name: "search tag3",
			tagSearch: "tag3",
			dataPages: []string{
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag2","id":"5-20","$type":"IssueTag"}]`,
			},
			itemsPerPage: 3,
			newExpectedValue: 0,
			errExpect: false,
		},
		{ // tes4 : 3 pages, result on page 3
			name: "search tag4",
			tagSearch: "tag4",
			dataPages: []string{
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag2","id":"5-20","$type":"IssueTag"}]`,
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag2","id":"5-20","$type":"IssueTag"}]`,
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag4","id":"5-20","$type":"IssueTag"}]`,
			},
			itemsPerPage: 3,
			newExpectedValue: 1,
			errExpect: false,
		},
		{ // tes4 : 3 pages, result not found
			name: "search tag4",
			tagSearch: "tag5",
			dataPages: []string{
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag2","id":"5-20","$type":"IssueTag"}]`,
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag2","id":"5-20","$type":"IssueTag"}]`,
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag4","id":"5-20","$type":"IssueTag"}]`,
			},
			itemsPerPage: 3,
			newExpectedValue: 0,
			errExpect: false,
		},
		{ // test : 3 pages, result on page 2
			name: "search tag6",
			tagSearch: "tag6",
			dataPages: []string{
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag2","id":"5-20","$type":"IssueTag"}]`,
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag6","id":"5-20","$type":"IssueTag"}]`,
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag4","id":"5-20","$type":"IssueTag"}]`,
			},
			itemsPerPage: 3,
			newExpectedValue: 1,
			errExpect: false,
		},
		{ // test : 4 pages, result on page 1
			name: "search tag7",
			tagSearch: "tag7",
			dataPages: []string{
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag7","id":"5-20","$type":"IssueTag"}]`,
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag4","id":"5-18","$type":"IssueTag"},{"name":"tag6","id":"5-20","$type":"IssueTag"}]`,
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag5","id":"5-18","$type":"IssueTag"},{"name":"tag4","id":"5-20","$type":"IssueTag"}]`,
				`[{"name":"tag6","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag4","id":"5-20","$type":"IssueTag"}]`,
			},
			itemsPerPage: 3,
			newExpectedValue: 1,
			errExpect: false,
		},
		{ // test : 4 pages, result not found
			name: "search tag8",
			tagSearch: "tag8",
			dataPages: []string{
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag7","id":"5-20","$type":"IssueTag"}]`,
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag4","id":"5-18","$type":"IssueTag"},{"name":"tag8","id":"5-20","$type":"IssueTag"}]`,
				`[{"name":"tag8","id":"5-17","$type":"IssueTag"},{"name":"tag5","id":"5-18","$type":"IssueTag"},{"name":"tag4","id":"5-20","$type":"IssueTag"}]`,
				`[{"name":"tag6","id":"5-17","$type":"IssueTag"},{"name":"tag8","id":"5-18","$type":"IssueTag"},{"name":"tag4","id":"5-20","$type":"IssueTag"}]`,
			},
			itemsPerPage: 3,
			newExpectedValue: 3,
			errExpect: false,
		},
		{ // test : 4 pages, result not found
			name: "search tag9",
			tagSearch: "tag8",
			dataPages: []string{
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag7","id":"5-20","$type":"IssueTag"}]`,
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag4","id":"5-18","$type":"IssueTag"},{"name":"tag6","id":"5-20","$type":"IssueTag"}]`,
				`[{"name":"tag1","id":"5-17","$type":"IssueTag"},{"name":"tag5","id":"5-18","$type":"IssueTag"},{"name":"tag4","id":"5-20","$type":"IssueTag"}]`,
				`[{"name":"tag6","id":"5-17","$type":"IssueTag"},{"name":"tag2","id":"5-18","$type":"IssueTag"},{"name":"tag4","id":"5-20","$type":"IssueTag"}]`,
			},
			itemsPerPage: 3,
			newExpectedValue: 0,
			errExpect: false,
		},
	}

	assert := require.New(t)

	for _, tt := range testCases{
		tc := tt
		client := TestClient{DataPages: tc.dataPages}

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			repo := TagRepository{
				Client: client,
				ItemsPerPage: tc.itemsPerPage,
			}
			tags, _ := repo.FindTagsByName(tc.tagSearch)
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
	itemsPerPage, _ := strconv.Atoi(request.QueryParams.Get("$top"))
	skipSize, _ := strconv.Atoi(request.QueryParams.Get("$skip"))
	pageNumber := (skipSize + itemsPerPage) / itemsPerPage
	fmt.Printf("$skip %s, $top %s\n",
		request.QueryParams.Get("$skip"),
		request.QueryParams.Get("$top"),
	)
	fmt.Printf("skip %d, itemsPerPage %d, pageNumber %d\n", skipSize, itemsPerPage, pageNumber)

	if len(this.DataPages) >= pageNumber {
		io.WriteString(w, this.DataPages[pageNumber-1])
	} else {
		fmt.Printf("No more page\n")
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
