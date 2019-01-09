package middleware

import (
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Methode um den Token zu prüfen und bei Erfolg in den context zu legen. Das Abfragen der userdb
// und erstellen des Tokens passiert in package internal/auth
var JWTAuthFunc = func(ctx context.Context) (context.Context, error) {
	token, err := authTokenFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	tokenInfo, err := parseToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}
	newCtx := context.WithValue(ctx, "tokenInfo", tokenInfo)
	return newCtx, nil
}

func parseToken(token string) (map[string]interface{}, error) {
	pubKey, err := ioutil.ReadFile("./keys/sample_key.pub")
	if err != nil {
		log.Fatal("Error reading public key")
		os.Exit(1)
	}
	rsaPublic, _ := crypto.ParseRSAPublicKeyFromPEM(pubKey)
	jwt, err := jws.ParseJWT([]byte(token))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "not parseable auth token: %v", err)

	}
	// Validate token
	if err = jwt.Validate(rsaPublic, crypto.SigningMethodRS256); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}
	return jwt.Claims(), nil
}

func authTokenFromMD(ctx context.Context, expectedScheme string) (string, error) {
	var val string

	cookie := metautils.ExtractIncoming(ctx).Get("cookie")

	//alle cookies
	cookies := strings.Split(cookie, "; ")
	for _, element := range cookies {
		if len(element) > 0 && element[0:13] == "Authorization" {
			val = element[14:len(element)]
			break
		}
	}

	if val == "" {
		val = metautils.ExtractIncoming(ctx).Get("authorization")
	}

	if val == "" {

		return "", status.Errorf(codes.Unauthenticated, "Request unauthenticated with "+expectedScheme)

	}
	splits := strings.SplitN(val, " ", 2)
	if len(splits) < 2 {
		return "", status.Errorf(codes.Unauthenticated, "Bad authorization string")
	}
	if strings.ToLower(splits[0]) != strings.ToLower(expectedScheme) {
		return "", status.Errorf(codes.Unauthenticated, "Request unauthenticated with "+expectedScheme)
	}
	return splits[1], nil
}
