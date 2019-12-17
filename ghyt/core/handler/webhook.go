package handler

import (
	"fmt"
	// "gopkg.in/go-playground/webhooks.v5/github"
	"github.com/kataras/iris"
	. "github.com/cifren/ghyt/core"
	. "github.com/cifren/ghyt/core/logger"
	. "github.com/cifren/ghyt/core/job"
	. "github.com/cifren/ghyt/core/job/tools"
	. "github.com/cifren/ghyt/core/config"
	"github.com/cifren/ghyt/config"
	// "reflect"
	// "github.com/cifren/ghyt/core/event"
	// "github.com/cifren/ghyt/core/action"

)

func GhWebhookHandler(ctx iris.Context, container Container)  {
	// hook, _ := github.New(github.Options.Secret(""))
	// payload, err := hook.Parse(ctx.Request(), github.PingEvent, github.PushEvent, github.PullRequestEvent)
	// if err != nil {
	// 	if err == github.ErrEventNotFound {
	// 		// ok event wasn't one of the ones asked to be parsed
	// 		fmt.Println(github.ErrEventNotFound)
	// 	} else {
	// 		fmt.Printf("Error on payload: %+v\n", err)
	// 	}
	// }
	// fmt.Printf("Payload type: %+v\n", reflect.TypeOf(payload))
	// switch payload.(type) {
	// 	case github.PushPayload:
	// 		release := payload.(github.PushPayload)
	// 		// Do whatever you want from here...
	// 		fmt.Printf("%+v\n", release)
	// 	case github.PullRequestPayload:
	// 		release := payload.(github.PullRequestPayload)

			//event := event.NewPullRequestEvent()
			// event.Variables["id"] = strconv.FormatInt(release.PullRequest.Number, 10)
			// event.Variables["description"] = string(release.PullRequest.Title)

			// myAction := action.AddTag{Name: event.GetVariables()["id"]}

			// runner := action.ActionRunner{}
			// runner.RunIt(myAction)

	// 	case github.PingPayload:
	// 		release := payload.(github.PingPayload)
	// 		// Do whatever you want from here...
	// 		fmt.Printf("%+v\n", release)
	// 	default:
	// 		fmt.Printf("Event without payload : %+v\n", reflect.TypeOf(payload))
	// }

	//event := event.EventManager.GetGhEvent(payload)


	conf := config.GetConf()

	jobContainer := NewJobContainer()
	jobContainer.Set("event.pull_request.state", "open")
	jobContainer.Set("event.pull_request.title", "connect-1517 lol")

	for _, job := range conf {
		runJob(jobContainer, job, container)
	}
}

func runJob(jobContainer *JobContainer, job Job, container Container) {
	logger := container.Get("logger").(Logger)

	logger.Debug(fmt.Sprintf(
		"Conditions found: %x",
		len(job.Conditions),
	))

	conditionChecker := ConditionChecker{}
	for _, condition := range job.Conditions {
		// varContainer is in ref in case persistName has been set
		if !conditionChecker.Check(condition, jobContainer, logger) {
			logger.Debug(fmt.Sprintf(
				"Condition refused '%s'",
				condition.Name,
			))
			// quit without executing actions
			return
		}
		logger.Debug(fmt.Sprintf("Condition success '%s'", condition.Name))
	}
	logger.Debug(fmt.Sprintf(
		"Actions found: %x",
		len(job.Actions),
	))
	actionRunner := container.Get("actionRunner").(ActionRunner)
	// run all actions
	for _, action := range job.Actions {
		logger.Debug(fmt.Sprintf(
			"Run action '%s'",
			action.Name,
		))
		actionRunner.Run(action, *jobContainer)
	}
}

type Container struct {
	All map[string]interface{}
}
func(this *Container) InitContainer() {
	this.All = make(map[string]interface{})
	this.All["logger"] = this.getLogger()
}
func(this Container) Get(reference string) interface{} {
	return this.All[reference]
}
func(this Container) getLogger() Logger {
	logger := NewLogger(DEBUG)
	return logger
}

func getConfs() []Conf {

	confs := []Conf {
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

	return confs
}

type Conf struct {
	Conditions []Condition
	Actions []Action
}

type VarContainer struct {
	Variables map[string]string
}
func NewVarContainer() *VarContainer {
	v := new(VarContainer)
	v.Variables = make(map[string]string)
	return v
}
func(this VarContainer) get(name string) string {
	return this.Variables[name]
}
func(this *VarContainer) set(name string, value string) {
	this.Variables[name] = value
}

type VariableRetriever struct {

}

type Condition struct {
	Name string
	Arguments map[string]string
	PersistName string
}

type Action struct {
	To string
	Name string
	Arguments map[string]string
}

type ConditionChecker struct {

}
func (this ConditionChecker) Check(
		conditionConfig Condition,
		varContainer VarContainer,
		logger Logger,
	) bool {
	conditionType := ConditionRetriever(conditionConfig.Name)

	result, validationErrorMessage := conditionType.Check(
		conditionConfig,
		varContainer,
	)
	if !result {
		logger.Debug(validationErrorMessage)
	}

	return result
}


type ActionRunner struct {

}
func (this ActionRunner) Run(actionConfig Action, varContainer VarContainer) {
	actionType := ActionRetriever(actionConfig.To, actionConfig.Name)

	actionType.Run(
		actionConfig,
		varContainer,
	)
}

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

type EqualConditionType struct {

}
func(this EqualConditionType) Check(conditionConfig Condition, varContainer VarContainer) (bool, string) {
	arguments := conditionConfig.Arguments
	variableName := arguments["variableName"]

	containerValue := varContainer.get(variableName)
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

type RegexConditionType struct {

}
func(this RegexConditionType) Check(conditionConfig Condition, varContainer VarContainer) (bool, string) {
	arguments := conditionConfig.Arguments
	variableName := arguments["variableName"]

	containerValue := varContainer.get(variableName)
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

type ConditionTypeInterface interface {
	Check(Condition, VarContainer) (bool, string)
}


type ActionTypeInterface interface {
	Run(Action, VarContainer)
}

type AddTagYoutrackActionType struct {

}
func(this AddTagYoutrackActionType) Run(actionConfig Action, varContainer VarContainer) {

}

type RemoveTagYoutrackActionType struct {

}
func(this RemoveTagYoutrackActionType) Run(actionConfig Action, varContainer VarContainer) {

}
