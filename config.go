package ganblib

type Config struct {
	Salt        string
	AuthBaseURL string
	AuthPath    string
	TokenPath   string
	JwtIssuer   string
}

// OverrideConfig overrides config
func OverrideConfig(c *Config) {
	config = conf{
		Salt:        c.Salt,
		AuthBaseURL: c.AuthBaseURL,
		AuthPath:    c.AuthPath,
		TokenPath:   c.TokenPath,
		JwtIssuer:   c.JwtIssuer,
	}
}
