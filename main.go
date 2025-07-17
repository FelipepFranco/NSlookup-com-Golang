package main

import (
	app "aplicacao_de_linha_de_comando/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil{
		log.Fatal(erro)
	}
}