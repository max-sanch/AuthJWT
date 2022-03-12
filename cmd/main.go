package main

import (
	"github.com/max-sanch/AuthJWT"
	"log"
)

func main() {
	srv := new(auth_jwt.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
