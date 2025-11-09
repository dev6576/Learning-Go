# Banking-App

A simple Go backend for a banking application.

## Table of Contents

* Overview
* Architecture & Key Concepts
* API Endpoints
* Worker Pooling and Goroutines
* Why Transaction Request and Execution are Decoupled
* Setup & Running
* Future Improvements

## Overview

This project implements a REST API backend for basic banking operations including account creation, balance checking, deposits, withdrawals, and money transfers. It demonstrates clean separation of concerns: API routing, business logic, transaction management, and concurrency handling.

## Architecture & Key Concepts

* `internals/api/router.go`: Defines all HTTP routes and binds them to handler functions.
* `internals/api/handlers`: Handler functions parse requests, call the business logic, and return responses.
* `internals/service`: Handles core banking operations such as account management and transactions.
* `internals/repository`: Interfaces with the data store to manage account and transaction records.

Transactions are decoupled into **request** and **execution** phases:

* The request phase receives and validates the transaction.
* The execution phase applies the changes atomically to ensure consistency and proper logging.

## API Endpoints

| HTTP Method               | Path                  | Description                                                                              |
| ------------------------- | --------------------- | ---------------------------------------------------------------------------------------- |
| `POST /account`           | `/account`            | Create a new bank account.                                                               |
| `GET /balance`            | `/balance`            | Retrieve the current balance of an account.                                              |
| `POST /transfer`          | `/transfer`           | Submit a transfer request from one account to another. Handles validation and execution. |
| `POST /deposit`           | `/deposit`            | Deposit funds into an account.                                                           |
| `POST /withdraw`          | `/withdraw`           | Withdraw funds from an account.                                                          |
| `GET /transaction-status` | `/transaction-status` | Check the status of a transaction.                                                       |

## Worker Pooling and Goroutines

The application uses goroutines with worker pools to handle transactions concurrently. This design allows:

* Efficient processing of multiple simultaneous transactions.
* Non-blocking request handling to improve throughput.
* Controlled concurrency using a fixed number of worker goroutines.
* Thread-safe execution of transactions by synchronizing access to shared account data.

Transactions are sent to a channel, and worker goroutines pick them up to execute in a safe and atomic manner. This ensures scalability and reliability in high-load scenarios.

## Why Transaction Request and Execution are Decoupled

* **Validation & Authorization**: Ensure user permissions, account ownership, and sufficient balance before execution.
* **Atomicity**: Apply debit/credit and record transactions in a single operation.
* **Separation of Concerns**: Keeps HTTP handling separate from business logic.
* **Extensibility**: Easier to implement asynchronous processing, retries, logging, and audit trails.
* **Error Handling & Retries**: Centralized execution logic allows robust failure handling.
* **Audit & Traceability**: Track request intent separately from execution outcome for compliance.

## Setup & Running

1. Clone the repository:

```bash
git clone https://github.com/dev6576/Learning-Go.git
cd Learning-Go/banking-app
```

2. Configure environment variables (DB connection, JWT secrets, etc.)
3. Run database migrations for accounts and transactions tables.
4. Start the server:

```bash
go run main.go
```

5. API listens on the configured port (e.g., `:8080`). Test endpoints using Postman or curl.

## Future Improvements

* Add worker queue monitoring and dynamic scaling of goroutines.
* Implement more transaction types: scheduled transfers, recurring deposits.
* Enhance concurrency handling and account locking during transfers.
* Expand audit trail and logging features.
* Add pagination and filtering for transaction history.
* Introduce rate limiting and fraud detection mechanisms.
