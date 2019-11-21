package job

import (	
	. "github.com/cifren/ghyt/core/job/condition"
)

func ConditionRetriever(name string) ConditionTypeInterface {
	var conditionType ConditionTypeInterface

	switch name {
		case "equal":
			conditionType = EqualConditionType{}
		case "regex":
			conditionType = RegexConditionType{}
	}

	return conditionType
}
