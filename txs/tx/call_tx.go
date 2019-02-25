package tx

import (
	"math/big"
	"encoding/json"
	"github.com/gallactic/gallactic/common/binary"
	"github.com/gallactic/gallactic/crypto"
	"github.com/gallactic/gallactic/errors"
)

type CallTx struct {
	data callData
}

type callData struct {
	Caller   TxInput         `json:"caller"`
	Callee   TxOutput        `json:"callee"`
	GasLimit uint64          `json:"gasLimit"`
	Data     binary.HexBytes `json:"data,omitempty"`
}

func NewCallTx(caller, callee crypto.Address, sequence uint64,data []byte, gasLimit uint64, amount *big.Int, fee *big.Int) (*CallTx, error) {

	sum := new(big.Int)
	sum.Add(amount,fee)
	return &CallTx{
		data: callData{
			Caller: TxInput{
				Address:  caller,
				Sequence: sequence,
				Amount: sum  ,
			},
			Callee: TxOutput{
				Address: callee,
				Amount:  amount,
			},
			GasLimit: gasLimit,
			Data:     data,
		},
	}, nil
}

func (tx *CallTx) Type() Type       { return TypeCall }
func (tx *CallTx) Caller() TxInput  { return tx.data.Caller }
func (tx *CallTx) Callee() TxOutput { return tx.data.Callee }
func (tx *CallTx) GasLimit() uint64 { return tx.data.GasLimit }
func (tx *CallTx) Data() []byte     { return tx.data.Data }

func (tx *CallTx) Signers() []TxInput {
	return []TxInput{tx.data.Caller}
}

func (tx *CallTx) Amount() *big.Int {
	return tx.data.Callee.Amount
}

func (tx *CallTx) Fee() *big.Int {
	var fee *big.Int
	fee =fee.Sub(tx.data.Caller.Amount,tx.data.Callee.Amount)
	return fee
}

func (tx *CallTx) EnsureValid() error {

	var result = tx.data.Callee.Amount.Cmp(tx.data.Caller.Amount)
	 if(result > 0){
		return e.Error(e.ErrInsufficientFunds)
	}

	if err := tx.data.Caller.ensureValid(); err != nil {
		return err
	}

	if !tx.data.Caller.Address.IsAccountAddress() {
		return e.Error(e.ErrInvalidAddress)
	}

	if !tx.CreateContract() {
		if err := tx.data.Callee.ensureValid(); err != nil {
			return err
		}

		if !tx.data.Callee.Address.IsContractAddress() {
			return e.Error(e.ErrInvalidAddress)
		}
	}

	return nil
}

func (tx *CallTx) CreateContract() bool {
	return tx.data.Callee.Address == crypto.Address{}
}

/// ----------
/// MARSHALING

func (tx CallTx) MarshalAmino() ([]byte, error) {
	return cdc.MarshalBinaryLengthPrefixed(tx.data)
}

func (tx *CallTx) UnmarshalAmino(bs []byte) error {
	return cdc.UnmarshalBinaryLengthPrefixed(bs, &tx.data)
}

func (tx CallTx) MarshalJSON() ([]byte, error) {
	return json.Marshal(tx.data)
}

func (tx *CallTx) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &tx.data)
}
