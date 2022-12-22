package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
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

func (c AES) Encrypt(plainText string) string {
	plainByte := []byte(plainText)
	encryptedByte := make([]byte, aes.BlockSize+len(plainByte))
	iv := encryptedByte[:aes.BlockSize]
	encryptStream := cipher.NewCTR(c.block, iv)
	encryptStream.XORKeyStream(encryptedByte[aes.BlockSize:], plainByte)
	// convert to base64 using base64 package
	encryptedText := base64.StdEncoding.EncodeToString(encryptedByte)
	return encryptedText
}

func (c AES) Decrypt(encryptedText string) (string, error) {
	encryptedByte, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}
	decryptedText := make([]byte, len(encryptedByte[aes.BlockSize:]))
	decryptStream := cipher.NewCTR(c.block, encryptedByte[:aes.BlockSize])
	decryptStream.XORKeyStream(decryptedText, encryptedByte[aes.BlockSize:])
	return string(decryptedText), nil
}
