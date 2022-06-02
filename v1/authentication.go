package v1

import "github.com/golang-jwt/jwt"

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
