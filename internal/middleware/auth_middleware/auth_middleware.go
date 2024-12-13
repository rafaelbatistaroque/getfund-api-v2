package auth_middleware

import (
	"context"
	"encoding/json"
	"errors"
	"getfund-api-v2/internal/proxy"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/service/session_service"
	"net/http"
	"strings"
)

type sessionModel struct {
	IdAdmin int `json:"is_admin"`
}

type AuthMiddleware interface {
	Authenticate(next http.Handler) http.Handler
	AuthenticateAdmin(next http.Handler) http.Handler
}

type authMiddleware struct {
	session session_service.SessionService
}

func New(sessionService session_service.SessionService) AuthMiddleware {
	return &authMiddleware{session: sessionService}
}

func (a *authMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := extractToken(r)
		if token == "" {
			proxy.SetError(w, result_app.UNAUTHORIZED_CODE, errors.New("unauthorized"))
			return
		}

		sessionSerialized, err := a.session.GetSession(token)
		if err != nil || sessionSerialized == "" {
			proxy.SetError(w, result_app.UNAUTHORIZED_CODE, errors.New("unauthorized"))
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, session_service.TokenKey{}, token)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *authMiddleware) AuthenticateAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := extractToken(r)
		if token == "" {
			proxy.SetError(w, result_app.UNAUTHORIZED_CODE, errors.New("unauthorized"))
			return
		}

		sessionSerialized, err := a.session.GetSession(token)
		if err != nil || sessionSerialized == "" {
			proxy.SetError(w, result_app.UNAUTHORIZED_CODE, errors.New("unauthorized"))
			return
		}

		session := &sessionModel{}
		errSession := json.Unmarshal([]byte(sessionSerialized), &session)
		if errSession != nil || session.IdAdmin == 1 {
			proxy.SetError(w, result_app.UNAUTHORIZED_CODE, errors.New("unauthorized"))
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, session_service.SessionKey{}, sessionSerialized)
		ctx = context.WithValue(ctx, session_service.TokenKey{}, token)

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
