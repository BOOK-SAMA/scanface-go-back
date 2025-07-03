package handle

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func Testrsa() {
	// The GenerateKey method takes in a reader that returns random bits, and
	// the number of bits
	/*
		สร้าง RSA Key Pair ความยาว 2048 bits
		privateKey คือกุญแจลับ
		publicKey คือกุญแจสาธารณะสำหรับแจกจ่ายให้ผู้อื่น
	*/
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// The public key is a part of the *rsa.PrivateKey struct
	publicKey := privateKey.PublicKey

	// use the public and private keys
	// ...
	// https://play.golang.org/p/tldFUt2c4nx
	/*
		พิมพ์ค่าของ modulus N, exponent D ของ private key และ exponent E ของ public key
		แปลงเป็น base64 เพื่อให้ง่ายในการส่งข้ามระบบ เช่น JSON, API
	*/
	modulusBytes := base64.StdEncoding.EncodeToString(privateKey.N.Bytes())
	privateExponentBytes := base64.StdEncoding.EncodeToString(privateKey.D.Bytes())
	fmt.Println(modulusBytes)
	fmt.Println(privateExponentBytes)
	fmt.Println(publicKey.E)

	/*
		ใช้ EncryptOAEP เข้ารหัสด้วย public key
		ใช้ SHA256 ในการ padding ข้อมูลก่อนเข้ารหัส
		ใช้สำหรับส่งข้อมูลลับไปให้ผู้ที่ถือ private key ถอดได้เท่านั้น
	*/
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		[]byte("super secret message"),
		nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("encrypted bytes: ", encryptedBytes)

	// The first argument is an optional random data generator (the rand.Reader we used before)
	// we can set this value as nil
	// The OEAPOptions in the end signify that we encrypted the data using OEAP, and that we used
	// SHA256 to hash the input.
	/*
		ถอดรหัสข้อมูลที่ถูกเข้ารหัสด้วย public key
		ต้องใช้ private key เท่านั้นถึงจะถอดได้
	*/
	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}

	// We get back the original information in the form of bytes, which we
	// the cast to a string and print
	fmt.Println("decrypted message: ", string(decryptedBytes))

	/*
		สร้าง hash ของข้อความ msg แล้วใช้ private key ลงลายเซ็น
		ใช้ SignPSS ซึ่งเป็นมาตรฐาน RSA-PSS ที่ปลอดภัยกว่า PKCS1
	*/

	msg := []byte("verifiable message")

	// Before signing, we need to hash our message
	// The hash is what we actually sign
	msgHash := sha256.New()
	_, err = msgHash.Write(msg)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)

	// In order to generate the signature, we provide a random number generator,
	// our private key, the hashing algorithm that we used, and the hash sum
	// of our message
	/*
		ใช้ public key ในการตรวจสอบว่า signature ถูกต้องหรือไม่
		ถ้าลายเซ็นถูกต้อง จะไม่มี error
	*/
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}

	// To verify the signature, we provide the public key, the hashing algorithm
	// the hash sum of our message and the signature we generated previously
	// there is an optional "options" parameter which can omit for now
	err = rsa.VerifyPSS(&publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		fmt.Println("could not verify signature: ", err)
		return
	}
	// If we don't get any error from the `VerifyPSS` method, that means our
	// signature is valid
	fmt.Println("signature verified")
}
