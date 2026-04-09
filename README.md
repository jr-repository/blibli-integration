# Blibli Integration API (Go)

A professional, production-ready Go template for integrating with the Blibli Seller API. This repository handles authentication, signature generation, and basic API communication, specifically demonstrating how to fetch a product list.

This template is built with security in mind, ensuring no credentials are hardcoded. It is designed to be easily extensible for other Blibli API endpoints.

---

## 🚀 Current Integration Step

* **Phase 1 Completed:** API Client Initialization, Authentication/Signature Generation, and "Hit List Product" (Fetch Products) endpoint testing.

## 📂 Project Architecture

```text
blibli-integration/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point & test execution
├── internal/
│   ├── blibli/
│   │   ├── client.go            # Core HTTP client, signature generation, headers
│   │   └── product.go           # Product-related API implementations
│   ├── config/
│   │   └── config.go            # Environment variable loading & validation
│   └── models/
│       └── product.go           # Structs for request/response payloads
├── .env.example                 # Template for environment variables
├── go.mod                       # Go module dependencies
└── README.md                    # Documentation
```

## ⚙️ Setup & Installation

**1. Clone the repository:**
```bash
git clone https://github.com/jr-repository/blibli-integration.git
cd blibli-integration
```

**2. Install dependencies:**
```bash
go mod tidy
```

**3. Configure Environment:**
Copy the example environment file and fill in your actual Blibli Seller credentials.
```bash
cp .env.example .env
```
> **Note:** Never commit your `.env` file to version control.

**4. Run the application:**
```bash
go run cmd/api/main.go
```

## 🛠️ Development Guidelines

* **Payloads:** Add new API payloads in `internal/models/`.
* **Endpoints:** Implement new endpoint logic in `internal/blibli/` by extending the `Client` struct.
* **Best Practices:** Maintain clean error handling and avoid hardcoding paths.

---

## 📦 Dependencies (`go.mod`)

*(Reference for your `go.mod` file)*

```go
module github.com/jr-repository/blibli-integration

go 1.21

require github.com/joho/godotenv v1.5.1
```