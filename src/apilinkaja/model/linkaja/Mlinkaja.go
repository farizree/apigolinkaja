package Mlinkaja

type (
	Account struct {
		ID             int
		AccountNumber  int
		CustomerNumber int
		Balance        int
		DTM_CRT string
		DTM_UPD string
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
		AccountID int32 `json:"account_id"`
		AccountNumber int32  `json:"account_number"`
		CustomerName  string `json:"customer_name"`
		Balance       int32  `json:"balance"`
		CustomerNumber int32 `json:"customer_number"`
	}

	TransactionDetail struct {
		AccountID int `json:"account_id"`
		CustomerNumber int `json:"customer_number"`
		TrxDebit int `json:"trx_debit"`
		TrxCredit int `json:"trx_credit"`
		TotalBalance int `json:"total_balance"`
		UserCreate string `json:"user_create"`
	}
)
