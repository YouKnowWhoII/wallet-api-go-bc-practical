package store

import (
	"sync"
	"wallet-api-go-bc/models"
)

var (
	Wallets = make(map[string]*models.Wallet)
	Lock    sync.Mutex
)
