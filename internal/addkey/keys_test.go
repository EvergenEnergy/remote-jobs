package addkey

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestAppendPublicKey(t *testing.T) {
	// hardcode an arbitrary timestamp (2023-1-10 10:30)
	// to get deterministic behaviour in the tests below
	currentTime := time.Date(2023, time.January, 10, 10, 30, 0, 0, time.UTC)
	publicKey := "my-new-public-key"

	tests := []struct {
		name        string
		currentKeys string
		expiration  string
		wantOutput  string
	}{
		{
			name:        "Empty input, add key for 30 minutes",
			currentKeys: "",
			expiration:  "30m",
			wantOutput:  fmt.Sprintf("expiry-time=\"202301101100\" %s\n", publicKey),
		},
		{
			name:        "Empty input, add key for 8 hours",
			currentKeys: "",
			expiration:  "8h",
			wantOutput:  fmt.Sprintf("expiry-time=\"202301101830\" %s\n", publicKey),
		},
		{
			name:        "Other keys are not touched",
			currentKeys: "expiry-time=\"202301101830\" other-public-key\n",
			expiration:  "30m",
			wantOutput:  fmt.Sprintf("expiry-time=\"202301101830\" other-public-key\nexpiry-time=\"202301101100\" %s\n", publicKey),
		},
		{
			name:        "Removes old entries for the same key (it effectively updates the expiration timestamp)",
			currentKeys: fmt.Sprintf("expiry-time=\"202201101830\" %s\n", publicKey),
			expiration:  "30m",
			wantOutput:  fmt.Sprintf("expiry-time=\"202301101100\" %s\n", publicKey),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := KeysStore{
				authorizedKeys: strings.NewReader(test.currentKeys),
				timeFunc:       func() time.Time { return currentTime },
			}

			expiration, err := time.ParseDuration(test.expiration)
			assert.NilError(t, err)

			output, err := store.AddKey(publicKey, expiration)
			assert.NilError(t, err)

			assert.Equal(t, output, test.wantOutput)
		})
	}
}
