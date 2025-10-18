package auth_middleware

import (
	"context"
	"encoding/json"
	"errors"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	shared_error "getfund-api-v2/internal/shared/error"
	shared_response_proxy "getfund-api-v2/internal/shared/proxy"
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
	session auth_contract.SessionService
}

func New(sessionService auth_contract.SessionService) AuthMiddleware {
	return &authMiddleware{session: sessionService}
}

func (a *authMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := extractToken(r)
		if token == "" {
			shared_response_proxy.SetError(w, shared_error.UNAUTHORIZED_CODE, errors.New("unauthorized"))
			return
		}

		sessionSerialized, err := a.session.GetSession(token)
		if err != nil || sessionSerialized == "" {
			shared_response_proxy.SetError(w, shared_error.UNAUTHORIZED_CODE, errors.New("unauthorized"))
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, auth_contract.TokenKey{}, token)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *authMiddleware) AuthenticateAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := extractToken(r)
		if token == "" {
			shared_response_proxy.SetError(w, shared_error.UNAUTHORIZED_CODE, errors.New("unauthorized"))
			return
		}

		sessionSerialized, err := a.session.GetSession(token)
		if err != nil || sessionSerialized == "" {
			shared_response_proxy.SetError(w, shared_error.UNAUTHORIZED_CODE, errors.New("unauthorized"))
			return
		}

		session := &sessionModel{}
		errSession := json.Unmarshal([]byte(sessionSerialized), &session)
		if errSession != nil || session.IdAdmin == 1 {
			shared_response_proxy.SetError(w, shared_error.UNAUTHORIZED_CODE, errors.New("unauthorized"))
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, auth_contract.SessionKey{}, sessionSerialized)
		ctx = context.WithValue(ctx, auth_contract.TokenKey{}, token)

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
