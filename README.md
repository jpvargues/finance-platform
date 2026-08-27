# Finance Platform

A backend for analyzing ETF diversification — investigating how much
overlap exists between ETFs that look different on paper but may hold
similar underlying assets.

## Status

Early development. Currently has:

- PostgreSQL schema for ETFs, holdings, and market/sector exposure
- A minimal Go API with a health check endpoint (`/health`)

## Tech Stack

- **Backend:** Go (standard library `net/http`, no framework yet)
- **Database:** PostgreSQL
- **Driver:** [pgx](https://github.com/jackc/pgx)

## Project Structure

```
cmd/api/       — application entry point (main.go)
schema.sql     — database schema (tables for ETFs, holdings, exposure)
```

## Setup

### Prerequisites

- Go 1.27+
- PostgreSQL running locally

### 1. Clone and install dependencies

```bash
git clone git@github.com:jpvargues/finance-platform.git
cd finance-platform
go mod download
```

### 2. Set up the database

```bash
createdb finance_platform
psql finance_platform -f schema.sql
```

### 3. Configure environment variables

Copy the example env file and adjust as needed:

```bash
cp .env.example .env
```

### 4. Run the server

```bash
go run cmd/api/main.go
```

Verify it's working:

```bash
curl http://localhost:8080/health
# {"status":"ok","database":"connected"}
```
