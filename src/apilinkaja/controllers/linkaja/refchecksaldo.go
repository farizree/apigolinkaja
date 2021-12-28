package linkaja

import (
	"net/http"

	"github.com/alexcesaro/log/stdlog"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"

	Conf "apigolinkaja/src/config"

	logger "github.com/sirupsen/logrus"

	Mlinkaja "apigolinkaja/src/apilinkaja/model/linkaja"
)

func GetCheckSaldo(c *gin.Context) {
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
	var txt Mlinkaja.FindAccount
	c.BindJSON(&txt)
	var accountNumber int = txt.Accountnumber
	var resultAccount Mlinkaja.CustomerAccount

	var stmt = dblinkaja.QueryRow("select * from view_customeraccount where account_number = ?", accountNumber).Scan(&resultAccount.AccountID, &resultAccount.AccountNumber, &resultAccount.CustomerName, &resultAccount.Balance, &resultAccount.CustomerNumber)

	if stmt != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": resultAccount,
		})
	}
}
