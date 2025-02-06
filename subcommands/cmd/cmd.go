package main

import (
	"context"
	"flag"
	"log/slog"
	"strings"

	"github.com/google/subcommands"
)

type TestCommand struct {
	test string
}

func (*TestCommand) Name() string     { return "test" }
func (*TestCommand) Synopsis() string { return "test command" }
func (*TestCommand) Usage() string {
	return `test:
This is test command
`
}
func (t *TestCommand) SetFlags(f *flag.FlagSet) {
}

func (t *TestCommand) Execute(c context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	slog.Info(strings.Join(f.Args(), " "))
	slog.Info("いえーーーい")
	return subcommands.ExitSuccess
}
