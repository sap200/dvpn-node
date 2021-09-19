package utils

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"encoding/hex"
	"io"
	"math/rand"
	"time"
)

// GenerateRandomAESKey generates a random AES Key
func GenerateRandomAESKey() []byte {
	rand.Seed(time.Now().UnixNano())
	bs := make([]byte, 32)

	for i := 0; i < 32; i++ {
		bs[i] = byte(rand.Intn(256))
	}

	return bs

}

// EncryptAES encrypts the plaintext using key
func EncryptAES(key []byte, plaintext []byte) string {
	// create cipher
	block, err := aes.NewCipher(key)
	check(err)

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(crand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	// return hex string
	return hex.EncodeToString(ciphertext) + "\n"
}

// DecryptAES decrypts ciphertext using key
func DecryptAES(key []byte, ct string) []byte {
	ciphertext, _ := hex.DecodeString(ct)

	block, err := aes.NewCipher(key)
	check(err)

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return plaintext
}

// checks the error and panics
func check(err error) {
	if err != nil {
		panic(err)
	}
}
