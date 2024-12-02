package authmiddleware

import (
	"context"
	"errors"
	"getfund-api-v2/internal/shared/applicationcode"
	"getfund-api-v2/internal/shared/proxy"
	"getfund-api-v2/internal/shared/service/sessionservice"
	"net/http"
	"strings"
)

type AuthMiddleware interface {
	Authenticate(next http.Handler) http.Handler
}

type authMiddleware struct {
	session sessionservice.SessionService
}

func New(sessionService sessionservice.SessionService) AuthMiddleware {
	return &authMiddleware{session: sessionService}
}

func (a *authMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := extractToken(r)
		if token == "" {
			proxy.SetError(w, applicationcode.CODE_UNAUTHORIZED, errors.New("unauthorized"))
			return
		}

		sessionSerialized, err := a.session.GetSession(token)
		if err != nil || sessionSerialized == "" {
			proxy.SetError(w, applicationcode.CODE_UNAUTHORIZED, errors.New("unauthorized"))
			return
		}

		ctx := context.WithValue(r.Context(), "session", sessionSerialized)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}
	return parts[1]
}
