package main

import (
	"github.com/urfave/cli"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "electrumx_client"

	app.Commands = []cli.Command{
		getBlockHeaderCommand,
		getBlockHeadersCommand,
		estimateFeeCommand,
		relayFeeCommand,
		decodeAddressCommand,
		scriptHashGetBalanceCommand,
		scriptHashGetHistoryCommand,
		scriptHashGetMempoolCommand,
	}

	if err := app.Run(os.Args); err != nil {
		checkErr(err)
	}
}
