package hdwallet

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func init() {
	coins[EVER] = newEVER
}

type everscale struct {
	name   string
	symbol string
	key    *Key

	// Everscale token
	contract string
}

func newEVER(key *Key) Wallet {
	return &everscale{
		name:   "Everscale",
		symbol: "EVER",
		key:    key,
	}
}

func (c *everscale) GetType() uint32 {
	return c.key.opt.CoinType
}

func (c *everscale) GetName() string {
	return c.name
}

func (c *everscale) GetSymbol() string {
	return c.symbol
}

func (c *everscale) GetKey() *Key {
	return c.key
}

func (c *everscale) GetAddress() (string, error) {
	return crypto.PubkeyToAddress(*c.key.PublicECDSA).Hex(), nil
}
