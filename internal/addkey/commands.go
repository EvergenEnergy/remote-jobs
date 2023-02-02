package addkey

import (
	"context"
	"flag"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/google/subcommands"
)

type AddPublicKeyCmd struct {
	publicKey string
	duration  time.Duration
}

func (*AddPublicKeyCmd) Name() string { return "add-public-key" }

func (*AddPublicKeyCmd) Synopsis() string {
	return "Add a new public key to the ~/.authorized_keys file"
}

func (*AddPublicKeyCmd) Usage() string {
	return `add-public-key --public-key <key content> --duration <duration string>:
  Add a new public key to the ~/.authorized_keys file with an expiration timestamp.
`
}

func (cmd *AddPublicKeyCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.publicKey, "public-key", "", "Public key that will be added to the device")
	f.DurationVar(&cmd.duration, "duration", 30*time.Minute, "How long should the new key be allowed to ssh for?")
}

func (cmd *AddPublicKeyCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if len(cmd.publicKey) == 0 {
		log.Println("required flag public-key is missing")
		return subcommands.ExitUsageError
	}

	currentUser, err := user.Current()
	if err != nil {
		log.Printf("failed to get current user: %v", err)
		return subcommands.ExitFailure
	}

	authKeysFile := currentUser.HomeDir + "/.ssh/authorized_keys"

	authorizedKeysFile, err := os.Open(authKeysFile)
	if err != nil {
		log.Printf("failed to open authorized_keys file: %v", err)
		return subcommands.ExitFailure
	}

	keysStore := NewKeysStore(authorizedKeysFile)

	newKeys, err := keysStore.AddKey(cmd.publicKey, cmd.duration)
	if err != nil {
		log.Printf("failed to add new key: %v", err)
		return subcommands.ExitFailure
	}

	if err = os.WriteFile(authKeysFile, []byte(newKeys), 0o600); err != nil {
		log.Printf("failed to write new content to the file: %v", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
