# Concert Management API

A practical example of building a REST API with [Goa](https://goa.design/),
covering design patterns, validation, error handling, and documentation
generation.

> **📚 Tutorial Reference**: This example is featured in the official
> [Goa REST API Tutorial](https://goa.design/docs/3-tutorials/1-rest-api/) on
> the Goa documentation site. The tutorial provides step-by-step guidance for
> building REST APIs with Goa's design-first approach.

> **⚠️ Example Code**: This is a simplified example for learning purposes. For 
> simplicity, the entire service implementation lives in the `main.go` file. 
> In practice, you would typically organize code into separate packages as 
> shown in the
> [Multiple Services guide](https://goa.design/docs/6-advanced/2-multiple-services/).

## What's This All About?

This is a concert management system that demonstrates how to build a clean REST
API. You can manage concerts, artists, venues, and ticket prices through
well-structured endpoints. The key advantage is that everything is defined
upfront in the design, and Goa generates all the boilerplate code for you.

What makes this approach useful:
- Clean, well-structured endpoints that follow REST conventions
- Validation that happens before your business logic runs
- Consistent error responses across all endpoints
- Pagination that handles large datasets properly
- Auto-generated OpenAPI documentation that stays current
- Strong typing that catches issues at compile time

## What You Can Do

The API covers the standard concert management operations:

**Concert Management**: List concerts with pagination, create new events, get
 details for specific concerts, update existing information, and remove
 concerts.

**Data Model**: Each concert includes a unique ID (auto-generated), artist or
 band name, date in ISO format, venue information, and ticket price stored in
 cents.

## Getting It Running

### What You'll Need
- Go 1.24 or later
- [Goa v3](https://goa.design/) framework

### Setup

Get the code and dependencies:
```bash
git clone <repository>
cd concerts
go mod tidy
```

Generate the Goa code:
```bash
goa gen concerts/design
```

Start the server:
```bash
go run cmd/concerts/main.go
```

The API will be available at `http://localhost:8080`.

### Documentation

Once running, you can view the auto-generated documentation:
- **OpenAPI JSON**: http://localhost:8080/openapi3.json
- **OpenAPI YAML**: http://localhost:8080/openapi3.yaml

## Try It Out

### List Concerts

```bash
# Get the first page of concerts
curl "http://localhost:8080/concerts"

# Get page 2 with 5 results
curl "http://localhost:8080/concerts?page=2&limit=5"
```

### Create a Concert

```bash
curl -X POST "http://localhost:8080/concerts" \
  -H "Content-Type: application/json" \
  -d '{
    "artist": "The White Stripes",
    "date": "2024-12-25",
    "venue": "Madison Square Garden, New York, NY",
    "price": 8500
  }'
```

### Get Concert Details

```bash
curl "http://localhost:8080/concerts/550e8400-e29b-41d4-a716-446655440000"
```

### Update a Concert

You can update all fields or just specific ones:

```bash
# Update multiple fields
curl -X PUT "http://localhost:8080/concerts/550e8400-e29b-41d4-a716-446655440000" \
  -H "Content-Type: application/json" \
  -d '{
    "artist": "The White Stripes",
    "date": "2024-12-26",
    "venue": "Madison Square Garden, New York, NY",
    "price": 9000
  }'

# Update just the price
curl -X PUT "http://localhost:8080/concerts/550e8400-e29b-41d4-a716-446655440000" \
  -H "Content-Type: application/json" \
  -d '{
    "price": 9500
  }'
```

### Delete a Concert

```bash
curl -X DELETE "http://localhost:8080/concerts/550e8400-e29b-41d4-a716-446655440000"
```

## Design Details

### Type Structure

The API uses a layered type approach for flexible operations:

- **ConcertData**: Base type with all concert fields, no required constraints
- **ConcertPayload**: For creation - extends base type and requires all fields
- **UpdatePayload**: For updates - extends base type but only requires the concert ID
- **Concert**: For responses - complete concert with ID and all details

This structure allows creation to require complete data while updates can be partial.

### Validation Rules

All validation is defined in the design layer:
- Artist names: 1-200 characters
- Dates: ISO 8601 format (YYYY-MM-DD) 
- Venues: 1-300 characters
- Prices: Non-negative integers, maximum $1000 (stored as cents)
- Pagination: Page ≥ 1, Limit 1-100

### Error Handling

The API returns consistent error responses:

```json
{
  "message": "Concert with ID abc123 not found",
  "code": "not_found"
}
```

HTTP status codes follow standard conventions:
- `200 OK` for successful reads and updates
- `201 Created` for successful creation  
- `204 No Content` for successful deletion
- `400 Bad Request` for validation errors
- `404 Not Found` for missing resources

## Project Structure

```bash
concerts/
├── design/
│   └── design.go   # API design specification
├── cmd/concerts/
│   └── main.go     # Service implementation 
├── gen/            # Generated code (don't modify)
│   ├── concerts/   # Service interfaces and types
│   └── http/       # HTTP transport layer
├── go.mod          # Dependencies
└── README.md       # This file
```

> **Note**: This example keeps everything in a single `main.go` file for 
> simplicity. In real applications, you would typically organize the service 
> implementation into separate packages, with clean separation between business 
> logic, transport handlers, and service interfaces as demonstrated in the 
> [Multiple Services documentation](https://goa.design/docs/6-advanced/2-multiple-services/).

## Development

To modify the API, edit `design/design.go` and regenerate:

```bash
goa gen concerts/design
```

Goa will update all the generated code while preserving your service implementation.

## Key Benefits

1. **Design-first development**: API specification drives implementation
2. **Centralized validation**: Rules defined once in the design
3. **Type safety**: Compile-time error detection
4. **Consistent error handling**: Standardized response format
5. **Always-current documentation**: Generated from the same source as code
6. **RESTful design**: Proper HTTP methods and resource modeling
7. **Efficient pagination**: Handles large datasets appropriately
8. **Clean separation**: Business logic separate from transport concerns

## Dependencies

- [Goa v3](https://goa.design/): Design framework for building APIs
- [UUID](https://github.com/google/uuid): For generating unique identifiers

This example demonstrates Goa's design-first approach for educational purposes.