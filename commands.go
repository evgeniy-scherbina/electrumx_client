package main

import (
	"fmt"
	"github.com/urfave/cli"
)

var getBlockHeaderCommand = cli.Command{
	Name:   "getblockheader",
	Usage:  "Return the block header at the given height.",
	Action: getBlockHeader,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "height",
			Usage: "The height of the block, a non-negative integer.",
		},
		cli.BoolFlag{
			Name: "verbose",
		},
	},
}

func getBlockHeader(ctx *cli.Context) error {
	height := ctx.Int("height")
	verbose := ctx.Bool("verbose")

	if height == 0 {
		return fmt.Errorf("`height` flag must be set")
	}

	client := NewElectrumxClient(defaultElectrumxServerHost, defaultElectrumxServerPort)
	if err := client.Dial(); err != nil {
		return err
	}

	resp, err := client.GetBlockHeader(height)
	if err != nil {
		return err
	}

	if !verbose {
		fmt.Println(resp)
		return nil
	}

	bh, err := DecodeBlockHeaderResp(resp)
	if err != nil {
		return err
	}

	fmt.Println(BlockHeaderDescription(bh))
	return nil
}

var getBlockHeadersCommand = cli.Command{
	Name:   "getblockheaders",
	Usage:  "Return a concatenated chunk of block headers from the main chain.",
	Action: getBlockHeaders,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "start_height",
			Usage: "The height of the first header requested, a non-negative integer.",
		},
		cli.StringFlag{
			Name:  "count",
			Usage: "The number of headers requested, a non-negative integer.",
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

	if startHeight == 0 || count == 0 {
		return fmt.Errorf("both `start_height` and `count` flags must be set")
	}

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

var estimateFeeCommand = cli.Command{
	Name: "estimatefee",
	Usage: "Return the estimated transaction fee per kilobyte for a transaction to be confirmed" +
		"within a certain number of blocks.",
	Action: estimateFee,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "number",
			Usage: "The number of blocks to target for confirmation.",
		},
	},
}

func estimateFee(ctx *cli.Context) error {
	number := ctx.Int("number")
	if !ctx.IsSet("number") {
		return fmt.Errorf("`number` flag must be set")
	}

	client := NewElectrumxClient(defaultElectrumxServerHost, defaultElectrumxServerPort)
	if err := client.Dial(); err != nil {
		return err
	}

	resp, err := client.EstimateFee(number)
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}

var relayFeeCommand = cli.Command{
	Name: "relayfee",
	Action: relayFee,
}

func relayFee(ctx *cli.Context) error {
	client := NewElectrumxClient(defaultElectrumxServerHost, defaultElectrumxServerPort)
	if err := client.Dial(); err != nil {
		return err
	}

	resp, err := client.RelayFee()
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}