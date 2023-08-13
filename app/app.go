package app

import (
	"fmt"
	"log"
	"net"
	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		Name: "host",
		Value: "styvennoronha.github.io",
		
	},
}

func Gerar() *cli.App{
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPs de endereços na Internet",
			Flags: flags,
			Action: buscarips,
		},
		{
			Name: "servidor",
			Usage: "Busca de nomes de Servidor na Internet",
			Flags: flags,
			Action: buscarServidor,
		},

	}

	return app
}

func buscarips(c *cli.Context)  {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)	
	}

	for _, ip := range ips {
		log.Println(ip)
	}
}

func buscarServidor(c *cli.Context)  {
	host := c.String("host")
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)	
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}