package middleware

import (
	"github.com/Laelapa/CompanyRegistry/auth/tokenauthority"
	"github.com/Laelapa/CompanyRegistry/logging"
)

func AuthenticateWithJWT(
	tokenAuthority *tokenauthority.TokenAuthority,
	logger *logging.Logger,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Check for Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				logger.Warn("Unauthorized request: Missing Authorization header", logger.ReqFields(r)...)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			tokenString, err := net.StripBearer(authHeader)

