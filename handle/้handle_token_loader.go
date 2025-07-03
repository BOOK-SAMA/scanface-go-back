package handle

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	keyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil || block.Type != "PRIVATE KEY" && block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		// Try PKCS8 fallback
		pkcs8Key, err2 := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err2 != nil {
			return nil, err
		}
		return pkcs8Key.(*rsa.PrivateKey), nil
	}

	return privateKey, nil
}

func LoadPublicKey(path string) (*rsa.PublicKey, error) {
	keyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	pubKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := pubKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not RSA public key")
	}

	return publicKey, nil
}

func Loadinit() {
	var (
		privateKey *rsa.PrivateKey
		publicKey  *rsa.PublicKey
	)

	var err error
	privateKey, err = LoadPrivateKey("private.pem")
	if err != nil {
		panic("Error loading private key: " + err.Error())
	}

	publicKey, err = LoadPublicKey("public.pem")
	if err != nil {
		panic("Error loading public key: " + err.Error())
	}

	fmt.Println(privateKey)
	fmt.Println(publicKey)
}
