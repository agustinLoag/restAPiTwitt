package middlewares

import (
	"net/http"

	"github.com/agustinLoag/restAPiTwitt/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion Perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
