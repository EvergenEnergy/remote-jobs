package addkey

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type KeysStore struct {
	authorizedKeys io.Reader
	timeFunc       func() time.Time // this allows the unit tests to force a deterministic expiration time to compare against
}

func NewKeysStore(authorizedKeys io.Reader) *KeysStore {
	return &KeysStore{authorizedKeys: authorizedKeys, timeFunc: defaultTimeFunc}
}

func defaultTimeFunc() time.Time {
	return time.Now()
}

func (p KeysStore) AddKey(publicKey string, expirationDuration time.Duration) (string, error) {
	currentAuthorizedKeys, err := io.ReadAll(p.authorizedKeys)
	if err != nil {
		return "", fmt.Errorf("failed to read currently authorized keys: %w", err)
	}

	currentKeys := strings.Split(string(currentAuthorizedKeys), "\n")

	newKeys := make([]string, 0, len(currentKeys))

	// filter out any existing entries for the same key
	for _, line := range currentKeys {
		if len(line) > 0 && !strings.Contains(line, publicKey) {
			newKeys = append(newKeys, line)
		}
	}

	expirationTime := p.timeFunc().Add(expirationDuration)

	// this is the reference time used by time.Format (January 2, 15:04:05, 2006)
	// in YYYYMMDDhhmm format
	const timeFormat = "200601021504"

	newKeys = append(newKeys, fmt.Sprintf("expiry-time=\"%s\" %s\n", expirationTime.Format(timeFormat), publicKey))

	return strings.Join(newKeys, "\n"), nil
}
