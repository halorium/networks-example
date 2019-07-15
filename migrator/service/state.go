package main

import (
	"os"

	"github.com/halorium/networks-example/aws/client"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/postgres"
)

type State struct {
	AWSClient    *awsclient.Client `json:"-"`
	PGClient     *postgres.Client  `json:"-"`
	CreatedAt    string            `json:"created-at"`
	FlawError    *flaw.Error       `json:"flaw-error"`
	File         os.FileInfo       `json:"-"`
	FileID       string            `json:"file-id"`
	FileName     string            `json:"file-name"`
	FilepathName string            `json:"filepath-name"`
	Files        []os.FileInfo     `json:"-"`
	Message      string            `json:"message"`
}

func NewState() *State {
	state := &State{}

	state.AWSClient = awsclient.FromPool()
	state.PGClient = postgres.ClientFromPool()

	return state
}
