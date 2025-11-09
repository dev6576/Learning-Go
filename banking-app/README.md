# Banking Application

A simple banking application built with Go that demonstrates basic banking operations using REST APIs.

## Features

- Create bank accounts
- Check account balance
- Deposit money
- Transfer money between accounts

## API Endpoints

### Create Account
```bash
curl -X POST http://localhost:8080/accounts -H "Content-Type: application/json" -d "{\"id\":\"alice\"}"
```

### Check Balance
```bash
curl "http://localhost:8080/balance?id=alice"
```

### Deposit Money
```bash
curl -X POST http://localhost:8080/deposit -H "Content-Type: application/json" -d "{\"id\":\"alice\",\"amount\":100}"
```

### Transfer Money
```bash
curl -X POST http://localhost:8080/transfer -H "Content-Type: application/json" -d "{\"from\":\"alice\",\"to\":\"bob\",\"amount\":50}"
```

## Test Sequence

1. Create two accounts:
```bash
# Create first account
curl -X POST http://localhost:8080/accounts -H "Content-Type: application/json" -d "{\"id\":\"alice\"}"

# Create second account
curl -X POST http://localhost:8080/accounts -H "Content-Type: application/json" -d "{\"id\":\"bob\"}"
```

2. Deposit money into alice's account:
```bash
curl -X POST http://localhost:8080/deposit -H "Content-Type: application/json" -d "{\"id\":\"alice\",\"amount\":100}"
```

3. Check alice's balance:
```bash
curl "http://localhost:8080/balance?id=alice"
```

4. Transfer money from alice to bob:
```bash
curl -X POST http://localhost:8080/transfer -H "Content-Type: application/json" -d "{\"from\":\"alice\",\"to\":\"bob\",\"amount\":50}"
```

5. Check both balances:
```bash
# Check alice's balance
curl "http://localhost:8080/balance?id=alice"

# Check bob's balance
curl "http://localhost:8080/balance?id=bob"
```

## Project Structure

```
banking-app/
├── cmd/
│   └── server/
│       └── main.go         # Main application entry point
├── internal/
│   ├── api/
│   │   ├── handlers.go     # HTTP handlers
│   │   └── router.go       # Router setup
│   └── bank/
│       ├── account.go      # Account operations
│       ├── bank.go         # Bank operations
│       └── errors.go       # Error definitions
└── pkg/                    # Public packages (if any)
```

## Running the Application

From the project root:
```bash
go run ./cmd/server/main.go
```