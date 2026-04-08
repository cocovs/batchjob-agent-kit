package main

import (
	"os"

	"github.com/cocovs/batchjob-agent-kit/cli/internal/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
