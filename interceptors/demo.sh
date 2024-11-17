#!/bin/bash
set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo_step() {
    echo -e "\n${GREEN}$1${NC}"
    echo -e "${YELLOW}$2${NC}\n"
}

CLI="./bin/interceptors-cli"

# Ensure the CLI is built
if [ ! -f "$CLI" ]; then
    echo "CLI not found. Please run ./run-service.sh first"
    exit 1
fi

# Helper function to make requests
make_request() {
    $CLI --url http://localhost:8088 --verbose "$@"
}

# Create JSON payloads
TENANT_ID="bcd44737-a547-4294-8c51-a6665750b1b4"
CREATE_PAYLOAD='{"value": "Hello World"}'
GET_PAYLOAD='{}'

# Demo scenario
echo_step "1. Create a new record" \
         "This demonstrates the basic flow with authentication and audit logging"
RECORD_ID=$(make_request interceptors create --body "$CREATE_PAYLOAD" --tenant-id "$TENANT_ID" --auth "Bearer ${TENANT_ID}" | grep -o '"id":"[^"]*"' | cut -d'"' -f4)
echo "Created record with ID: $RECORD_ID"

echo_step "2. Retrieve the created record" \
         "This demonstrates caching interceptor in action"
make_request interceptors get --body "$GET_PAYLOAD" --record-id "$RECORD_ID" --tenant-id "$TENANT_ID" --auth "Bearer ${TENANT_ID}"
make_request interceptors get --body "$GET_PAYLOAD" --record-id "$RECORD_ID" --tenant-id "$TENANT_ID" --auth "Bearer ${TENANT_ID}"

echo_step "3. Test unavailable service (retry interceptor)" \
         "This demonstrates the retry mechanism when service is temporarily unavailable"
make_request interceptors get --body "$GET_PAYLOAD" --record-id "00000000-0000-0000-0000-000000000000" --tenant-id "$TENANT_ID" --auth "Bearer ${TENANT_ID}"

echo_step "4. Test timeout (deadline interceptor)" \
         "This demonstrates the deadline interceptor timing out a slow request"
make_request interceptors get --body "$GET_PAYLOAD" --record-id "00000000-0000-0000-0000-000000000001" --tenant-id "$TENANT_ID" --auth "Bearer ${TENANT_ID}"

echo_step "5. Test invalid tenant ID" \
         "This demonstrates the JWT authentication interceptor rejecting invalid tokens"
make_request interceptors get --body "$GET_PAYLOAD" --record-id "$RECORD_ID" --tenant-id "$TENANT_ID" --auth "Bearer wrong-token"
