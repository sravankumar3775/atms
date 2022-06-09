package bs

import (
	"atms/ds"
	"atms/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

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
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	key, err := GetPrivateKey()
	if err != nil {
		panic("couldn't get private key info")
	}

	idtoken := jwt.NewWithClaims(jwt.SigningMethodRS256, idTokenClaims)

	idTokenSigned, err := idtoken.SignedString(key)
	if err != nil {
		panic("couldn't get private key info")
	}

	w.Header().Set(models.IDTokenHeader, idTokenSigned)
}
