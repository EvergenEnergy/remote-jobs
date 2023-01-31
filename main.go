package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"

	"github.com/EvergenEnergy/remote-jobs/internal/addkey"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(&addkey.AddPublicKeyCmd{}, "")

	flag.Parse()

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
