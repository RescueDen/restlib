package authentication

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

/**
File of static support functions for passwords creating, editing, hashing, etc.
*/

/*
JWT claims struct
*/
type Token struct {
	UserId int
	jwt.StandardClaims
}

/**
Support function to hash the password
*/
func HashPassword(password string) string {

	//Hash the password, there should be a salt
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)

}

/**
  Support function to generate a JWT token
*/
func CreateJWTToken(userId int) string {

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: userId}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

	return tokenString

}

/**
  Compare passwords.  Determine if they match
*/
func ComparePasswords(currentPwHash string, testingPassword string) bool {

	//Now take the password and encrypt it
	err := bcrypt.CompareHashAndPassword([]byte(currentPwHash), []byte(testingPassword))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return false
	} else {
		return true
	}

}

/**
  Compare passwords.  Determine if they match
*/
func ValidateToken(tokenHeader string) (int, error) {

	//Token is missing, returns with error code 403 Unauthorized
	if tokenHeader == "" {
		return -1, errors.New("auth_missing_token")
	}

	//Now split the token to get the useful part
	splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	if len(splitted) != 2 {
		return -1, errors.New("auth_malformed_token")

	}

	//Grab the token part, what we are truly interested in
	tokenPart := splitted[1]

	//Get the token and take it back apart
	tk := &Token{}

	//Now parse the token
	token, err := jwt.ParseWithClaims(tokenPart, tk,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

	//check for mailformed data
	if err != nil { //Malformed token, returns with http code 403 as usual
		return -1, errors.New("auth_malformed_token")

	}

	//Token is invalid, maybe not signed on this server
	if !token.Valid {
		//Return the error
		return -1, errors.New("auth_forbidden")

	}

	return 1, nil

}
