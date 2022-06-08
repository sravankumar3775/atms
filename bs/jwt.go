package bs

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

// func (jwts JWT) Create(token string) (string, error) {
// 	key, err := jwt.ParseRSAPrivateKeyFromPEM(jwts.PrivateKey)
// 	if err != nil {
// 		return "", fmt.Errorf("create: parse key: %w", err)
// 	}

// 	claims := make(jwt.MapClaims)
// 	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
// 	claims["iat"] = time.Now().Unix()
// }
