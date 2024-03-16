package task2

import (
	"go.temporal.io/sdk/workflow"
)

// This sample workflow demonstrates how to use invoke child workflow from parent workflow execution.  Each child
// workflow execution is starting a new run and parent execution is notified only after the completion of last run.

// SampleParentWorkflow workflow definition
func SampleParentWorkflow(ctx workflow.Context) (string, error) {
	logger := workflow.GetLogger(ctx)
	cwo := workflow.ChildWorkflowOptions{}
	ctx = workflow.WithChildOptions(ctx, cwo)

	doneCh := workflow.NewChannel(ctx)

	for i := 0; i < 200; i++ {
		workflow.Go(ctx, func(ctx workflow.Context) {
			var result string
			err := workflow.ExecuteChildWorkflow(ctx, SampleChildWorkflow, "april").Get(ctx, &result)
			if err != nil {
				logger.Error("Parent execution received child execution failure.", "Error", err)
				return
			}
			doneCh.Send(ctx, result)
		})
	}

	results := ""
	for i := 0; i < 200; i++ {
		var v string
		doneCh.Receive(ctx, &v)
		results += v + ";"

	}

	logger.Info("Parent execution completed.")
	return results, nil
}
