	//comando.exe servidor --host example.com.br Buscar o DNS
	//comando.exe ip --host example.com.br Buscar o IP
	//go run main servidor --host example.com.br Buscar o DNS
	//go run main ip --host example.com.br Buscar o IP
package main

import (
	"comando/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Inicio")

	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}


}
