package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// const container = `
// {
// 	"container":{
// 		"first":"Janet",
// 		"last":"Prichard"
// 	},
// 	"age":47
// }`

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

func main() {

	// dosmth()
	s := "loooooollll"
	buf := bytes.NewBufferString(s)
	// publicKey := buf.String()

	resp, err := http.Post("http://localhost:8090/storage/lol", "text/plain", buf)

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		os.Exit(1)
	}
	publicKey := &privateKey.PublicKey

	e, err := encryptMessage(publicKey, []byte("KEK!"))
	d, err := decryptMessage(e, privateKey)

	fmt.Println("debug decrypted: ", string(d))
	pubkeyJSON, err := json.Marshal(*publicKey)

	req, err := http.NewRequest("GET", "http://localhost:8090/storage/lol", bytes.NewBuffer(pubkeyJSON))
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	decryptStr, err := decryptMessage(body, privateKey)

	// We get back the original information in the form of bytes, which we
	// the cast to a string and print
	fmt.Println("decrypted message: ", string(decryptStr))
}
