package job

import(
	. "github.com/cifren/ghyt/core/job/action"
)

func ActionRetriever(to string, name string) ActionTypeInterface {
	var actionType ActionTypeInterface
	actionName := to + "-" + name
	
	switch actionName {
		case "youtrack-addTag":
			actionType = AddTagYoutrackActionType{}
		case "youtrack-removeTag":
			actionType = RemoveTagYoutrackActionType{}
	}

	return actionType
}

