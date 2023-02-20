package main

import (
	"fmt"
	"log"
	"os"
	"primeira-aplicacao-go/app"
)

func main() {
    fmt.Println("Início")

    aplicacao := app.Gerar()
    
    if erro := aplicacao.Run(os.Args); erro != nil {
	log.Fatal(erro)
    }
  
}