package cmd

import "errors"

var (
	errProjectRequired     = errors.New("project key is required. Use --project flag or set 'project' in config file")
	errEnvironmentRequired = errors.New("environment key is required. Use --environment flag or set 'environment' in config file")
)
