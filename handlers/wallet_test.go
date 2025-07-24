package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"wallet-api-go-bc/models"
	"wallet-api-go-bc/store"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateWallet(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		expectedStatus int
		expectedError  string
	}{
		{"Valid Request", `{"name": "Test Wallet"}`, http.StatusCreated, ""},
		{"Invalid JSON", `{"name":`, http.StatusBadRequest, "Invalid request body"},
		{"Missing Name", `{}`, http.StatusBadRequest, "Validation failed"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			store.Wallets = make(map[string]*models.Wallet)

			req := httptest.NewRequest(http.MethodPost, "/wallets", bytes.NewReader([]byte(tt.body)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := CreateWallet(c)
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

func TestGetWallet(t *testing.T) {
	e := echo.New()
	store.Wallets = make(map[string]*models.Wallet)

	wallet := &models.Wallet{ID: "1", Name: "Test Wallet", Balance: 100}
	store.Wallets["1"] = wallet

	req := httptest.NewRequest(http.MethodGet, "/wallets/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := GetWallet(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var result models.Wallet
	err = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, wallet, &result)
}
