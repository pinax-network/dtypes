package authentication

import (
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"strings"
)

type JwtCredentials struct {
	jwt.StandardClaims
	IP       string                   `json:"-"`
	ApiKeyId string                   `json:"api_key_id"`
	Networks []NetworkPermissionClaim `json:"networks"`
}

type NetworkPermissionClaim struct {
	Name  string `json:"name"`
	Quota int    `json:"quota"`
	Rate  int    `json:"rate"`
}

func (c *JwtCredentials) GetUserID() string {
	userID := c.Subject
	return strings.TrimPrefix(userID, "uid:")
}

func (c *JwtCredentials) GetLogFields() []zap.Field {
	return []zap.Field{
		zap.String("subject", c.Subject),
		zap.String("jti", c.Id),
		zap.String("api_key_id", c.ApiKeyId),
		zap.String("ip", c.IP),
	}
}
