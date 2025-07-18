package app

import (
	"fmt"
	// "log"
	"net"

	"github.com/urfave/cli/v2"
)

// Gerar retorna a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "felipepfranco.com.br",
			Usage: "Host a ser buscado",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços da Internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca nomes de servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) error {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		return erro // ou log.Fatal(erro), mas retornar erro é mais indicado aqui
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}

func buscarServidores(c *cli.Context) error {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		return erro // ou log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
	return nil
}
