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
						"value": "connect-[^-]* ",
					},
					PersistName: "yt_id",
				},
			},
			Actions: []Action{
				{
					To: "youtrack",
		 			Name: "addTag", 
					Arguments: map[string]string{
						"youtrackId": "%yt_id%", 
						"value": "nok",
					},
				},
				{
					To: "youtrack",
		 			Name: "removeTag", 
					Arguments: map[string]string{
						"youtrackId": "%yt_id%", 
						"value": "nok",
					},
				},
			},
		},
	}

	return conf
}