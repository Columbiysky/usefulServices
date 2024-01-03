package logic

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	tokenDb "sso/dbLogic/token"
	"strconv"
	"time"
)

func GenerateTokenForAccount(accountId int, accountLogin string) string {
	nowInUtcAsString := time.Now().UTC().String()

	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalln(err)
	}

	token := encrypt(key, strconv.Itoa(accountId)+accountLogin+nowInUtcAsString)
	return token
}

func RegisterActivity(tokenValue string) {
	tokenDb.RegisterActivity(tokenValue)
}

func RegisterToken(accountId int, tokenValue string) {
	tokenDb.RegisterToken(accountId, tokenValue)
}

// fully stolen from https://gist.github.com/manishtpatel/8222606#file-main-go-L27
// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}
