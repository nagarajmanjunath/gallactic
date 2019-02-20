package tx

import (
	"math/big"
	"github.com/gallactic/gallactic/crypto"
)

type TxOutput struct {
	Address crypto.Address `json:"address"`
	Amount  *big.Int         `json:"amount"`
}

func (out *TxOutput) ensureValid() error {
	return out.Address.EnsureValid()
}
