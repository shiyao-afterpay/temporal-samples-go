package _child_workflow

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

// @@@SNIPSTART samples-go-child-workflow-example-child-workflow-definition
// SampleChildWorkflow is a Workflow Definition
func SampleChildWorkflow2(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	greeting := "Hello " + name + "!"
	logger.Info("222222222 Child workflow execution: " + greeting)
	workflow.Sleep(ctx, 3*time.Second)

	signal := "dummy"
	fmt.Println("KKKKKKKK" + childWfRunId)
	// workflow.SignalExternalWorkflow(ctx, "ABC-SIMPLE-CHILD-WORKFLOW-ID", childWfRunId, "child-workflow-signal", signal)
	// workflow.SignalExternalWorkflow(ctx, "ABC-SIMPLE-CHILD-WORKFLOW-ID", "", "child-workflow-signal", signal)
	workflow.SignalExternalWorkflow(ctx, "ABC-SIMPLE-CHILD-WORKFLOW-ID", "", "child-workflow-signal", signal)

	return greeting, nil
}

// @@@SNIPEND
