package tokenauthority

import "github.com/Laelapa/CompanyRegistry/internal/config"

type TokenAuthority struct {
	cfg *config.AuthConfig
}

func New(cfg *config.AuthConfig) *TokenAuthority {
	t := &TokenAuthority{
		cfg: cfg,
	}
	return t
}
