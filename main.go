package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

var pizza = []string{"Enjoy your pizza with delicious"}

func info() {
	app.Name = "Pizza test CLI"
	app.Author = "Tanohzana"
	app.Usage = "A simple CLI that helps you add ingredients to your pizza"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "peppers",
			Aliases: []string{"p"},
			Usage:   "Add peppers to your pizza",
			Action: func(c *cli.Context) {
				pe := "peppers"
				peppers := append(pizza, pe)
				m := strings.Join(peppers, " ")
				fmt.Println(m)
			},
		},
		{
			Name:    "pineapple",
			Aliases: []string{"pa"},
			Usage:   "Add pineapple to your pizza",
			Action: func(c *cli.Context) {
				pa := "pineapple"
				pineapple := append(pizza, pa)
				m := strings.Join(pineapple, " ")
				fmt.Println(m)
			},
		},
		{
			Name:    "cheese",
			Aliases: []string{"c"},
			Usage:   "Add cheese to your pizza",
			Action: func(c *cli.Context) {
				ch := "cheese"
				cheese := append(pizza, ch)
				m := strings.Join(cheese, " ")
				fmt.Println(m)
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
