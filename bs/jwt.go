package bs

import (
	"atms/models"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	PrivateKey []byte
	PublicKey  []byte
}

func NewJWT(privateKey, publicKey []byte) JWT {
	return JWT{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}

func GetPrivateKey() (*rsa.PrivateKey, error) {

	prKeyPath := "/Users/venkatachintapalli/cert/id_rsa"

	prKey, err := ioutil.ReadFile(prKeyPath)
	if err != nil {
		panic("error opening path")
	}

	block, _ := pem.Decode(prKey)
	if block == nil {
		panic("could not decode auth key")
	}

	fmt.Println("PEM ", block.Type)

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic("error parsing private key")
	}

	return privateKey, nil
}

func GetPublicKey() (*rsa.PublicKey, error) {

	pubKeyPath := "/Users/venkatachintapalli/cert/id_rsa.pub"

	prKey, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		panic("error opening path")
	}

	block, _ := pem.Decode(prKey)
	if block == nil {
		panic("could not decode auth key")
	}

	fmt.Println("PEM ", block.Type)

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("error parsing private key")
	}

	return publicKey.(*rsa.PublicKey), nil
}

func ValidateIDToken(idTokenHeader string) (*models.IDTokenClaims, bool, error) {

	validToken := true

	token, err := jwt.ParseWithClaims(idTokenHeader, &models.IDTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		key, err := GetPublicKey()
		if err != nil {
			panic("error accessing public key info")
		}

		return key, nil
	})
	if !token.Valid {
		panic("expired token")
	}
	if err != nil {
		if validation, ok := err.(*jwt.ValidationError); ok && validation.Errors&jwt.ValidationErrorExpired != 0 {
			validToken = false
		} else {
			panic("error validating token")
		}
	}

	idtokenClaims, ok := token.Claims.(*models.IDTokenClaims)
	if !ok {
		panic("error getting token claims")
	}

	return idtokenClaims, validToken, nil
}

func Authenticate(r *http.Request) error {
	idTokenHeader := r.Header.Get(models.IDTokenHeader)
	if idTokenHeader == "" {
		panic("header not provided")
	}
	checkToken, _, tknErr := ValidateIDToken(idTokenHeader)
	if tknErr != nil || checkToken == nil {
		panic("error verifying id token")
	}

	return nil
}

// func (jwts JWT) Create(token string) (string, error) {
// 	key, err := jwt.ParseRSAPrivateKeyFromPEM(jwts.PrivateKey)
// 	if err != nil {
// 		return "", fmt.Errorf("create: parse key: %w", err)
// 	}

// 	claims := make(jwt.MapClaims)
// 	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
// 	claims["iat"] = time.Now().Unix()
// }
