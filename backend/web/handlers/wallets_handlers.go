package handlers

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/i-invest/internal/wallets/b3"
)

type B3ImportForm struct {
	ReportFile *multipart.FileHeader `form:"report_file" binding:"required"`
}

func CreateWallet(c *gin.Context)            {}
func ShowWallet(c *gin.Context)              {}
func CreateWalletTransaction(c *gin.Context) {}

func ImportB3FileFormHandler(c *gin.Context) (string, string) {
	walletId, _ := c.Params.Get("id")

	var importFrom B3ImportForm
	err := c.ShouldBind(&importFrom)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return "", ""
	}

	return walletId, importFrom.ReportFile.Filename
}

func ImportWalletFromB3(c *gin.Context) {
	_, reportFilePath := ImportB3FileFormHandler(c)
	report, err := b3.ParseSummaryReport(reportFilePath)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": report,
	})
}

func ImportWalletTransactionsFromB3(c *gin.Context) {
	_, reportFilePath := ImportB3FileFormHandler(c)
	report, err := b3.ParseTransactionsReport(reportFilePath)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": report,
	})
}
