package common

type JwtConfig struct {
	Issuer    string `json:"issuer"`
	Audience  string `json:"audience"`
	Expires   int64  `json:"expires"`
	SecretKey string `json:"secret_key"`
}
