package routers

import (
	"encoding/json"
	"net/http"

	"github.com/pagus97/twetpad/bd"
	"github.com/pagus97/twetpad/models"
)

/*ModificarPerfil modifica el perfil del usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro. Reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se a logrado modificar el registro del usuario"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
