package job

import (
	"fmt"
	. "github.com/cifren/ghyt-api/ghyt/core/job/condition"
)

func ConditionRetriever(name string) ConditionTypeInterface {
	var conditionType ConditionTypeInterface

	switch name {
		case "equal":
			conditionType = EqualConditionType{}
		case "regex":
			conditionType = RegexConditionType{}
		default:
			panic(fmt.Sprintf("Condition type not found, given : %#v", name))
	}

	return conditionType
}
