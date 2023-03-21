package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pagus97/twetpad/bd"
	"github.com/pagus97/twetpad/models"
)

/*GraboTweet permite grabar el Tweet en la BD*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se a logrado insertar el tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
