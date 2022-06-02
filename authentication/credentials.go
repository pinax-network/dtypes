package authentication

import "go.uber.org/zap"

type Credentials interface {
	GetLogFields() []zap.Field
	GetUserID() string
}

type AnonymousCredentials struct {
	// MUST implement Credentials
	userID string
	ip     string
}

func newAnonymousCredentials() *AnonymousCredentials {
	return &AnonymousCredentials{
		userID: "anonymous",
		ip:     "0.0.0.0",
	}
}

func (c *AnonymousCredentials) GetUserID() string {
	return c.userID
}

func (c *AnonymousCredentials) GetLogFields() []zap.Field {
	return []zap.Field{
		zap.String("subject", c.userID),
		zap.String("api_key_id", c.userID),
		zap.String("ip", c.ip),
	}
}
