package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"wallet-api-go-bc/models"
	"wallet-api-go-bc/store"
)

var validate = validator.New()

func AddTransaction(c echo.Context) error {
	id := c.Param("id")

	var body models.Transaction
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	if err := validate.Struct(body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Validation failed"})
	}

	store.Lock.Lock()
	defer store.Lock.Unlock()

	wallet, ok := store.Wallets[id]
	if !ok {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Wallet not found"})
	}

	if body.Type == "debit" && wallet.Balance < body.Amount {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Insufficient balance"})
	}

	if body.Type == "credit" {
		wallet.Balance += body.Amount
	} else {
		wallet.Balance -= body.Amount
	}

	txn := models.Transaction{
		Type:   body.Type,
		Amount: body.Amount,
	}
	wallet.Txns = append(wallet.Txns, txn)

	return c.JSON(http.StatusCreated, txn)
}

func ListTransactions(c echo.Context) error {
	id := c.Param("id")

	store.Lock.Lock()
	wallet, ok := store.Wallets[id]
	store.Lock.Unlock()

	if !ok {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Wallet not found"})
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	end := offset + limit
	if end > len(wallet.Txns) {
		end = len(wallet.Txns)
	}

	return c.JSON(http.StatusOK, wallet.Txns[offset:end])
}
