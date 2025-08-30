# GoRedis - Simple Redis-like Caching Tool

A lightweight, high-performance in-memory key-value store with HTTP and gRPC interfaces, built in Go.

## MVP Overview

GoRedis is a simplified Redis-like caching solution designed for fast development and deployment. The MVP focuses on core caching functionality with a clean HTTP API, while providing a foundation for future gRPC implementation.

## Core Features (MVP)

### 1. Basic Key-Value Operations
- **SET**: Store a key-value pair with optional TTL
- **GET**: Retrieve value by key
- **DELETE**: Remove a key-value pair
- **EXISTS**: Check if a key exists
- **KEYS**: List all keys (with optional pattern matching)

### 2. Data Types Support
- **Strings**: Basic string values
- **Numbers**: Integer and float support
- **JSON**: Store and retrieve JSON objects

### 3. Expiration & TTL
- Set expiration time for keys (TTL - Time To Live)
- Automatic cleanup of expired keys
- Background garbage collection

### 4. HTTP REST API
- RESTful endpoints for all operations
- JSON request/response format
- Proper HTTP status codes
- Request validation and error handling

### 5. Persistence (Optional for MVP)
- In-memory storage (primary)
- Optional snapshot-to-disk functionality
- Configurable persistence intervals

## Technical Requirements

### Performance Targets
- Sub-millisecond response times for basic operations
- Support for 10,000+ concurrent connections
- Memory-efficient storage
- Horizontal scaling readiness

### Architecture
- **Language**: Go (Golang)
- **Storage**: In-memory with concurrent-safe data structures
- **HTTP Server**: High-performance HTTP server (gin-gonic/gin or net/http)
- **Configuration**: YAML/JSON configuration files
- **Logging**: Structured logging with levels
- **Health Checks**: Built-in health and metrics endpoints

## API Specification

### HTTP Endpoints

#### Key-Value Operations
```
POST   /api/v1/keys/{key}          # Set key-value
GET    /api/v1/keys/{key}          # Get value by key
DELETE /api/v1/keys/{key}          # Delete key
HEAD   /api/v1/keys/{key}          # Check if key exists
GET    /api/v1/keys                # List all keys
```

#### Administrative
```
GET    /health                     # Health check
GET    /metrics                    # Basic metrics
POST   /api/v1/flush               # Clear all data
GET    /api/v1/info                # Server information
```

### Request/Response Examples

#### Set Key-Value
```bash
POST /api/v1/keys/user:123
Content-Type: application/json

{
  "value": "john_doe",
  "ttl": 3600
}
```

#### Get Value
```bash
GET /api/v1/keys/user:123

Response:
{
  "key": "user:123",
  "value": "john_doe",
  "ttl": 3542,
  "created_at": "2025-08-27T10:30:00Z"
}
```

## Project Structure

```
goredis/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── cache/
│   │   ├── store.go             # Core cache implementation
│   │   ├── item.go              # Cache item structure
│   │   └── cleanup.go           # TTL cleanup routines
│   ├── api/
│   │   ├── handlers.go          # HTTP handlers
│   │   ├── middleware.go        # HTTP middleware
│   │   └── routes.go            # Route definitions
│   ├── config/
│   │   └── config.go            # Configuration management
│   └── metrics/
│       └── metrics.go           # Metrics collection
├── pkg/
│   └── client/
│       └── http.go              # HTTP client library
├── configs/
│   └── server.yaml              # Default configuration
├── docs/
│   ├── api.md                   # API documentation
│   └── deployment.md            # Deployment guide
├── scripts/
│   ├── build.sh                 # Build scripts
│   └── docker-build.sh          # Docker build script
├── Dockerfile                   # Container definition
├── docker-compose.yml           # Local development setup
├── go.mod                       # Go modules
├── go.sum                       # Dependencies checksum
├── Makefile                     # Build automation
└── README.md                    # This file
```

## Implementation Phases

### Phase 1: Core MVP (Week 1-2)
- [x] Basic in-memory key-value store
- [x] HTTP API with essential endpoints
- [x] TTL support and background cleanup
- [x] Configuration management
- [x] Basic logging and error handling

### Phase 2: Enhanced Features (Week 3)
- [ ] Pattern-based key searching
- [ ] Bulk operations (MGET, MSET)
- [ ] Basic persistence (snapshots)
- [ ] Metrics and monitoring endpoints
- [ ] Docker containerization

### Phase 3: gRPC Implementation (Week 4)
- [ ] Protocol buffer definitions
- [ ] gRPC server implementation
- [ ] Performance benchmarking
- [ ] Client libraries for both HTTP and gRPC

### Phase 4: Production Ready (Week 5+)
- [ ] Authentication and authorization
- [ ] Rate limiting
- [ ] Clustering support
- [ ] Advanced persistence options
- [ ] Comprehensive testing suite

## Configuration Example

```yaml
# configs/server.yaml
server:
  host: "0.0.0.0"
  http_port: 8080
  grpc_port: 9090
  read_timeout: 10s
  write_timeout: 10s

cache:
  max_memory: "1GB"
  cleanup_interval: "1m"
  default_ttl: "1h"
  max_key_size: 250
  max_value_size: "1MB"

logging:
  level: "info"
  format: "json"
  output: "stdout"

persistence:
  enabled: false
  snapshot_interval: "5m"
  snapshot_path: "./data/snapshots"
```

## Getting Started

### Prerequisites
- Go 1.21 or higher
- Docker (optional)

### Quick Start
```bash
# Clone the repository
git clone <repo-url>
cd goredis

# Build the application
make build

# Run the server
./bin/goredis-server

# Test the API
curl -X POST http://localhost:8080/api/v1/keys/test \
  -H "Content-Type: application/json" \
  -d '{"value": "hello world", "ttl": 60}'

curl http://localhost:8080/api/v1/keys/test
```

## Future Enhancements

### gRPC Interface
- High-performance binary protocol
- Streaming support for bulk operations
- Auto-generated client libraries
- Service mesh compatibility

### Advanced Features
- Lua scripting support
- Pub/Sub messaging
- Distributed caching with consistent hashing
- Cross-datacenter replication
- Advanced data types (Lists, Sets, Hashes)

## Performance Benchmarks (Target)

| Operation | Throughput | Latency (p99) |
|-----------|------------|---------------|
| SET       | 100k ops/s | < 1ms         |
| GET       | 150k ops/s | < 0.5ms       |
| DELETE    | 80k ops/s  | < 1ms         |

## Contributing

See `CONTRIBUTING.md` for development guidelines and contribution process.

## License

MIT License - see `LICENSE` file for details.
