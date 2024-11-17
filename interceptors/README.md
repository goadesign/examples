# Interceptors Example

This example demonstrates the concepts of various interceptors in a Goa service. It is intended as a learning tool to understand how interceptors work in Goa and should not be used directly in production environments.

## ⚠️ Important Note

This code is for educational purposes only. The implementations are intentionally simplified to highlight the concepts and mechanics of Goa interceptors. In a production environment, you would need to:

- Implement proper JWT validation and security measures
- Use appropriate caching strategies and storage backends
- Add comprehensive error handling and logging
- Follow your organization's security best practices
- Consider performance implications and scaling requirements

The interceptors demonstrated include:
- Authentication (JWT) - basic example of token validation
- Caching - simple in-memory implementation
- Audit Logging - basic request/response logging
- Request Tracing - simplified trace ID generation
- Retry Mechanism - basic retry logic
- Request Deadlines - simple timeout handling

## Running the Example

1. Start the service:
   ```bash
   ./run-service.sh
   ```

2. In another terminal, run the demo:
   ```bash
   ./demo.sh
   ```

## What the Demo Shows

The demo script demonstrates several interceptor features:

1. **Basic Flow**: Creates a record and retrieves it, showing authentication and audit logging in action
2. **Caching**: Makes the same request twice to demonstrate caching
3. **Retry Mechanism**: Tests the retry interceptor by requesting an unavailable resource
4. **Deadline**: Shows how the deadline interceptor handles slow requests
5. **Authentication**: Demonstrates JWT token validation

## Service Endpoints

The service runs on both HTTP on `http://localhost:8088` and gRPC on `localhost:8080`.
The HTTP service provides two endpoints:
- `POST /records/{tenantID}` - Create a new record
- `GET /records/{tenantID}/{recordID}` - Retrieve a record by ID

Each endpoint is protected by various interceptors that demonstrate different aspects of the middleware functionality. 
