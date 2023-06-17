package authspec

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SetPubKeyRequest struct {
	PublicKey  string `json:"public_key"`
	CustomerId int64  `json:"customerid"`
	Token      string `json:"token"`
}

type GenericRequestWithToken struct {
	Token string `json:"token"`
}

type VerifyJwtReq struct {
	TokenString     string `json:"token_string"`
	PublicKeyString string `json:"pub_key_string"`
}

type CreateUserRequest struct {
	CustomerId int64  `json:"customerid"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}
