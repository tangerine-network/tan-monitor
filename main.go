package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/urfave/cli"
	"github.com/tangerine-network/tan-monitor/monitor"
)

var app *cli.App

func init() {
	app = cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Usage = "Tangerine Newtwork Monitor"
	app.Commands = []cli.Command{
		commandStart,
	}
}

var commandStart = cli.Command{
	Name:      "start",
	Usage:     "Start monitor job",
	ArgsUsage: "<networkID> <ethThreshold>",
	Description: `Start network monitor with <networkID> and <ethThreshold>
   <ethThreshold>'s unit is ETH`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "emailPassword",
			Value: "./email.password",
			Usage: "Path to the email sender password",
		},
		cli.StringFlag{
			Name:  "ccList",
			Value: "./cc.txt",
			Usage: "Path to the cc list",
		},
	},
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) != 2 {
			log.Fatalf("invalid argument count")
		}
		networkID, err := strconv.Atoi(ctx.Args()[0])
		if err != nil {
			panic(err)
		}
		threshold := ctx.Args()[1]
		emailPassword, err := ioutil.ReadFile(ctx.String("emailPassword"))
		if err != nil {
			panic(err)
		}
		ccList, _ := ioutil.ReadFile(ctx.String("ccList"))
		backend := monitor.NewBlockchainBackend(networkID)
		m := monitor.NewMonitor(networkID, backend, threshold)
		email := monitor.NewEmail(
			"no-reply@byzantine-lab.io",
			string(emailPassword),
			"smtp-relay.gmail.com",
			string(ccList),
		)
		m.Register(email)
		m.Run()
		return nil
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Println("app.Run error:", err)
		os.Exit(1)
	}
}
