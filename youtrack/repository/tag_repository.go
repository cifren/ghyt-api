package repository

import (
	. "github.com/cifren/ghyt-api/youtrack/core"
	"github.com/thedevsaddam/gojsonq"
	"fmt"
	"errors"
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

const (
	// Can see all tags in YT
	TAGS_ENDPOINT string = "issueTags"
)

type TagRepository struct {
	Client ClientInterface
	ItemsPerPage int
}

// Id is empty at anytime
func (this TagRepository) Find(id string) interface{} {
	return this.FindTag(id)
}

func (this TagRepository) FindTagsByName(name string) []Tag {
	itemsPerPage := this.ItemsPerPage
	if itemsPerPage == 0 {
		panic("Value 0 is not possible, please select another 'top' value")
	}

	request := NewRequest(TAGS_ENDPOINT)
	request.QueryParams.Add("fields", TagFields)
	request.QueryParams.Add("$top", fmt.Sprintf("%d", itemsPerPage))
	request.QueryParams.Add("$skip", fmt.Sprintf("%d", 0))
    var tempTags []Tag
    tags := []Tag{}

	i := 0

	var itemPosition int
	var nextItemPosition int
	fmt.Printf("Run #1 look for #name %s, Request '%#v'\n", name, request)
	for done := false;
		done == false;
		i = i + 1 {
		respResult, respErr := this.Client.Get(*request)

		if respErr != nil {
			panic(respErr)
		}

		if !strings.Contains(respResult.Header.Get("Content-Type"), "application/json") {
			panic(errors.New(fmt.Sprintf(
				"Content-type detected is not '%s', instead '%s'",
				"application/json",
				respResult.Header.Get("Content-Type"),
			)))
		}

		jq := gojsonq.New().Reader(respResult.Body)
		// it means page is empty so request should stop
		if jq.Count() <= 0 {
			fmt.Println("Page is empty")
			done = true
			continue
		}

		result := jq.Where("name", "=", name)

		// Usually when Json is empty or malformed, it means no more results
		if result.Error() != nil {
			fmt.Printf("Usually end of response or malformed json, Jsonq error '%v'\n", jq.Error())
			// Exit for loop
			done = true
			continue
		}

		itemPosition = i * itemsPerPage
		nextItemPosition = itemPosition + itemsPerPage
		// no results found on this page
		if len(result.Get().([]interface{})) == 0 {
			fmt.Printf(
				"No results found from row '%d' to '%d', for tag name '%s'\n",
				itemPosition,
				nextItemPosition - 1,
				name,
			)
		} else {
			var resultBytes bytes.Buffer
			tempTags = []Tag{}
			result.Writer(&resultBytes)
			json.Unmarshal(resultBytes.Bytes(), &tempTags)

			tags = append(tags, tempTags...)
		}

		fmt.Printf("Skip '%v'\n", nextItemPosition)
		request.QueryParams.Set("$skip", strconv.Itoa(nextItemPosition))

		// Decide when to finish for loop
		if i >= 1000 {
			panic("Too many loop for Repository, means 1000 pages have been searched before being killed")
		}
		if !done {
			fmt.Printf("Run #%d for '%v'\n", i+2, request)
		}
	}
	fmt.Printf("Results found : %d\n", len(tags))

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
	myTag := tagPointer.(*Tag)
    this.FlushTag(myTag)
}

func (this TagRepository) FlushTag(tag *Tag) {
	endpoint := TAGS_ENDPOINT
	if (*tag).Id != "" {
		endpoint = endpoint + "/" + (*tag).Id
	}

	jsonTag := struct{
		Name string `json:"name"`
	}{
		Name: tag.Name,
	}

	this.getRepository().Flush(tag, endpoint, this.Client, TagFields, jsonTag)

	fmt.Printf("%#v\n", tag)
}
