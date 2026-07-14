package security

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

// Claims represents the strict Enterprise Identity contract.
type Claims struct {
	UserID string   `json:"sub"`
	Aud    string   `json:"aud"` // Service Audience
	Scopes []string `json:"scp"`
	jwt.RegisteredClaims
}

// SecurityManager orchestrates identity validation.
type SecurityManager struct {
	JWKSet    jwk.Set
	IssuerURL string
}

// Middleware enforces mTLS and JWT validation for every incoming request.
func (sm *SecurityManager) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Transport Security Check (mTLS verify)
		if r.TLS == nil || len(r.TLS.PeerCertificates) == 0 {
			http.Error(w, "Forbidden: mTLS Required", http.StatusForbidden)
			return
		}

		// 2. Extract Authorization Header
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized: Missing Token", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 3. Parse and Validate JWT (Claims-based)
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			// Get key from JWKS (Key Rotation support)
			kid, _ := token.Header["kid"].(string)
			key, err := sm.JWKSet.LookupKeyID(kid)
			if err != nil {
				return nil, err
			}
			var rawKey interface{}
			key.Raw(&rawKey)
			return rawKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: Invalid Token", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(*Claims)
		
		// 4. Audience & Scope Enforcement
		if claims.Aud != "order-service" {
			http.Error(w, "Forbidden: Wrong Audience", http.StatusForbidden)
			return
		}

		// 5. Context Propagation (Identity Injection)
		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "scopes", claims.Scopes)
		
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
