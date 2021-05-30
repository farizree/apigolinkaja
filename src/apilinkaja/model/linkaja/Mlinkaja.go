package Mlinkaja

type (
	Account struct {
		ID             int
		AccountNumber  int
		CustomerNumber int
		Balance        int
	}

	FindAccount struct {
		Accountnumber int `json:"account_number"`
	}

	TransferAccount struct {
		FromAccountNumber int `json:"from_account_number"`
		ToAccountNumber   int `json:"to_account_number"`
		Amount            int `json:"amount"`
	}

	CustomerAccount struct {
		AccountNumber int32  `json:"account_number"`
		CustomerName  string `json:"customer_name"`
		Balance       int32  `json:"balance"`
	}
)
