package customerspec

type CreateCustomerRequest struct {
	CustomerId string `json:"customerid"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Age        int32  `json:"age"`
	Gender     string `json:"gender"`
	Password   string `json:"password"`
}

type DeleteCustomerRequest struct {
	CustomerId string `json:"customerid"`
}

type GetCustomerRequest struct {
	CustomerId string `json:"customerid"`
}

type UpdateCustomerRequest struct {
	CustomerId int64  `json:"customerid"`
	FirstName  string `json:"fisrtname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Password   string `json:"pass"`
	Age        string `json:"age"`
	Secret     string `json:"secret"`
	Gender     string `json:"gender"`
}

type GetOffersRequest struct {
	TokenString string `json:"auth_token"`
	PublicKey   string `json:"pub_key"`
	CustomerId  string `json:"customerid"`
}
