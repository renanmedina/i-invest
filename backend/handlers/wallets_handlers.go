package handlers

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/investment-warlock/internal/wallets/b3"
)

type B3ImportForm struct {
	ReportFile *multipart.FileHeader `form:"report_file" binding:"required"`
}

func CreateWallet(c *gin.Context)            {}
func ShowWallet(c *gin.Context)              {}
func CreateWalletTransaction(c *gin.Context) {}

func ImportWalletFromB3(c *gin.Context) {
	// walletId, _ := c.Params.Get("id")

	// c.JSON(http.StatusOK, gin.H{
	// 	"wallet_id": walletId,
	// })

	var importFrom B3ImportForm
	err := c.ShouldBind(&importFrom)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	report, err := b3.ParseB3SummaryReport(importFrom.ReportFile.Filename)
	// wallet.Id = walletId
	// wallets.SaveWallet(wallet)

	c.JSON(http.StatusOK, gin.H{
		"items": report,
	})
}

func ImportWalletTransactionsFromB3(c *gin.Context) {}
