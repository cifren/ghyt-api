package condition

import (
	. "github.com/cifren/ghyt/core/job/tools"
	. "github.com/cifren/ghyt/core/config"
	"regexp"
	"fmt"
)

type RegexConditionType struct {

}
func(this RegexConditionType) Check(conditionConfig Condition, jobContainer JobContainer) (bool, string) {
	arguments := conditionConfig.Arguments
	variableName := arguments["variableName"]

	containerValue := jobContainer.Get(variableName)
	proposedValue := arguments["value"]
	matched, _ := regexp.Match(proposedValue, []byte(containerValue))

	validationErrorMessage := ""
	if matched {
		return true, validationErrorMessage
	} else {
		validationErrorMessage = fmt.Sprintf(
			"Variable '%s' with value '%s' does not match with regex '%s'", 
			variableName,
			containerValue,
			proposedValue,
		)
		return false, validationErrorMessage
	}
}