package app

import "github.com/urfave/cli"

// Gerar vai retorna a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nome de servidores na internet"
	
	return app
}