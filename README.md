# Wallet API using Echo Framework - Practical Assignment Botcalm

Build a wallet API using Golang and the Echo framework. No external databases, use in-memory storage only.

## Requirements

- Use Echo (https://echo.labstack.com/)
- Properly structured code (handlers, models, etc.)
- Config Management (Load configs from .env (e.g. PORT). Use `github.com/joho/godotenv`)
- Use `map` or other in-memory storage
- Return meaningful status codes & error messages
- Must be clean and readable

## Features to Implement

### POST `/wallets`
- Create a new wallet with a name.
- Auto-generate a wallet ID and starting balance `0`.

### GET `/wallets/:id`
- Get wallet details including current balance.

### POST `/wallets/:id/transactions`
- Accept a JSON body with:
  ```json
  {
    "type": "credit", // or "debit"
    "amount": 100
  }
  ```
- Update the wallet's balance accordingly.
- Reject negative balances or insufficient funds.

### GET `/wallets/:id/transactions`
- Return list of all transactions for the given wallet and add pagination on transactions (Support `limit` and `offset` query params)

---

### Bonus Tasks
- Add basic middleware to log each request’s method and path.
- Rate Limiting Middleware (Limit number of API requests per minute per IP)
- Docker Support (Add `Dockerfile` and optionally a `docker-compose.yml` to run app)



---

## Testing (Optional Advanced Requirement)

Write **unit tests** for the following:

- Wallet creation logic (`POST /wallets`)
- Transaction logic (`POST /wallets/:id/transactions`)
- Validation errors (e.g., invalid JSON)

Use Go’s native `testing` package and organize test files as:

```
handlers/
├── wallet.go
├── wallet_test.go
```

> Use `go test ./...` to run all tests.

Bonus if:
- You cover edge cases
- Use table-driven tests
- Mock dependencies where needed

## API Endpoints

- `POST /wallets` - Create wallet
- `GET /wallets/:id` - Get wallet
- `POST /wallets/:id/transactions` - Add transaction
- `GET /wallets/:id/transactions` - List all transactions