package router

import (
	"github.com/labstack/echo/v4"
	"wallet-api-go-bc/handlers"
)

func RegisterRoutes(e *echo.Echo) {
	wallets := e.Group("/wallets")
	wallets.POST("", handlers.CreateWallet)
	wallets.GET("/:id", handlers.GetWallet)
	wallets.POST("/:id/transactions", handlers.AddTransaction)
	wallets.GET("/:id/transactions", handlers.ListTransactions)
}
