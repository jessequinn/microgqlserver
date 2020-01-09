package services

import (
	"github.com/dgrijalva/jwt-go"
	pb "github.com/jessequinn/microgqlserver/srv/authsrv/proto/auth"
	rs "github.com/jessequinn/microgqlserver/srv/authsrv/repositories"
	"time"
)

var (
	// Define a secure key string used
	// as a salt when hashing our tokens.
	// Please make your own way more secure than this,
	// use a randomly generated md5 hash or something.
	// md5sum <<< "keypassphrase"
	key = []byte("871a7b7013cdcb64e2710e3e7cb4d1b9")
)

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	Repo rs.Repository
}

// Decode a token string into a token object
func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {
	// Parse the token
	tokenType, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	// Validate the token and return the custom claims
	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// Encode a claim into a JWT
func (srv *TokenService) Encode(user *pb.User) (string, error) {
	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			//ExpiresAt: 15000,
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    "go.micro.srv.user",
		},
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign token and return
	return token.SignedString(key)
}
