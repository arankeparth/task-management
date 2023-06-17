package customerspec

type GetCustomerResponse struct {
	CustomerId int64  `json:"customerid"`
	FirstName  string `json:"fisrtname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Age        string `json:"age"`
	Gender     string `json:"gender"`
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
