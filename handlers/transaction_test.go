package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"wallet-api-go-bc/models"
	"wallet-api-go-bc/store"

	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAddTransaction(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		wallet         *models.Wallet
		expectedStatus int
		expectedError  string
	}{
		{"Valid Credit", `{"type": "credit", "amount": 50}`, &models.Wallet{ID: "1", Balance: 100}, http.StatusCreated, ""},
		{"Valid Debit", `{"type": "debit", "amount": 50}`, &models.Wallet{ID: "1", Balance: 100}, http.StatusCreated, ""},
		{"Insufficient Balance", `{"type": "debit", "amount": 150}`, &models.Wallet{ID: "1", Balance: 100}, http.StatusBadRequest, "Insufficient balance"},
		{"Invalid JSON", `{"type":`, &models.Wallet{ID: "1", Balance: 100}, http.StatusBadRequest, "Invalid request body"},
		{"Validation Error", `{"type": "invalid", "amount": 50}`, &models.Wallet{ID: "1", Balance: 100}, http.StatusBadRequest, "Validation failed"},
		{"Wallet Not Found", `{"type": "credit", "amount": 50}`, nil, http.StatusNotFound, "Wallet not found"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			store.Wallets = make(map[string]*models.Wallet)
			if tt.wallet != nil {
				store.Wallets["1"] = tt.wallet
			}

			req := httptest.NewRequest(http.MethodPost, "/wallets/1/transactions", bytes.NewReader([]byte(tt.body)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("1")

			err := AddTransaction(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			if tt.expectedError != "" {
				var resp map[string]string
				err = json.Unmarshal(rec.Body.Bytes(), &resp)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedError, resp["error"])
			}
		})
	}
}

func TestListTransactions(t *testing.T) {
	e := echo.New()
	store.Wallets = make(map[string]*models.Wallet)

	wallet := &models.Wallet{
		ID:      "1",
		Name:    "Test Wallet",
		Balance: 100,
		Txns: []models.Transaction{
			{Type: "credit", Amount: 100},
			{Type: "debit", Amount: 50},
		},
	}
	store.Wallets["1"] = wallet

	req := httptest.NewRequest(http.MethodGet, "/wallets/1/transactions?limit=1&offset=1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ListTransactions(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var txns []models.Transaction
	err = json.Unmarshal(rec.Body.Bytes(), &txns)
	assert.NoError(t, err)
	assert.Len(t, txns, 1)
	assert.Equal(t, "debit", txns[0].Type)
	assert.Equal(t, 50.0, txns[0].Amount)
}
