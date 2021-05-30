package linkaja

import (
	"fmt"

	"github.com/alexcesaro/log/stdlog"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"

	Conf "apigolinkaja/src/config"

	logger "github.com/sirupsen/logrus"

	Mlinkaja "apigolinkaja/src/apilinkaja/model/linkaja"
)

func PostTransfer(c *gin.Context) {
	logkoe := stdlog.GetFromFlags()

	env, errenv := Conf.Environment()
	if errenv != nil {
		logger.Println(errenv)
		logkoe.Info(errenv)
	} else {
		if env == "production" {
			gin.SetMode(gin.ReleaseMode)
			// router := gin.New()
		} else if env == "development" {
			gin.SetMode(gin.DebugMode)
		}
	}
	dblinkaja := Conf.Init()
	var resultAccountFrom Mlinkaja.Account
	var resultAccountTo Mlinkaja.Account
	var transfer Mlinkaja.TransferAccount
	c.BindJSON(&transfer)
	var fromAccountNumber int = transfer.FromAccountNumber
	var toAccountNumber int = transfer.ToAccountNumber
	var amount int = transfer.Amount

	var dataFromAccountNumber = dblinkaja.QueryRow("select * from accounts where account_number = ?", fromAccountNumber).Scan(&resultAccountFrom.ID, &resultAccountFrom.AccountNumber, &resultAccountFrom.CustomerNumber, &resultAccountFrom.Balance)
	if dataFromAccountNumber != nil {
		fmt.Println("Data Account Number From Tidak Ditermukan")
		c.JSON(404, gin.H{
			"dataAccountFrom": "Data Account Number From Tidak Ditermukan",
		})
	}
	var dataToAccountNumber = dblinkaja.QueryRow("select * from accounts where account_number = ?", toAccountNumber).Scan(&resultAccountTo.ID, &resultAccountTo.AccountNumber, &resultAccountTo.CustomerNumber, &resultAccountTo.Balance)
	if dataToAccountNumber != nil {
		fmt.Println("Data Account Number To Tidak Ditermukan")
		c.JSON(404, gin.H{
			"dataAccountTo": "Data Account Number To Tidak Ditermukan",
		})
	}

	var getFromBalance int = resultAccountFrom.Balance
	var getToBalance int = resultAccountTo.Balance
	var dbFromAccountNumber int = resultAccountFrom.AccountNumber
	var dbToAccountNumber int = resultAccountTo.AccountNumber

	if fromAccountNumber == dbFromAccountNumber && toAccountNumber == dbToAccountNumber {
		if getFromBalance > amount {
			var totalBalanceFrom = getFromBalance - amount
			var totalBalanceTo = getToBalance + amount
			_, updateBalanceFrom := dblinkaja.Query("update accounts set balance = ? where account_number = ?", totalBalanceFrom, fromAccountNumber)
			if updateBalanceFrom != nil {
				fmt.Println("failed update balance account number from")
			}
			_, updateBalanceTo := dblinkaja.Query("update accounts set balance = ? where account_number = ?", totalBalanceTo, toAccountNumber)
			if updateBalanceTo != nil {
				fmt.Println("failed update balance account number to")
			}
			c.Writer.WriteHeader(201)
		} else {
			fmt.Println("Balance Tidak Mencukupi")
			c.JSON(404, gin.H{
				"Message": "Balance Tidak Mencukupi",
			})
		}
	} else {
		fmt.Println("Data doesn't match")
		c.JSON(404, gin.H{
			"Message": "Data doesn't match",
		})
	}
}
