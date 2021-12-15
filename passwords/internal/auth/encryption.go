package auth

import (
	"crypto/aes"
	"crypto/cipher"
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

type CookedCipher struct {
	Cipher string
	Salt   string
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

func encryptWithAESGCM(plain, secret string) (*CookedCipher, error) {
	key := []byte(secret)
	bytes := []byte(plain)
	block, err := aes.NewCipher(key)
	if err != nil {
		return &CookedCipher{}, err
	}

	nonce, err := getRandomBytes(12)
	if err != nil {
		return &CookedCipher{}, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return &CookedCipher{}, err
	}
	cipher := aesgcm.Seal(nil, nonce, bytes, nil)

	return &CookedCipher{
		Cipher: hex.EncodeToString(cipher),
		Salt:   hex.EncodeToString(nonce),
	}, nil
}

func getPhoneSecret() string {
	// return os.Getenv("AES_PHONE_SECRET_KEY")
	return "AES256Key-32Characters1234567890"
}

func getAddressSecret() string {
	// return os.Getenv("AES_ADDRESS_SECRET_KEY")
	return "AES256Key-32Characters1234567890"
}

func EncryptPhone(phone string) (*CookedCipher, error) {
	secret := getPhoneSecret()

	return encryptWithAESGCM(phone, secret)
}

func EncryptAddress(address string) (*CookedCipher, error) {
	secret := getAddressSecret()

	return encryptWithAESGCM(address, secret)
}

func decryptAESGCM(ciphertext *CookedCipher, secret string) (string, error) {
	bytes, _ := hex.DecodeString(ciphertext.Cipher)
	nonce, _ := hex.DecodeString(ciphertext.Salt)
	key := []byte(secret)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, nonce, bytes, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func DecryptPhone(ciphertext *CookedCipher) (string, error) {
	secret := getPhoneSecret()

	return decryptAESGCM(ciphertext, secret)
}

func DecryptAddress(ciphertext *CookedCipher) (string, error) {
	secret := getAddressSecret()

	return decryptAESGCM(ciphertext, secret)
}
