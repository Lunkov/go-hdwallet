package hdwallet

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func init() {
	coins[ECOS] = newECOS
}

type ecos struct {
	name   string
	symbol string
	key    *Key

	// ECOS token
	contract string
}

func newECOS(key *Key) Wallet {
	return &ecos{
		name:   "EcoSphere",
		symbol: "ECOS",
		key:    key,
	}
}

func (c *ecos) GetType() uint32 {
	return c.key.opt.CoinType
}

func (c *ecos) GetName() string {
	return c.name
}

func (c *ecos) GetSymbol() string {
	return c.symbol
}

func (c *ecos) GetKey() *Key {
	return c.key
}

func (c *ecos) GetAddress() (string, error) {
	return crypto.PubkeyToAddress(*c.key.PublicECDSA).Hex(), nil
}
