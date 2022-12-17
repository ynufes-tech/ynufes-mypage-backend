package cipher

import (
	"crypto/aes"
	"crypto/cipher"
)

type AES struct {
	key   string
	block cipher.Block
}

func NewAES(key string) (*AES, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	return &AES{
		key:   key,
		block: block,
	}, nil
}

func (aes AES) Encrypt(plainText string) string {
	encryptedText := make([]byte, len(plainText))
	aes.block.Encrypt(encryptedText, []byte(plainText))
	return string(encryptedText)
}

func (aes AES) Decrypt(encryptedText string) (string, error) {
	decryptedText := make([]byte, len(encryptedText))
	aes.block.Decrypt(decryptedText, []byte(encryptedText))
	return string(decryptedText), nil
}
