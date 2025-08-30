# GoRedis Postman Collection

This directory contains Postman collection and environment files for testing the GoRedis cache API.

## Files

- `GoRedis_Cache_API.postman_collection.json` - Main collection with all API endpoints
- `GoRedis_Local.postman_environment.json` - Environment variables for local testing

## How to Import

### Option 1: Import Collection File
1. Open Postman
2. Click "Import" button
3. Select "Upload Files"
4. Choose `GoRedis_Cache_API.postman_collection.json`
5. Import the environment file `GoRedis_Local.postman_environment.json`

### Option 2: Import via URL (if hosted)
```
https://raw.githubusercontent.com/yourusername/goredis/main/postman/GoRedis_Cache_API.postman_collection.json
```

## Setup

1. **Import both files** into Postman
2. **Select the environment**: Choose "GoRedis Local Environment" from the environment dropdown
3. **Start your server**: Make sure GoRedis is running on `localhost:8080`
4. **Test the endpoints**: Start with the Health Check request

## Available Endpoints

### Health & Info
- `GET /health` - Check server health
- `GET /info` - Get server information

### Cache Operations
- `POST /api/v1/keys/{key}` - Set a key-value pair
- `GET /api/v1/keys/{key}` - Get value by key
- `DELETE /api/v1/keys/{key}` - Delete a key
- `HEAD /api/v1/keys/{key}` - Check if key exists
- `GET /api/v1/keys` - List all keys
- `POST /api/v1/flush` - Clear all cached data

## Testing Scenarios

### Basic Operations
1. **Health Check** - Verify server is running
2. **Set Simple Key** - Store a string value
3. **Get Key** - Retrieve the stored value
4. **Delete Key** - Remove the key
5. **List Keys** - See all stored keys

### Advanced Operations
1. **Set with TTL** - Store value with expiration
2. **Set JSON Object** - Store complex data structures
3. **Check Expiration** - Wait and verify TTL works
4. **Flush All** - Clear entire cache

## Environment Variables

| Variable | Default Value | Description |
|----------|---------------|-------------|
| `base_url` | `http://localhost:8080` | GoRedis server URL |
| `api_version` | `v1` | API version |

## Example Usage

### Setting a Simple Value
```bash
curl -X POST http://localhost:8080/api/v1/keys/mykey \
  -H "Content-Type: application/json" \
  -d '{"value": "Hello World"}'
```

### Setting with TTL (60 seconds)
```bash
curl -X POST http://localhost:8080/api/v1/keys/temp_key \
  -H "Content-Type: application/json" \
  -d '{"value": "Expires in 60 seconds", "ttl": 60}'
```

### Getting a Value
```bash
curl http://localhost:8080/api/v1/keys/mykey
```

### Setting JSON Object
```bash
curl -X POST http://localhost:8080/api/v1/keys/user:123 \
  -H "Content-Type: application/json" \
  -d '{
    "value": {
      "name": "John Doe",
      "age": 30,
      "email": "john@example.com"
    },
    "ttl": 300
  }'
```

## Response Formats

### Success Response (SET)
```json
{
  "status": "success",
  "key": "mykey"
}
```

### Success Response (GET)
```json
{
  "status": "success",
  "value": "Hello World"
}
```

### Error Response
```json
{
  "error": "key not found"
}
```

## Tips

1. **Use variables**: The collection uses `{{base_url}}` variable for easy environment switching
2. **Test order**: Run Health Check first to ensure server is running
3. **TTL testing**: Set a key with short TTL (5-10 seconds) to test expiration
4. **JSON validation**: Use Postman's JSON validator in Tests tab
5. **Automation**: Set up Tests in Postman to automate response validation
