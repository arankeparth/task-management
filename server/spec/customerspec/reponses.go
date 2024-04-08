package customerspec

type GetCustomerResponse struct {
	CustomerId string `json:"customerid" db:"CustomerId"`
	FirstName  string `json:"fisrtname" db:"FirstName"`
	LastName   string `json:"lastname" db:"LastName"`
	Email      string `json:"email" db:"Email"`
	Age        string `json:"age" db:"Age"`
	Gender     string `json:"gender" db:"Gender"`
}

type CreateCustomerRsponse struct {
	Status     string `json:"status"`
	CustomerId int64  `json:"customerid"`
}

type DeleteCustomerResponse struct {
	Status string `json:"status"`
}

type GetOffersResponse struct {
	CustomerId int64  `json:"customerid" sql:"customerid"`
	Title      string `json:"title" sql:"title"`
	Desc       string `json:"desc" sql:"description"`
}
