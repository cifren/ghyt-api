package repository

import (
	. "github.com/cifren/youtrack/core"
	. "github.com/cifren/youtrack/repository"
	"github.com/thedevsaddam/gojsonq"
)

const (
	// Can see all tags in YT
	TAGS_ENDPOINT string = "issueTags"
	PAGIMATION_SIZE int = 400
)

type TagRepository struct {
	client Client
	repository RepositoryHelper
}

// Id is empty at anytime
func (this TagRepository) Find(id string) interface{} {
	return this.FindTag(id)
}

func (this TagRepository) FindTagsByName(name string) []Tag {
	request := NewRequest(TAGS_ENDPOINT)
	request.QueryParams.Add("fields", TagFields)
	request.QueryParams.Add("$top", PAGIMATION_SIZE)
	request.QueryParams.Add("$skip", 0)

	i := 0
	for res, err := this.client.Get(request); ok; ok != nil {
		jq := gojsonq.New().Reader(res)
		res := jq.From(".").Where("name", ">", 1200).OrWhere("id", "=", nil).Out()
		
		gojsonq.New().FromString(json).Where("name", "=", name).Out()
		i += PAGIMATION_SIZE
		request.QueryParams.Add("$skip", 0)
	}
}

// Id is empty at anytime
func(this TagRepository) FindTags(id string) interface{} {
	endpoint := USER_TAGS_GET_ENDPOINT
	tag := Tags{}
	this.repository.Find(
		&tags, 
		endpoint, 
		this.client, 
		TagsFields,
	)

	return tag
}

func (this TagRepository) Flush(tagPointer interface{}) {
	
}