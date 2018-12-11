package ecc

import (
	"fmt"
	"github.com/eosforce/goeosforce/btcsuite/btcd/btcec"
//	"github.com/eosforce/goeosforce/btcsuite/btcutil/base58"

)

type innerK1PublicKey struct {
}

func (p *innerK1PublicKey) key(content []byte) (*btcec.PublicKey, error) {
	key, err := btcec.ParsePubKey(content, btcec.S256())
	if err != nil {
		return nil, fmt.Errorf("parsePubKey: %s", err)
	}

	return key, nil
}

func (p *innerK1PublicKey) prefix() string {
	return PublicKeyPrefixCompat
}
