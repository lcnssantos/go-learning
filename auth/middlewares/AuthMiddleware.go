package middlewares

import (
	"context"
	"errors"
	"main/auth/services"
	"main/shared"
	"net/http"
	"strings"
)

type AuthenticationMiddleware struct {
	authService *services.AuthService
}

func NewAuthenticationMiddleware(authService *services.AuthService) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{authService: authService}
}

func (this *AuthenticationMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			shared.ThrowHttpError(w, http.StatusUnauthorized, "Missing JWT Token")
			return
		}

		bearerToken := strings.Replace(authorization, "Bearer ", "", 1)

		user, err := this.authService.GetByToken(bearerToken)

		if err != nil {
			shared.ThrowHttpError(w, http.StatusUnauthorized, err.Error())
			return
		}

		if !user.EmailConfirmed || !user.IsActive {
			shared.ThrowHttpError(w, http.StatusUnauthorized, errors.New("User not active").Error())
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
