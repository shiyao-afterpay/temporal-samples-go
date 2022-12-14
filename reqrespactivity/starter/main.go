package main

import (
	"context"
	"log"

	"github.com/temporalio/samples-go/reqrespactivity"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "reqrespactivity_workflow",
		TaskQueue: "reqrespactivity",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, reqrespactivity.UppercaseWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
