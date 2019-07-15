package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/git"
	"github.com/halorium/networks-example/halt"
)

func main() {
	branchName, flawError := git.CurrentBranch()

	if flawError != nil {
		halt.Panic(flawError)
	}

	_, err := fmt.Fprintln(os.Stderr, "retrying redrive messages for", branchName)

	if err != nil {
		halt.Panic(flaw.From(err))
	}

	awsSession, err := session.NewSession()

	if err != nil {
		halt.Panic(flaw.From(err))
	}

	flawError = iterateQueues(
		&state{
			sqsClient:  sqs.New(awsSession),
			branchName: branchName,
		},
	)

	if flawError != nil {
		halt.Panic(flawError)
	}
}
