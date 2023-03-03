package middlewares

import (
	"net/http"

	"github.com/Dedemahendra1/go_tutorial/config"
	"github.com/Dedemahendra1/go_tutorial/helper"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}
		// mengambil token value
		tokenString := c.Value

		claims := &config.JWTClaim{}
		// parsing token jwt
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				// token invalid
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// token expired
				response := map[string]string{"message": "Unauthorized, Token expired!"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			} else {
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
