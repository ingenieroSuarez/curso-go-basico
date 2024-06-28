package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"platzi.com/go/rest-ws/models"
	"platzi.com/go/rest-ws/server"
)

var (
	NO_AUTH_NEEDED = []string{ // rutas excentas del middleware
		"login",
		"signup",
	}
)

func shoulCheckToken(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false // quiere decir que la ruta no debe estar protegida, esta excenta.
		}
	}
	return true // es una ruta no excenta
}

func CheckAuthMiddleware(s server.Server) func(h http.Handler) http.Handler { // el middleware hace un salto al siguiente handler
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shoulCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r) // next=> puede seguir
				return
			}
			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
			// de acá arriba tomamos el token
			_, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) { //models.AppClaims{} es del modelo que tengo en la carpeta models
				return []byte(s.Config().JWTSecret), nil // el nil, significa que paso un error
			})

			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized) //401
				return
			}
			next.ServeHTTP(w, r)
		})

	}
}
