package job

import(
	"fmt"
	. "github.com/cifren/ghyt-api/ghyt/core/job/action"
)

func ActionRetriever(to string, name string) ActionTypeInterface {
	var actionType ActionTypeInterface
	actionName := to + "-" + name

	switch actionName {
		case "youtrack-addTag":
			actionType = AddTagYoutrackActionType{}
		case "youtrack-removeTag":
			actionType = RemoveTagYoutrackActionType{}
		default:
			panic(fmt.Sprintf("Action type not found, given : %#v", actionName))
	}

	return actionType
}

