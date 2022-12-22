package cipher

import "testing"

func TestAES(t *testing.T) {
	plainText := "Hello Shion! This is testtesttest!😂!!"
	key := "testing1testing1testing1testing1"
	aes, err := NewAES(key)
	if err != nil {
		t.Error(err)
	}
	encryptedText := aes.Encrypt(plainText)
	decryptedText, err := aes.Decrypt(encryptedText)
	if err != nil {
		t.Error(err)
	}
	if plainText != decryptedText {
		t.Errorf("plainText: %s, decryptedText: %s", plainText, decryptedText)
	}
}
