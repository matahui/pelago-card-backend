# Pelago Card Backend

A microservices-based card management backend system built with Go-Zero framework, focusing on merchant onboarding functionality.

## Architecture

- **API Service**: RESTful API gateway handling HTTP requests
- **RPC Service**: Internal business logic service with gRPC communication
- **Database**: MySQL for data persistence
- **Cache**: Redis for caching (optional)

## Technology Stack

- **Framework**: [Go-Zero](https://go-zero.dev/)
- **Language**: Go 1.21+
- **Database**: MySQL 8.0+
- **Protocol**: gRPC, HTTP/REST
- **Tools**: goctl (Go-Zero CLI tool)

## Project Structure

```
pelago-card-backend/
├── api/                    # API service (HTTP gateway)
│   ├── apis/              # API definitions
│   ├── internal/          # Internal logic
│   └── card.go           # Main entry point
├── rpc/                   # RPC service (business logic)
│   ├── card/             # Card service
│   ├── pb/               # Generated protobuf files
│   └── proto/            # Protobuf definitions
├── sql/                   # Database schemas
├── docker/               # Docker configurations
├── bin/                  # Compiled binaries
└── scripts/              # Utility scripts
```

## Prerequisites

- Go 1.21 or higher
- MySQL 8.0+
- Docker & Docker Compose (optional)

### Install Go-Zero CLI Tool

```bash
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

## Quick Start

### 1. Database Setup

Start MySQL using Docker:

```bash
docker-compose up -d mysql
```

Or manually create database and import schema:

```bash
mysql -u root -p < sql/v1.0.0/init.sql
```

### 2. Configuration

Update database connection in configuration files:
- `rpc/card/etc/card.yaml`
- `api/etc/card-api.yaml`

### 3. Build Services

```bash
# Build RPC service
go build -o bin/card-app app/card/card.go

# Build API service  
go build -o bin/card-portal portal/card.go
```

### 4. Run Services

Start RPC service:
```bash
./bin/card-app -f app/card/etc/card.yaml
```

Start API service:
```bash
./bin/card-portal -f portal/etc/card-portal.yaml
```

## API Endpoints

Base URL: `http://localhost:8888`

### Merchant Management

- `POST /api/merchant/register` - Register new merchant
- `GET /api/merchant/{merchant_id}` - Get merchant details
- `PUT /api/merchant/update` - Update merchant information
- `GET /api/merchant/list` - List merchants with pagination

## Development

### Code Generation

Generate API code:
```bash
cd portal && goctl portal go -portal api/card.portal -dir . --style=gozero
```

Generate RPC code:
```bash
cd app && goctl app protoc proto/merchant.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=./card --style=gozero
```

Generate model code:
```bash
cd app/card/internal && goctl model mysql datasource -url="root:password@tcp(localhost:3306)/pelago_card" -table="merchants" -dir="./model" --style=gozero
```

### Database Schema

The system uses a single `merchants` table for merchant onboarding:

```sql
CREATE TABLE merchants (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    merchant_id VARCHAR(50) UNIQUE NOT NULL,
    entity_name VARCHAR(200) NOT NULL,
    brand_name VARCHAR(100) NOT NULL,
    website_url VARCHAR(500),
    merchant_logo VARCHAR(500),
    merchant_country VARCHAR(50) NOT NULL,
    contact_name VARCHAR(100) NOT NULL,
    contact_email VARCHAR(100) NOT NULL,
    api_key VARCHAR(100) NOT NULL,
    verified_status VARCHAR(50) DEFAULT 'Pending for approval',
    merchant_status VARCHAR(50) DEFAULT 'Inactive',
    business_description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## License

This project is proprietary software owned by Pelago.

## Contact

For questions or support, please contact the development team. 