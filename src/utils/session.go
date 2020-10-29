package utils

import (
	"context"

	"github.com/satori/go.uuid"
	"google.golang.org/grpc/metadata"
)

// NewContext: new context with session id
func NewContext(ctx context.Context) (context.Context, string) {
	sessionId := NewSessionId()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("session-id", sessionId))
	return ctx, sessionId
}

// NewContextWithSession: create new session with given session id
func NewContextWithSession(ctx context.Context, session string) context.Context {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("session-id", session))
	return ctx
}

// GetSessionIdFromContext: get session id in context
func GetSessionIdFromContext(ctx context.Context) string {
	out_md, out_ok := metadata.FromOutgoingContext(ctx)
	in_md, in_ok := metadata.FromIncomingContext(ctx)
	if out_ok {
		sessionIds, ok := out_md["session-id"]
		if !ok {
			return ""
		}
		return sessionIds[0]
	}

	if in_ok {
		sessionIds, ok := in_md["session-id"]
		if !ok {
			return ""
		}
		return sessionIds[0]
	}
	return ""
}

// NewSessionId: new session id, uuid
func NewSessionId() string {
	return uuid.NewV4().String()
}
