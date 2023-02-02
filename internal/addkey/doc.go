// Package addkey defines a custom job handler
// that will take 2 parameters:
// - a public key
// - a time duration
// The handler will add the key to the ~/.authorized_keys file
// where it will be valid according to the expiration time passed in.
// If the same public key had been added already in the past
// those previous entries will be removed to avoid cluttering the file.
// The expiration feature relies on the `expiry-time` option available in OpenSSH since
// version 7.7 (https://www.openssh.com/txt/release-7.7)
package addkey
