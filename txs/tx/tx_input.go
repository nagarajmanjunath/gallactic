package tx

import (
	"math/big"
	"github.com/gallactic/gallactic/crypto"
)

type TxInput struct {
	Address  crypto.Address `json:"address"`
	Amount   *big.Int        `json:"amount"`
	Sequence uint64        `json:"sequence"`
}

func (in *TxInput) ensureValid() error {
	return in.Address.EnsureValid()
}
