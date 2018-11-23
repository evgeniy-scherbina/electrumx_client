package main

import (
	"fmt"
	"github.com/urfave/cli"
)

var getBlockHeadersCommand = cli.Command{
	Name:   "getblockheaders",
	Action: getBlockHeaders,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "start_height",
		},
		cli.StringFlag{
			Name: "count",
		},
		cli.BoolFlag{
			Name: "verbose",
		},
	},
}

func getBlockHeaders(ctx *cli.Context) error {
	startHeight := ctx.Int("start_height")
	count := ctx.Int("count")
	verbose := ctx.Bool("verbose")

	client := NewElectrumxClient(defaultElectrumxServerHost, defaultElectrumxServerPort)
	if err := client.Dial(); err != nil {
		return err
	}

	resp, err := client.GetBlockHeaders(startHeight, count)
	if err != nil {
		return err
	}

	if !verbose {
		fmt.Println(resp)
		return nil
	}

	bh, err := DecodeBlockHeadersResp(resp)
	if err != nil {
		return err
	}

	fmt.Println(BlockHeadersDescription(bh))
	return nil
}