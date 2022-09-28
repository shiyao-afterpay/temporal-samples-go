package _child_workflow

import (
	"fmt"

	"go.temporal.io/sdk/workflow"
)

var childWfRunId string

// @@@SNIPSTART samples-go-child-workflow-example-child-workflow-definition
// SampleChildWorkflow is a Workflow Definition
func SampleChildWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	greeting := "Hello " + name + "!"
	logger.Info("Child workflow execution: " + greeting)

	childWfRunId = workflow.GetInfo(ctx).WorkflowExecution.RunID
	// fmt.Println("kkkkkkkkkkk" + workflow.GetInfo(ctx).WorkflowExecution.RunID)

	var ret string
	selector := workflow.NewSelector(ctx)
	selector.AddReceive(workflow.GetSignalChannel(ctx, "child-workflow-signal"),
		func(rc workflow.ReceiveChannel, ok bool) {
			fmt.Println("PPPPPPPPPPPPPPP")
			rc.Receive(ctx, &ret)
			logger.Info("PPPPPPPPPPPPPP In child workflow receive")
		})

	// selector.AddDefault(
	// 	func() {
	// 		logger.Info("In child workflow begin sleep")
	// 		workflow.Sleep(ctx, 10*time.Second)
	// 		logger.Info("In child workflow end sleep")
	// 	})
	// selector.Select(ctx)

	selector.Select(ctx)
	fmt.Println("CCCCCC I am out")

	return greeting, nil
}

// @@@SNIPEND
