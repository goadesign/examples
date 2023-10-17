package sessionapi

import (
	"context"
	"fmt"
	"log"

	"github.com/rs/xid"
	"goa.design/examples/cookies/gen/session"
)

// session service example implementation.
// The example methods log the requests and return zero values.
type sessionsrvc struct {
	logger *log.Logger
}

// NewSession returns the session service implementation.
func NewSession(logger *log.Logger) session.Service {
	return &sessionsrvc{logger}
}

// CreateSession implements create_session.
func (s *sessionsrvc) CreateSession(ctx context.Context, p *session.CreateSessionPayload) (res *session.CreateSessionResult, err error) {
	sid := xid.New().String()
	res = &session.CreateSessionResult{
		SessionID: sid,
		Message:   "new session created",
	}
	s.logger.Printf("created session %q", sid)
	return
}

// UseSession implements use_session.
func (s *sessionsrvc) UseSession(ctx context.Context, p *session.UseSessionPayload) (res *session.UseSessionResult, err error) {
	res = &session.UseSessionResult{
		Message: fmt.Sprintf("using session %q", p.SessionID),
	}
	s.logger.Printf("session %q", p.SessionID)
	return
}
