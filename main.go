package main

import (
	"log"

	"github.com/agustinLoag/restAPiTwitt/bd"
	"github.com/agustinLoag/restAPiTwitt/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
