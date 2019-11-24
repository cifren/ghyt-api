package config 

import (
	. "github.com/cifren/ghyt/core/config"
)

func GetConf() []Job {
	
	conf := []Job {
		{
			Conditions: []Condition{
				{
					Name: "equal",
					Arguments: map[string]string{
						"variableName": "event.pull_request.state", 
						"value": "open",
					},
				},
				{
					Name: "regex",
					Arguments: map[string]string{
						"variableName": "event.pull_request.title", 
						"value": "connect-[^-][0-9]*",
						"persistName": "yt_id",
					},
				},
			},
			Actions: []Action{
				{
					To: "youtrack",
		 			Name: "addTag", 
					Arguments: map[string]string{
						"youtrackId": "yt_id", 
						"tagName": "nok",
					},
				},
				{
					To: "youtrack",
		 			Name: "removeTag", 
					Arguments: map[string]string{
						"youtrackId": "yt_id", 
						"tagName": "nok",
					},
				},
			},
		},
	}

	return conf
}