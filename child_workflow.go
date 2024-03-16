package task2

import (
	"go.temporal.io/sdk/workflow"
)

// SampleChildWorkflow workflow definition
func SampleChildWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	greeting := "Hi " + name + "!"
	logger.Info("Child workflow execution: " + greeting)
	return greeting, nil
}
