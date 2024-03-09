package trace

import "github.com/google/uuid"

// TraceID is unique across the lifecycle of a single 'event', reagardless of many
// requests it takes to complete. Carried in the `X-Trace-ID` header.
// RequestID is unique to each request. Carried in the `X-Request-ID` header.
type Trace struct {
	TraceID   uuid.UUID
	RequestID uuid.UUID
}
