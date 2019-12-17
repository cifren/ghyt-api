package condition

import (
	. "github.com/cifren/ghyt/core/job/tools"
	. "github.com/cifren/ghyt/core/config"
	"fmt"
)

type EqualConditionType struct {

}
func(this EqualConditionType) Check(conditionConfig Condition, jobContainer *JobContainer) (bool, string) {
	arguments := conditionConfig.Arguments
	variableName := arguments["variableName"]

	containerValue := jobContainer.Get(variableName)
	proposedValue := arguments["value"]

	validationErrorMessage := ""
	if containerValue == proposedValue {
		return true, validationErrorMessage
	} else {
		validationErrorMessage = fmt.Sprintf(
			"Variable '%s' with value '%s' does not match with value '%s'", 
			variableName,
			containerValue,
			proposedValue,
		)
		return false, validationErrorMessage
	}
}
