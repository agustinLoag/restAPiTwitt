package routers

import (
	"encoding/json"
	"net/http"

	"github.com/agustinLoag/restAPiTwitt/bd"
	"github.com/agustinLoag/restAPiTwitt/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido"+err.Error(), http.StatusBadRequest)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "El Password debe de ser al menos de 6 caracteres"+err.Error(), http.StatusBadRequest)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "El Email esta ya registrado"+err.Error(), 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al guardar"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado realizar el registro"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
