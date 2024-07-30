package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

const usage = `
   ,d88b.d88b,			Orangedocker!!!
  d8'88bd8'8b'8d			It is a simple container runtime implementation.
 d8'   88bd8'   88			The purpose of this project is to learn how docker works and how to write a docker by ourselves Learn it.
88       88     88			If you has any problem when using it, please contact me at:
88     . 88     88			https://github.com/OrangeDou
88    . .88     88
d8'   ...88...  d8'
  d8'    88   8d'
   ,88888888888ba,
   d8'       8d'`

func main() {
	app := cli.NewApp() // create new cli app
	app.Name = "Orangedocker"
	app.Usage = usage

	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
