package condition

import (
	"regexp"
	"fmt"
	"errors"
	. "github.com/cifren/ghyt-api/ghyt/core/job/tools"
	. "github.com/cifren/ghyt-api/ghyt/core/config"
)

type RegexConditionType struct {

}
func(this RegexConditionType) Check(conditionConfig Condition, jobContainer *JobContainer) (bool, error) {
	arguments := conditionConfig.Arguments
	persistName := arguments["persistName"]
	variableName := arguments["variableName"]

	containerValue := jobContainer.Get(variableName)
	regex := arguments["value"]
	re := regexp.MustCompile(regex)
	matched := string(re.Find([]byte(containerValue)))

	if persistName != "" {
		jobContainer.Set(persistName, matched)
	}

	validationErrorMessage := ""
	if matched != "" {
		return true, nil
	} else {
		validationErrorMessage = fmt.Sprintf(
			"Variable '%s' with value '%s' does not match with regex '%s'",
			variableName,
			containerValue,
			regex,
		)
		return false, errors.New(validationErrorMessage)
	}
}
