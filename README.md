# Airplane Dealers Divar

A Divar-like advertising platform for buying and selling airplanes between airlines, built with Go and Echo. Supports role-based access, expert/repair inspection requests, bookmarks, payment, and a full Swagger API.

## Features

**Ads**
- Post, edit, and manage airplane listings (model, price, flight hours, age, images)
- Filter ads by category, price, age, and more
- Bookmark ads
- Request expert inspection or repair check on an ad

**Users & Roles**
- Sign up / log in with JWT authentication
- Roles: `SuperUser (Matin)`, `Admin`, `Expert`, `Airline`
- Admin panel for managing users, ads, and system config

**Payments**
- Payment gateway creation and verification

**Other**
- Activity logging service layer
- Swagger UI at `/swagger/index.html`
- Database migrations with `golang-migrate`

## Tech Stack
- **Go** + **Echo** framework
- **PostgreSQL** + **GORM**
- **JWT** authentication
- **Docker** + **Docker Compose**
- **Swagger** (swaggo/echo-swagger)

## Setup

### With Docker (recommended)
```bash
cp .env.example .env       # configure DB credentials and secrets
docker compose up --build
```
Runs migrations and seeds example data automatically.

### Without Docker
```bash
cp .env.example .env
go mod download

# Run migrations
migrate -path database/migrations -database 'postgres://...' up

# Start server
go run main.go
```

Server runs at `http://localhost:8080`  
Swagger UI: `http://localhost:8080/swagger/index.html`

## Environment Variables
| Variable | Description |
|----------|-------------|
| `POSTGRES_*` | PostgreSQL connection settings |
| `SECRET` | JWT signing secret |
| `ADMIN_CODE` | Registration code for admin role |
| `EXPERT_CODE` | Registration code for expert role |

## API Overview
| Domain | Endpoints |
|--------|-----------|
| Users | Register, login, profile |
| Ads | CRUD, filter, search |
| Bookmarks | Add/remove/list saved ads |
| Expert | Request expert inspection |
| Repair | Request repair check |
| Payment | Create and verify payments |

## Running Tests
```bash
go test ./...
```

## Project Structure
```
├── handlers/       # HTTP handlers per domain
├── datastore/      # DB layer (PostgreSQL via GORM)
├── models/         # Data models
├── server/         # Route registration
├── service/        # Service layer (logging)
├── filter/         # Query filtering logic
├── middlewares/    # JWT auth middleware
├── database/       # Connection + migrations
├── config/         # Config loader (.env + YAML)
└── docs/           # Swagger generated docs
```
