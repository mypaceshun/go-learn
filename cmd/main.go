package main

import (
	"context"
	"errors"
	"flag"

	"github.com/google/subcommands"
	"github.com/mypaceshun/go-learn/log"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&TestCommand{}, "")

	flag.Parse()
	ctx := context.Background()
	if rc := int(subcommands.Execute(ctx)); rc != 0 {
		log.ErrExit(errors.New("エラーだね"), rc)
	}
}
