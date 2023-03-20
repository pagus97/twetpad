package main

import (
	"log"

	"github.com/pagus97/twetpad/bd"
	"github.com/pagus97/twetpad/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
