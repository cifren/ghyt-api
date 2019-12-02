package repository

import (
	. "github.com/cifren/youtrack/core"
	"github.com/thedevsaddam/gojsonq"
	"fmt"
)

const (
	// Can see all tags in YT
	TAGS_ENDPOINT string = "issueTags"
	PAGIMATION_SIZE int = 400
)

type TagRepository struct {
	Client Client
	Repository RepositoryHelper
}

// Id is empty at anytime
func (this TagRepository) Find(id string) interface{} {
	return this.FindTag(id)
}

func (this TagRepository) FindTagsByName(name string) []Tag {
	request := NewRequest(TAGS_ENDPOINT)
	request.QueryParams.Add("fields", TagFields)
	request.QueryParams.Add("$top", fmt.Sprintf("%s", PAGIMATION_SIZE))
	request.QueryParams.Add("$skip", fmt.Sprintf("%s", 0))
    var tempTags []Tag
    tags := []Tag{}

	i := 0

	var jsonResult *gojsonq.Result
	var jsonErr error
	for respResult, respErr := this.Client.Get(*request); respErr != nil; i = i + PAGIMATION_SIZE {
        jsonResult, jsonErr = gojsonq.New().Reader(respResult.Body).Where("name", "=", name).GetR()
        if jsonErr != nil {
            panic(jsonErr)
        }
		tempTags = []Tag{}
		jsonResult.As(&tempTags)

		append(tags, tempTags...)

		request.QueryParams.Add("$skip", fmt.Sprintf("%s", i))
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
