package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

func deriveKey(passphrase string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha256.New), salt
}

/*
Encrypt encrypts a given string using AES-256-GCM encryption with a passphrase.

Usage:

	passphrase := "super secret password"
	value := "value to be encrypted"
	encrypted, err := encryption.Encrypt(passphrase, value)
*/
func Encrypt(passphrase, plaintext string) (encrypted string, err error) {
	key, salt := deriveKey(passphrase, nil)
	iv := make([]byte, 12)
	rand.Read(iv)
	b, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	aesgcm, err := cipher.NewGCM(b)
	if err != nil {
		return
	}
	value := []byte(plaintext)
	data := aesgcm.Seal(nil, iv, value, nil)
	encrypted = strings.Join([]string{
		hex.EncodeToString(salt),
		hex.EncodeToString(iv),
		hex.EncodeToString(data),
	},
		"-")
	return
}

/*
Decrypt decrypts a cipher that was encrypted using AES-256-GCM encryption.

Usage:

	passphrase := "super secret password"
	cipherText := "cipher value"
	decrypted, err := encryption.Decrypt(passphrase, cipherText)
*/
func Decrypt(passphrase, ciphertext string) (decrypted string, err error) {
	arr := strings.Split(ciphertext, "-")
	salt, err := hex.DecodeString(arr[0])
	if err != nil {
		return
	}
	iv, err := hex.DecodeString(arr[1])
	if err != nil {
		return
	}
	data, err := hex.DecodeString(arr[2])
	if err != nil {
		return
	}
	key, _ := deriveKey(passphrase, salt)
	b, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	aesgcm, err := cipher.NewGCM(b)
	if err != nil {
		return
	}
	data, err = aesgcm.Open(nil, iv, data, nil)
	if err != nil {
		return
	}
	decrypted = string(data)
	return
}
