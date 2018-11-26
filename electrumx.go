package main

import (
	"bufio"
	"fmt"
	"net"
)

const (
	defaultElectrumxServerHost = "127.0.0.1"
	defaultElectrumxServerPort = 60401
)

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

func (client *ElectrumxClient) call0(id int, method string) error {
	tmpl := `{"id":%v, "method": "%v", "params": []}` + "\n"
	raw := fmt.Sprintf(tmpl, id, method)
	_, err := client.conn.Write([]byte(raw))
	return err
}

func (client *ElectrumxClient) call1(id int, method string, params ...interface{}) error {
	tmpl := `{"id":%v, "method": "%v", "params": [%v]}` + "\n"
	allParams := append([]interface{}{id, method}, params...)
	raw := fmt.Sprintf(tmpl, allParams...)
	_, err := client.conn.Write([]byte(raw))
	return err
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