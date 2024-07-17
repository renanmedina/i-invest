package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/renanmedina/investment-warlock/web/handlers"
)

func main() {
	router := gin.Default()
	initializeHandlers(router)
	startWebserver(router)
}

func initializeHandlers(router *gin.Engine) {
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"health": "up",
		})
	})

	authGroup := router.Group("/auth")
	authGroup.POST("/login", handlers.AuthenticateUser)
	authGroup.POST("/register", handlers.RegisterUser)

	apiGroup := router.Group("/api")
	apiGroup.GET("/wallets/:id", handlers.ShowWallet)
	apiGroup.POST("/wallets", handlers.CreateWallet)
	apiGroup.POST("/wallets/:id/create-transaction", handlers.CreateWalletTransaction)
	apiGroup.POST("/wallets/:id/import-b3", handlers.ImportWalletFromB3)
	apiGroup.POST("/wallets/:id/import-b3-transactions", handlers.ImportWalletTransactionsFromB3)

	watchlistsGroup := apiGroup.Group("/watchlists")
	watchlistsGroup.POST("/import-b3", handlers.ImportWatchlistFromB3SummaryReport)
	watchlistsGroup.GET("/fetch-announcements", handlers.FetchNewAnnouncements)

	marketGroup := router.Group("/market")
	marketGroup.GET("/:tickerId/announcements", handlers.MarketTickerAnnouncements)
	marketGroup.GET("/:tickerId/fetch-new-announcements", handlers.FetchCompanyNewAnnouncements)
}

func startWebserver(router *gin.Engine) {
	err := router.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic("failed to start server at 0.0.0.0:8080")
	}

	fmt.Println("[WEBSERVER] Successfully listening at 0.0.0.0:8080")
}
