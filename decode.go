package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/wire"
)

func DecodeBlockHeadersResp(resp *BlockHeadersResp) ([]*wire.BlockHeader, error) {
	raw, err := hex.DecodeString(resp.Result.Hex)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer(raw)

	blockHeaders := make([]*wire.BlockHeader, 0)
	for buff.Len() != 0 {
		fmt.Println(buff.Len())

		blockHeader := wire.BlockHeader{}
		if err := blockHeader.Deserialize(buff); err != nil {
			return nil, err
		}
		blockHeaders = append(blockHeaders, &blockHeader)
	}
	return blockHeaders, nil
}