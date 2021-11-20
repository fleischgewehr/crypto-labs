package auth

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"io"

	"golang.org/x/crypto/argon2"
)

type argonParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

type CookedPassword struct {
	Hash      string
	ArgonSalt string
}

func getSha512Hash(password []byte) []byte {
	hasher := sha512.New()
	hasher.Write(password)

	return hasher.Sum(nil)
}

func getArgonHash(bytes, salt []byte) *CookedPassword {
	conf := &argonParams{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
	hash := argon2.IDKey(
		bytes, salt, conf.iterations, conf.memory, conf.parallelism, conf.keyLength,
	)

	return &CookedPassword{
		Hash: hex.EncodeToString(hash), ArgonSalt: hex.EncodeToString(salt),
	}
}

func HashPassword(password string) (*CookedPassword, error) {
	shaHash := getSha512Hash([]byte(password))
	salt, err := getRandomBytes(16)
	if err != nil {
		return &CookedPassword{}, err
	}
	return getArgonHash(shaHash, salt), nil
}

func CheckPassword(password string, stored *CookedPassword) bool {
	shaHash := getSha512Hash([]byte(password))
	salt, err := hex.DecodeString(stored.ArgonSalt)
	if err != nil {
		return false
	}
	argonRes := getArgonHash(shaHash, salt)

	return stored.Hash == argonRes.Hash
}

func getRandomBytes(n uint32) ([]byte, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, buf)

	return buf, err
}
