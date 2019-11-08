package main

import (
	"fmt"
	"github.com/FLYyyyy2016/goTools/api"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "my api cli"
	app.Version = "v 0.0.1"
	app.Action = func(c *cli.Context) {
		weather, err := api.QueryByCity("北京")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(weather)
	}
	app.Run(os.Args)
}
