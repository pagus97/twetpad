package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pagus97/twetpad/bd"
	"github.com/pagus97/twetpad/models"
)

/*Email valor de Email usado en todos los EndsPoints */
var Email string

/*IDUsuario es el ID devuelto del modelo que se usara en todos los EndsPoints */
var IDUsuario string

/*ProcesoToken proceso token para extraer*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encotrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encotrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encotrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}

	return claims, false, string(""), err
}
