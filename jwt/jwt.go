package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pagus97/twetpad/models"
)

/*GenerJWT genera el encriptado con JWT */
func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellido":         t.Apellido,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
