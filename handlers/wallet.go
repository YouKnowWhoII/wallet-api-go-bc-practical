package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"wallet-api-go-bc/models"
	"wallet-api-go-bc/store"
)

func CreateWallet(c echo.Context) error {

	var body models.Wallet
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	id := uuid.NewString()
	wallet := &models.Wallet{
		ID:      id,
		Name:    body.Name,
		Balance: 0,
		Txns:    []models.Transaction{},
	}

	if err := validate.Struct(wallet); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Validation failed"})
	}

	store.Lock.Lock()
	store.Wallets[id] = wallet
	store.Lock.Unlock()

	return c.JSON(http.StatusCreated, wallet)
}

func GetWallet(c echo.Context) error {
	id := c.Param("id")
	store.Lock.Lock()
	wallet, ok := store.Wallets[id]
	store.Lock.Unlock()

	if !ok {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Wallet not found"})
	}

	return c.JSON(http.StatusOK, wallet)
}
