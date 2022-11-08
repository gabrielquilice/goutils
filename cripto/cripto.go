package critpo

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

type Key interface {
	[]byte | string
}

func EncryptFile[K Key](filename string, data []byte, key K) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Erro ao criptografar arquivo: %v\n", err)
	}
	defer f.Close()
	f.Write(encrypt(data, key))
}

func encrypt[K Key](data []byte, key K) []byte {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func DecryptFile[K Key](filename string, key K) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Erro ao efetuar o decryptfile: %v\n", err)
	}
	return decrypt(data, key)
}

func decrypt[K Key](data []byte, key K) []byte {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	return plaintext
}
