package cipher

import (
	aesCrypto "crypto/aes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAES(t *testing.T) {
	plainText := "Hello Shion! This is testtesttest!😂!!"
	// key must be 16, 24, 32 bytes
	key := "testing1testing1testing1testing1"
	invalidKey := "testing1testing1testing1testing"
	aes, err := NewAES(key)
	if err != nil {
		t.Error(err)
	}

	_, err = NewAES(invalidKey)
	assert.ErrorIs(t, err, aesCrypto.KeySizeError(len(invalidKey)))

	encryptedText := aes.Encrypt(plainText)
	decryptedText, err := aes.Decrypt(encryptedText)
	if err != nil {
		t.Error(err)
	}
	if plainText != decryptedText {
		t.Errorf("plainText: %s, decryptedText: %s", plainText, decryptedText)
	}
}
