package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/wire"
)

func BlockHeaderDescription(blockHeader *wire.BlockHeader) string {
	tmpl := `
	Version    %v
	PrevBlock  %v
	MerkleRoot %v
	Timestamp  %v
	Bits       %v
	Nonce      %v
	`
	return fmt.Sprintf(
		tmpl,
		blockHeader.Version,
		hex.EncodeToString(blockHeader.PrevBlock[:]),
		hex.EncodeToString(blockHeader.MerkleRoot[:]),
		blockHeader.Timestamp,
		blockHeader.Bits,
		blockHeader.Nonce,
	)
}

func BlockHeadersDescription(blockHeaders []*wire.BlockHeader) string {
	repr := ""
	for _, blockHeader := range blockHeaders {
		repr += BlockHeaderDescription(blockHeader)
	}
	return repr
}