package handle

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func Gentoken(human NewPersonData) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"person": human,
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)

	fmt.Println(tokenString, err)
	return tokenString
}

func GenerateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// func ReciveToken(tokenString string) {
// 	// sample token string taken from the New example
// 	tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

// 	// Parse takes the token string and a function for looking up the key. The latter is especially
// 	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
// 	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
// 	// to the callback, providing flexibility.
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")

// 		return hmacSampleSecret, nil
// 	},
// 		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok {
// 		fmt.Println(claims["foo"], claims["nbf"])
// 	} else {
// 		fmt.Println(err)
// 	}
// }

func ReciveToken(tokenString string) {

	// อย่า override tokenString ด้วยค่าใหม่
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// ตรวจสอบว่า alg ถูกต้องหรือไม่
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		log.Println("Error parsing token:", err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Name:", claims["name"])
		fmt.Println("Age:", claims["age"])
		fmt.Println("NBF:", claims["nbf"])

	} else {
		fmt.Println("Invalid token or claims.")
	}
}

func ReciveTokenNew(tokenString string) NewPersonData {
	var result NewPersonData

	// ใช้ parser แบบ v5
	parser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	token, err := parser.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		log.Println("Error parsing token:", err)
		return result
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		personData, ok := claims["person"].(map[string]interface{})
		if !ok {
			log.Println("Invalid person data in token")
			return result
		}

		// แปลง map[string]interface{} → JSON → struct
		jsonBytes, err := json.Marshal(personData)
		if err != nil {
			log.Println("Error marshalling person data:", err)
			return result
		}

		if err := json.Unmarshal(jsonBytes, &result); err != nil {
			log.Println("Error unmarshalling to struct:", err)
		} else {
			fmt.Println("Decoded Person:", result)
		}
	} else {
		log.Println("Invalid token or claims")
	}

	return result
}
