package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
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
	Usage: "Return the minimum fee a low-priority transaction must pay in order to be accepted" +
		"to the daemon’s memory pool.",
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

var decodeAddressCommand = cli.Command{
	Name: "decodeaddress",
	Usage: "see details https://electrumx.readthedocs.io/en/latest/protocol-basics.html#script-hashes",
	Action: decodeAddress,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "address",
		},
	},
}

func decodeAddress(ctx *cli.Context) error {
	encodedAddress := ctx.String("address")

	reversed, err := decodeAddressHelper(encodedAddress, &chaincfg.MainNetParams)
	if err != nil {
		return err
	}

	fmt.Println(hex.EncodeToString(reversed))
	return nil
}

func decodeAddressHelper(encodedAddress string, params *chaincfg.Params) ([]byte, error) {
	address, err := btcutil.DecodeAddress(encodedAddress, params)
	if err != nil {
		return nil, err
	}

	script, err := txscript.PayToAddrScript(address)
	if err != nil {
		return nil, err
	}

	scriptHash := sha256.Sum256(script)
	if err != nil {
		return nil, err
	}

	reverse := func(arr []byte) []byte {
		for i := 0; i < len(arr) / 2; i++ {
			arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
		}
		return arr
	}
	reversed := reverse(scriptHash[:])
	return reversed, nil
}

// ScriptHashGetBalance
var scriptHashGetBalanceCommand = cli.Command{
	Name: "scripthashgetbalance",
	Action: scriptHashGetBalance,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "address",
		},
	},
}

func scriptHashGetBalance(ctx *cli.Context) error {
	encodedAddress := ctx.String("address")

	reversed, err := decodeAddressHelper(encodedAddress, &chaincfg.MainNetParams)
	if err != nil {
		return err
	}

	client := NewElectrumxClient(defaultElectrumxServerHost, defaultElectrumxServerPort)
	if err := client.Dial(); err != nil {
		return err
	}

	resp, err := client.ScriptHashGetBalance(reversed)
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}