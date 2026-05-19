// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package credentials // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/sumologicextension/internal/credentials"

func _getHasher() Hasher { _ = "STUB: not implemented"; return *new(Hasher) }

const (
	filenamePrefix      = "filename"
	encryptionKeyPrefix = "encryption"
)

type Hasher interface {
	Write(p []byte) (n int, err error)
	Sum(b []byte) []byte
}

func hashWith(hasher Hasher, key []byte) (string, error) { _ = "STUB: not implemented"; return "", nil }

// HashKeyToFilename creates a filename using the default hasher and provided key
// as input. It returns this filename and an error.
func HashKeyToFilename(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// HashKeyToFilenameWith creates a filename using the provided key as input and
// using the provided hasher.
func HashKeyToFilenameWith(hasher Hasher, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// HashKeyToEncryptionKey creates an encryption key using a default hasher.
// It returns the created key and an error.
func HashKeyToEncryptionKey(key string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// HashKeyToEncryptionKeyWith creates a 32 bytes long key from the provided
// key using the provided hasher.
func HashKeyToEncryptionKeyWith(hasher Hasher, key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// encrypt encrypts provided byte slice with AES using the encryption key.
func encrypt(data, encryptionKey []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// decrypt decrypts provided byte slice with AES using the encryptionKey.
func decrypt(data, encryptionKey []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
