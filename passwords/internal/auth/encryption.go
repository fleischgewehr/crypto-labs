package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/argon2"
)

type argonParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func getSecretKey() ([]byte, error) {
	return []byte("very secret key goes here"), nil
}

func getSha512Hash(password []byte) []byte {
	hasher := sha512.New()
	hasher.Write(password)

	return hasher.Sum(nil)
}

func getArgonHash(bytes []byte) ([]byte, error) {
	conf := &argonParams{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
	salt, err := getRandomBytes(conf.saltLength)
	if err != nil {
		return []byte(""), err
	}

	return argon2.IDKey(
		bytes, salt, conf.iterations, conf.memory, conf.parallelism, conf.keyLength,
	), nil
}

func encryptWithAES(bytes []byte) ([]byte, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return []byte(""), err
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return []byte(""), err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return []byte(""), err
	}
	nonce, err := getRandomBytes(uint32(aesGCM.NonceSize()))
	if err != nil {
		return []byte(""), err
	}

	return aesGCM.Seal(nonce, nonce, bytes, nil), nil
}

func decryptWithAES(encrypted string) ([]byte, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return []byte(""), err
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return []byte(""), err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return []byte(""), err
	}

	nonceLength := aesGCM.NonceSize()
	enc, err := hex.DecodeString(encrypted)
	if err != nil {
		return []byte(""), err
	}
	nonce, ciphertext := enc[:nonceLength], enc[nonceLength:]

	return aesGCM.Open(nil, nonce, ciphertext, nil)
}

func HashPassword(password string) (string, error) {
	shaHash := getSha512Hash([]byte(password))
	argonHash, err := getArgonHash(shaHash)
	if err != nil {
		return "", err
	}
	ciphertext, err := encryptWithAES(argonHash)

	return string(ciphertext), err
}

func CheckPassword(password, hash string) bool {
	shaHash := getSha512Hash([]byte(password))
	argonHash, err := getArgonHash(shaHash)
	if err != nil {
		return false
	}
	plaintext, err := decryptWithAES(string(argonHash))
	if err != nil {
		return false
	}

	return string(plaintext) == hash
}

func getRandomBytes(n uint32) ([]byte, error) {
	buf := make([]byte, n)
	_, err := rand.Read(buf)

	return buf, err
}
