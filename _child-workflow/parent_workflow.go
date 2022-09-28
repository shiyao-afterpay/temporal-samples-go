package _child_workflow

import (
	"fmt"

	"go.temporal.io/sdk/workflow"
)

// @@@SNIPSTART samples-go-child-workflow-example-parent-workflow-definition
// SampleParentWorkflow is a Workflow Definition
// This Workflow Definition demonstrates how to start a Child Workflow Execution from a Parent Workflow Execution.
// Each Child Workflow Execution starts a new Run.
// The Parent Workflow Execution is notified only after the completion of last Run of the Child Workflow Execution.
func SampleParentWorkflow(ctx workflow.Context) error {
	fmt.Println("I AM RUNNNNNNNNING")
	logger := workflow.GetLogger(ctx)

	cwo := workflow.ChildWorkflowOptions{
		WorkflowID: "ABC-SIMPLE-CHILD-WORKFLOW-ID",
	}
	childCtx := workflow.WithChildOptions(ctx, cwo)
	wf := workflow.ExecuteChildWorkflow(childCtx, SampleChildWorkflow, "World")

	cwo2 := workflow.ChildWorkflowOptions{
		WorkflowID: "ABC-SIMPLE-CHILD-WORKFLOW-ID2",
	}
	childCtx2 := workflow.WithChildOptions(ctx, cwo2)
	workflow.ExecuteChildWorkflow(childCtx2, SampleChildWorkflow2, "World")

	err := wf.Get(ctx, nil)
	fmt.Println(err)

	logger.Info("Parent execution completed.")
	return nil
}

// @@@SNIPEND
