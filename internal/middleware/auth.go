package middleware

import(
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key")

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")

		if authHeader == ""{
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{},error){
			return jwtKey, nil
		})

		if err != nil || !token.Valid{
			http.Error(w, "imvalid token", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		userID := int(claims["user_id"].(float64))

		ctx := context.WithValue(r.Context(), "user_id", userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}