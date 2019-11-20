package core

import (
	"github.com/cifren/ghyt/core/logger"
)

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
func(this Container) getLogger() logger.Logger {
	logger := logger.NewLogger(logger.DEBUG)
	return logger
}