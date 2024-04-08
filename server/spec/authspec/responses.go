package authspec

import "github.com/dgrijalva/jwt-go"

type LoginResponse struct {
	IsLoggedIn   bool   `json:"is_loggedin"`
	ErrorMessage string `json:"error_message"`
	AuthToken    string `json:"auth_token"`
	PublicKey    string `json:"public_key"`
	CustomerId   string `json:"customerid"`
	CustomerName string `json:"customername"`
}

type SetPubKeyResponse struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message"`
}

type CreateUserResponse struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message"`
}

type VerifyJwtResponse struct {
	Status       int           `json:"status"`
	ErrorMessage string        `json:"error_message"`
	ParsedToken  jwt.MapClaims `json:"token"`
}
