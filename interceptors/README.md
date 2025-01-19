# Interceptors Example

This example demonstrates how to structure and organize interceptors in a Goa service. It provides skeleton implementations that show the basic patterns for both client-side and server-side interceptors. Note that these are educational examples and not production-ready implementations.

## Design Overview

The example shows how to organize different types of interceptors:

### Server-Side Interceptors

```go
// Example interceptor chain - implementations are skeleton/logging only
ServerInterceptor(TraceRequest)  // Shows where to add trace context
ServerInterceptor(RequestAudit)  // Shows where to add timing/logging
ServerInterceptor(JWTAuth)       // Shows where to add auth checks
ServerInterceptor(SetDeadline)   // Shows where to add timeouts
ServerInterceptor(Cache)         // Shows where to add caching
```

### Client-Side Interceptors

```go
// Example client interceptors - implementations are skeleton/logging only
ClientInterceptor(EncodeTenant)  // Shows where to modify auth tokens
ClientInterceptor(Retry)         // Shows where to add retry logic
```

## Example Patterns Demonstrated

The example shows common interceptor patterns (note: implementations are skeleton only):

1. **Authentication Pattern**
   - Where to add JWT token validation
   - How to pass tenant information
   - Structure for auth checks

2. **Observability Pattern**
   - Where to add request tracing
   - How to structure audit logging
   - Points for adding metrics

3. **Caching Pattern**
   - Where to add cache checks
   - Structure for cache updates
   - Points for cache invalidation

4. **Resilience Pattern**
   - Where to add timeouts
   - Structure for retry logic
   - Error handling points

## Running the Example

1. Start the service:
   ```bash
   ./run-service.sh
   ```
   This builds and starts both the service and CLI on port 8088.

2. In another terminal, run the demo:
   ```bash
   ./demo.sh
   ```

## Demo Scenarios

The demo script shows the interceptor chain in action with test scenarios:

1. **Basic Request Flow**
   ```bash
   # Shows the complete interceptor chain with logging
   interceptors-cli create --tenant-id <UUID> --auth "Bearer <token>"
   ```

2. **Cache Structure**
   ```bash
   # Shows where cache checks would happen
   interceptors-cli get --record-id <UUID> --tenant-id <UUID> --auth "Bearer <token>"
   ```

3. **Retry Pattern**
   ```bash
   # Shows where retry logic would be implemented
   interceptors-cli get --record-id "00000000-0000-0000-0000-000000000000"
   ```

4. **Timeout Pattern**
   ```bash
   # Shows where deadline handling would occur
   interceptors-cli get --record-id "00000000-0000-0000-0000-000000000001"
   ```

5. **Auth Pattern**
   ```bash
   # Shows where token validation would happen
   interceptors-cli get --record-id <UUID> --auth "Bearer wrong-token"
   ```

## Implementation Notes

The interceptors in this example are skeleton implementations that only log their execution. They demonstrate:

1. How to structure interceptors in the design
2. Where to place different types of middleware logic
3. How to chain interceptors in a specific order
4. How client and server interceptors interact

Key Files:
- `design/design.go`: Shows how to define interceptors in Goa DSL
- `interceptors/interceptors_server.go`: Shows server interceptor structure
- `interceptors/interceptors_client.go`: Shows client interceptor structure
- `demo.sh`: Shows how to test interceptor patterns

## Service Endpoints

The service runs on both HTTP on `http://localhost:8088` and gRPC on `localhost:8080`.
The HTTP service provides two endpoints:
- `POST /records/{tenantID}` - Create a new record
- `GET /records/{tenantID}/{recordID}` - Retrieve a record by ID

Each endpoint is protected by various interceptors that demonstrate different aspects of the middleware functionality. 
