package bs

import (
	"atms/ds"
	"atms/models"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// func RegisterLoginHandlers() {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/api/login", userLoginHandler).Methods(http.MethodGet)
// 	log.Fatal(http.ListenAndServe("localhost:1235", router))
// }

func userLoginHandler(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	if !ok {
		panic("credentials not provided")
	}

	authenticate := models.Authenticate{
		UserName: username,
		Password: password,
	}

	userID, err := ds.AuthenticateUserAccount(dbds, authenticate)
	if err != nil || !userID {
		panic("could not identify user")
	}

	idTokenClaims := models.IDTokenClaims{
		UserID: authenticate.UserName,
	}

	key, err := GetPrivateKey()
	if err != nil {
		panic("couldn't get private key info")
	}

	idtoken := jwt.NewWithClaims(jwt.SigningMethodRS256, idTokenClaims)

	idTokenSigned, err := idtoken.SignedString(key)

	w.Header().Set("id-token", idTokenSigned)

	fmt.Println("authenticate: ", authenticate)
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
