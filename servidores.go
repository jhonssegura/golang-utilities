package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, servidor := range servidores {
		revisarServidor(servidor)
	}
	tiempo := time.Since(inicio)
	fmt.Println("Tiempo de ejecucion %s", tiempo)
}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "Servidor caido")
	} else {
		fmt.Println(servidor, "Servidor funcionando")
	}
}
