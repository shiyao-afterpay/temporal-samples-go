package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	_child_workflow "github.com/temporalio/samples-go/_child-workflow"
)

// @@@SNIPSTART samples-go-child-workflow-example-worker-starter
func main() {
	// The client is a heavyweight object that should be created only once per process.
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "child-workflow", worker.Options{})

	w.RegisterWorkflow(_child_workflow.SampleParentWorkflow)
	w.RegisterWorkflow(_child_workflow.SampleChildWorkflow)
	w.RegisterWorkflow(_child_workflow.SampleChildWorkflow2)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}

// @@@SNIPEND
