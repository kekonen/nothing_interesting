package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"net/http"
	"os"

	"bytes"

	"github.com/gin-gonic/gin"
)

func decryptMessage(message []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	decryptedBytes, err := privateKey.Decrypt(nil, message, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return make([]byte, 0), err
	}

	return decryptedBytes, nil
}

func encryptMessage(pk *rsa.PublicKey, message []byte) ([]byte, error) {
	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		pk,
		message,
		nil,
	)
	return ciphertext, err
}

// https://www.sohamkamani.com/golang/rsa-encryption/
func main() {

	storage := make(map[string]string)

	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/storage/:path", func(c *gin.Context) {
		path := c.Param("path")
		println("GOT", path)

		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)
		publicKeyU8 := buf.Bytes()

		var publicKey rsa.PublicKey
		err := json.Unmarshal(publicKeyU8, &publicKey)
		if err != nil {
			os.Exit(1)
		}

		cipherText, err := encryptMessage(&publicKey, []byte(storage[path]))

		c.String(http.StatusOK, string(cipherText))
	})

	router.POST("/storage/:path", func(c *gin.Context) {
		path := c.Param("path")
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)
		data := buf.String()

		println("GOT", data)

		storage[path] = data

		c.String(http.StatusOK, "Hello %s", data)
	})

	router.Run(":8090")
}
