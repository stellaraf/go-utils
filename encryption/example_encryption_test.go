package encryption_test

import (
	"fmt"

	"go.stellar.af/utils/encryption"
)

func Example() {
	passphrase := "super secret password"
	data := "value to encrypt"
	encrypted, err := encryption.Encrypt(passphrase, data)
	if err != nil {
		panic(err)
	}

	decrypted, err := encryption.Decrypt(passphrase, encrypted)
	if err != nil {
		panic(err)
	}
	fmt.Println(decrypted)
	// Output: value to encrypt
}
