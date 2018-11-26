package main

import (
	"encoding/json"
	"fmt"
)

type BlockHeaderResp struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`

	// The raw block header as a hexadecimal string.
	Result string `json:"result"`
}

func (resp *BlockHeaderResp) String() string {
	tmpl := `
	ID:      %v
	Jsonrpc: %v
	Result:  %v
	`
	return fmt.Sprintf(tmpl, resp.ID, resp.Jsonrpc, resp.Result)
}

// Return the block header at the given height.
// * The height of the block, a non-negative integer.
func (client *ElectrumxClient) GetBlockHeader(height int) (*BlockHeaderResp, error) {
	if err := client.call1(0, "blockchain.block.header", height); err != nil {
		return nil, err
	}

	resp, err := client.recv()
	if err != nil {
		return nil, err
	}

	rez := BlockHeaderResp{}
	if err := json.Unmarshal(resp, &rez); err != nil {
		return nil, err
	}

	return &rez, nil
}

type BlockHeadersResp struct {
	ID      int                `json:"id"`
	Jsonrpc string             `json:"jsonrpc"`
	Result  BlockHeadersResult `json:"result"`
}

func (resp *BlockHeadersResp) String() string {
	tmpl := `
	ID:      %v
	Jsonrpc: %v
	Count:   %v
	Hex:     %v
	Max:     %v
	`
	return fmt.Sprintf(tmpl, resp.ID, resp.Jsonrpc, resp.Result.Count, resp.Result.Hex, resp.Result.Max)
}

type BlockHeadersResult struct {
	// The number of headers returned, between zero and the number requested. If
	// the chain has not extended sufficiently far, only the available headers will be
	// returned. If more headers than max were requested at most max will be returned.
	Count int `json:"count"`

	// The binary block headers concatenated together in-order as a hexadecimal string.
	Hex string `json:"hex"`

	// The maximum number of headers the server will return in a single request.
	Max int `json:"max"`
}

// Return a concatenated chunk of block headers from the main chain.
// * start_height - the height of the first header requested, a non-negative integer.
// * count - the number of headers requested, a non-negative integer.
func (client *ElectrumxClient) GetBlockHeaders(startHeight int, count int) (*BlockHeadersResp, error) {
	if err := client.call2(0, "blockchain.block.headers", startHeight, count); err != nil {
		return nil, err
	}

	resp, err := client.recv()
	if err != nil {
		return nil, err
	}

	rez := BlockHeadersResp{}
	if err := json.Unmarshal(resp, &rez); err != nil {
		return nil, err
	}

	return &rez, nil
}

type EstimateFeeResp struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`

	// The estimated transaction fee in coin units per kilobyte, as a floating point number.
	// If the daemon does not have enough information to make an estimate, the integer
	// -1 is returned.
	Result float64 `json:"result"`
}

func (resp *EstimateFeeResp) String() string {
	tmpl := `
	ID:      %v
	Jsonrpc: %v
	Result:  %v
	`
	return fmt.Sprintf(tmpl, resp.ID, resp.Jsonrpc, resp.Result)
}

// Return the estimated transaction fee per kilobyte for a transaction to be confirmed
// within a certain number of blocks.
// * number - the number of blocks to target for confirmation.
func (client *ElectrumxClient) EstimateFee(number int) (*EstimateFeeResp, error) {
	if err := client.call1(0, "blockchain.estimatefee", number); err != nil {
		return nil, err
	}

	resp, err := client.recv()
	if err != nil {
		return nil, err
	}

	rez := EstimateFeeResp{}
	if err := json.Unmarshal(resp, &rez); err != nil {
		return nil, err
	}

	return &rez, nil
}

type RelayFeeResp struct {
	ID      int     `json:"id"`
	Jsonrpc string  `json:"jsonrpc"`
	Result  float64 `json:"result"`
}

func (resp *RelayFeeResp) String() string {
	tmpl := `
	ID:      %v
	Jsonrpc: %v
	Result:  %v
	`
	return fmt.Sprintf(tmpl, resp.ID, resp.Jsonrpc, resp.Result)
}

// Return the minimum fee a low-priority transaction must pay in order to be accepted
// to the daemon’s memory pool.
func (client *ElectrumxClient) RelayFee() (*RelayFeeResp, error) {
	if err := client.call0(0, "blockchain.relayfee"); err != nil {
		return nil, err
	}

	resp, err := client.recv()
	if err != nil {
		return nil, err
	}

	rez := RelayFeeResp{}
	if err := json.Unmarshal(resp, &rez); err != nil {
		return nil, err
	}

	return &rez, nil
}
