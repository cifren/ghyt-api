package repository

import (
	. "github.com/cifren/youtrack/core"
	"github.com/thedevsaddam/gojsonq"
	"fmt"
	"errors"
	//"io/ioutil"
)

const (
	// Can see all tags in YT
	TAGS_ENDPOINT string = "issueTags"
	PAGIMATION_SIZE int = 400
)

type TagRepository struct {
	Client ClientInterface
	Repository RepositoryHelper
}

// Id is empty at anytime
func (this TagRepository) Find(id string) interface{} {
	return this.FindTag(id)
}

func (this TagRepository) FindTagsByName(name string) []Tag {
	request := NewRequest(TAGS_ENDPOINT)
	request.QueryParams.Add("fields", TagFields)
	request.QueryParams.Add("$top", fmt.Sprintf("%d", PAGIMATION_SIZE))
	request.QueryParams.Add("$skip", fmt.Sprintf("%d", 0))
    var tempTags []Tag
    tags := []Tag{}

	i := 0

	var jsonResult *gojsonq.Result
	var jsonErr error
	var currentPagination int
	done := false

	for respResult, respErr := this.Client.Get(*request);
		done == false;
		i = i + 1 {

		if respErr != nil {
			panic(respErr)
		}
		
		currentPagination = i * PAGIMATION_SIZE
		if respResult.Header.Get("Content-Type") != "application/json" {
			panic(errors.New(fmt.Sprintf(
				"Content-type detected is not '%s', instead '%s'",
				"application/json",
				respResult.Header.Get("Content-Type"),
			)))
		}
		
		jq := gojsonq.New().Reader(respResult.Body).Where("name", "=", name)
		
		if jq.Error() != nil {
			fmt.Println(jq.Error())
			done = true
			continue
		}
		
		// no results found on this page
		if len(jq.Get().([]interface{})) == 0 {
			fmt.Printf(
				"No results found from '%d' and '%d', for tag name '%s'\n",
				currentPagination,
				currentPagination + PAGIMATION_SIZE,
				name,
			)
			continue
		}

        jsonResult, jsonErr = jq.GetR()
        if jsonErr != nil {
            panic(jsonErr)
		}
		fmt.Printf("%#v\n", jsonResult)
		tempTags = []Tag{}
		jsonResult.As(&tempTags)

		// Means body was empty
		if len(tempTags) == 0 {
			done = true
			continue
		}

		tags = append(tags, tempTags...)

		request.QueryParams.Add("$skip", fmt.Sprintf("%d", currentPagination))

		// Decide when to finish for loop
		if i >= 1000 {
			done = true
		}
	}

	return tags
}

// Id is empty at anytime
func(this TagRepository) FindTag(id string) interface{} {
	endpoint := TAGS_ENDPOINT
	tag := Tag{}
	this.getRepository().Find(
		&tag,
		endpoint,
		this.Client,
		TagFields,
	)

	return tag
}

func(this TagRepository) getRepository() RepositoryHelper {
    return RepositoryHelper{}
}

func (this TagRepository) Flush(tagPointer interface{}) {

}