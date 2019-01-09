package auth

import (
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/oklog/ulid"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

func createJWT() string {
	bytes, err := ioutil.ReadFile("./keys/sample_key.priv")
	if err != nil {
		log.Fatal("Error reading private key")

	}
	rsaPrivate, keyErr := crypto.ParseRSAPrivateKeyFromPEM(bytes)
	if keyErr != nil {
		log.Fatal("Error parsing private key")
	}

	claims := jws.Claims{}
	claims.Set("AccessToken", "level1")
	claims.Set("rollen", "admin, fibu")
	claims.SetIssuer("veith")
	claims.SetSubject("task")
	now := time.Now()
	claims.SetIssuedAt(now)
	claims.SetExpiration(now.Add(expiresDuration))
	claims.SetNotBefore(now)
	claims.SetJWTID(GenerateUlidString())
	jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)
	byteToken, err := jwt.Serialize(rsaPrivate)

	if err != nil {
		log.Fatal("Error signing the key. ", err)
		os.Exit(1)
	}

	return string(byteToken)
}

// Erzeuge eine ULID
func GenerateUlidString() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	newID, _ := ulid.New(ulid.Timestamp(t), entropy)
	return newID.String()
}
