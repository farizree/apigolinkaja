package linkaja

import (
	"fmt"
	"net/http"

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

	var dataFromAccountNumber = dblinkaja.QueryRow("select * from accounts where account_number = ?", fromAccountNumber).Scan(&resultAccountFrom.ID, &resultAccountFrom.AccountNumber, &resultAccountFrom.CustomerNumber, &resultAccountFrom.Balance, &resultAccountFrom.DTM_CRT, &resultAccountFrom.DTM_UPD)

	if dataFromAccountNumber != nil {
		fmt.Println("Data Account Number From Tidak Ditermukan")
		c.JSON(404, gin.H{
			"dataAccountFrom": "Data Account Number From Tidak Ditemukan",
		})
	} else {
		var dataToAccountNumber = dblinkaja.QueryRow("select * from accounts where account_number = ?", toAccountNumber).Scan(&resultAccountTo.ID, &resultAccountTo.AccountNumber, &resultAccountTo.CustomerNumber, &resultAccountTo.Balance, &resultAccountFrom.DTM_CRT, &resultAccountFrom.DTM_UPD)
		if dataToAccountNumber != nil {
			fmt.Println("Data Account Number To Tidak Ditermukan")
			c.JSON(404, gin.H{
				"dataAccountTo": "Data Account Number To Tidak Ditemukan",
			})
		}
	}

	var accountIDFrom int = resultAccountFrom.ID
	var customerNumberFrom int = resultAccountFrom.CustomerNumber
	var dbFromAccountNumber int = resultAccountFrom.AccountNumber
	var getFromBalance int = resultAccountFrom.Balance

	var accountIDTo int = resultAccountTo.ID
	var customerNumberTo int = resultAccountTo.CustomerNumber
	var dbToAccountNumber int = resultAccountTo.AccountNumber
	var getToBalance int = resultAccountTo.Balance

	var zeroValueDebit int = 0
	var zeroValueCredit int = 0

	var usrCrt string = "Farizree"

	if fromAccountNumber == dbFromAccountNumber && toAccountNumber == dbToAccountNumber {
		if getFromBalance > amount {
			var totalBalanceFrom = getFromBalance - amount
			var totalBalanceTo = getToBalance + amount

			if fromAccountNumber == toAccountNumber {
				c.JSON(404, gin.H{
					"Error": "Cannot transfer with the same account number",
				})
			} else {
				_, insertSPTRXDetailFrom := dblinkaja.Query("CALL spTransactionDetail (?,?,?,?,?,?)", accountIDFrom, customerNumberFrom, zeroValueDebit, amount, totalBalanceFrom, usrCrt)
				if insertSPTRXDetailFrom != nil {
					c.JSON(404, gin.H{
						"Error": "failed update balance account number from",
					})
				} else {

					_, insertSPTRXDetailTo := dblinkaja.Query("CALL spTransactionDetail (?,?,?,?,?,?)", accountIDTo, customerNumberTo, amount, zeroValueCredit, totalBalanceTo, usrCrt)
					if insertSPTRXDetailTo != nil {
						c.JSON(404, gin.H{
							"Error": "failed update balance account number to",
						})
					} else {
						c.Writer.WriteHeader(201)
						c.JSON(http.StatusOK, gin.H{
							"Message": "Yeay! Success Tranfer From Your Account",
						})
					}
				}
			}
		} else {
			c.JSON(404, gin.H{
				"Error": "Balance Tidak Mencukupi",
			})
		}
	} else {
		c.JSON(404, gin.H{
			"Error": "Data doesn't match",
		})
	}
}
