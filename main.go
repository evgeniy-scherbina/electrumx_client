package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"net"
	"os"
)

const (
	defaultElectrumxServerHost = "127.0.0.1"
	defaultElectrumxServerPort = 60401
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type ElectrumxClient struct {
	host string
	port int
	conn net.Conn
}

func NewElectrumxClient(host string, port int) *ElectrumxClient {
	return &ElectrumxClient{
		host: host,
		port: port,
	}
}

func (client *ElectrumxClient) Dial() error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", client.host, client.port))
	if err != nil {
		return err
	}
	client.conn = conn
	return nil
}

func (client *ElectrumxClient) call2(id int, method string, params ...interface{}) error {
	tmpl := `{"id":%v, "method": "%v", "params": [%v, %v]}` + "\n"
	allParams := append([]interface{}{id, method}, params...)
	raw := fmt.Sprintf(tmpl, allParams...)
	_, err := client.conn.Write([]byte(raw))
	return err
}

func (client *ElectrumxClient) recv() ([]byte, error) {
	r := bufio.NewReader(client.conn)
	return r.ReadBytes('\n')
}

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

func main() {
	app := cli.NewApp()
	app.Name = "electrumx_client"

	app.Commands = []cli.Command{
		getBlockHeadersCommand,
	}

	if err := app.Run(os.Args); err != nil {
		checkErr(err)
	}
}
