package main

import (
	"log"
)

func main() {
	srv := new(www.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}
