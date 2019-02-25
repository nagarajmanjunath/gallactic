package tx

import (
	"math/big"
	"encoding/json"

	"github.com/gallactic/gallactic/crypto"
	"github.com/gallactic/gallactic/errors"
)

type SendTx struct {
	data sendData
}
type sendData struct {
	Senders   []TxInput  `json:"senders"`
	Receivers []TxOutput `json:"receivers"`
}

func EmptySendTx() (*SendTx, error) {
	return &SendTx{}, nil
}

func NewSendTx(from, to crypto.Address, seq uint64, amt, fee *big.Int) (*SendTx, error) {
	tx := &SendTx{}
	sum := new(big.Int)
	sum.Add(amt,fee)
	tx.AddSender(from, seq, sum)
	tx.AddReceiver(to, amt)

	return tx, nil
}

func (tx *SendTx) Type() Type            { return TypeSend }
func (tx *SendTx) Senders() []TxInput    { return tx.data.Senders }
func (tx *SendTx) Receivers() []TxOutput { return tx.data.Receivers }

func (tx *SendTx) Signers() []TxInput {
	return tx.data.Senders
}

func (tx *SendTx) Amount() *big.Int {
	return tx.outAmount()
}

func (tx *SendTx) Fee() *big.Int {
	var fee *big.Int
	fee =fee.Sub(tx.inAmount(),tx.outAmount())
	return fee
}

func (tx *SendTx) AddSender(addr crypto.Address, seq uint64, amt *big.Int) {
	tx.data.Senders = append(tx.data.Senders, TxInput{
		Address:  addr,
		Amount:   amt,
		Sequence: seq,
	})
}

func (tx *SendTx) AddReceiver(addr crypto.Address, amt *big.Int) {
	tx.data.Receivers = append(tx.data.Receivers, TxOutput{
		Address: addr,
		Amount:  amt,
	})
}

func (tx *SendTx) inAmount() *big.Int {
	inAmt := big.NewInt(0)
	for _, in := range tx.data.Senders {

		inAmt.Add(in.Amount,inAmt)
	}
	return inAmt
}

func (tx *SendTx) outAmount() *big.Int {
	outAmt := big.NewInt(0)
	for _, out := range tx.data.Receivers {
		outAmt.Add(out.Amount,outAmt)
	}
	return outAmt
}

func (tx *SendTx) EnsureValid() error {

	var result = tx.outAmount().Cmp(tx.inAmount())
	if(result > 0){
	   return e.Error(e.ErrInsufficientFunds)
   }

	for _, in := range tx.data.Senders {
		if err := in.ensureValid(); err != nil {
			return err
		}

		if !in.Address.IsAccountAddress() {
			return e.Error(e.ErrInvalidAddress)
		}
	}

	for _, out := range tx.data.Receivers {
		if err := out.ensureValid(); err != nil {
			return err
		}

		if !out.Address.IsAccountAddress() {
			return e.Error(e.ErrInvalidAddress)
		}
	}

	return nil
}

/// ----------
/// MARSHALING

func (tx SendTx) MarshalAmino() ([]byte, error) {
	return cdc.MarshalBinaryLengthPrefixed(tx.data)
}

func (tx *SendTx) UnmarshalAmino(bs []byte) error {
	return cdc.UnmarshalBinaryLengthPrefixed(bs, &tx.data)
}

func (tx SendTx) MarshalJSON() ([]byte, error) {
	return json.Marshal(tx.data)
}

func (tx *SendTx) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &tx.data)
}
