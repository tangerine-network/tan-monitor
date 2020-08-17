package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"bufio"

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
func getLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
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
		cli.StringFlag{
			Name:  "skipList",
			Value: "./skip-list.txt",
			Usage: "Path to the skip list",
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
		skipList, _ := getLines(ctx.String("skipList"))
		fmt.Println("skip list:")
		for _, skip := range skipList {
			fmt.Println(skip)
		}
		backend := monitor.NewBlockchainBackend(networkID)
		m := monitor.NewMonitor(networkID, backend, threshold)
		email := monitor.NewEmail(
			"no-reply@byzantine-lab.io",
			string(emailPassword),
			"smtp-relay.gmail.com",
			string(ccList),
			skipList,
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
