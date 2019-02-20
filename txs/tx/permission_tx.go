package tx

import (
	"math/big"
	"encoding/json"
	"fmt"

	"github.com/gallactic/gallactic/core/account"
	"github.com/gallactic/gallactic/crypto"
	"github.com/gallactic/gallactic/errors"
)

type PermissionsTx struct {
	data permissionsData
}
type permissionsData struct {
	Modifier    TxInput             `json:"modifier"`
	Modified    TxOutput            `json:"modified"`
	Permissions account.Permissions `json:"permissions"`
	Set         bool                `json:"set"`
}

func NewPermissionsTx(modifier, modified crypto.Address, perm account.Permissions, set bool, seq uint64, fee *big.Int) (*PermissionsTx, error) {
    var amt = big.NewInt(0)
	return &PermissionsTx{
		data: permissionsData{
			Modifier: TxInput{
				Address:  modifier,
				Sequence: seq,
				Amount:   fee,
			},
			Modified: TxOutput{
				Address: modified,
				Amount:  amt,
			},

			Permissions: perm,
			Set:         set,
		},
	}, nil
}

func (tx *PermissionsTx) Type() Type                       { return TypePermissions }
func (tx *PermissionsTx) Modifier() TxInput                { return tx.data.Modifier }
func (tx *PermissionsTx) Modified() TxOutput               { return tx.data.Modified }
func (tx *PermissionsTx) Permissions() account.Permissions { return tx.data.Permissions }
func (tx *PermissionsTx) Set() bool                        { return tx.data.Set }

func (tx *PermissionsTx) Signers() []TxInput {
	return []TxInput{tx.data.Modifier}
}

func (tx *PermissionsTx) Amount() *big.Int {
	return big.NewInt(0)
}

func (tx *PermissionsTx) Fee() *big.Int {
	return tx.data.Modifier.Amount
}

func (tx *PermissionsTx) EnsureValid() error {
	/// Just modifying permission, not transferring money

	var amt = big.NewInt(0)
	var result = tx.data.Modified.Amount.Cmp(amt)
	if(result == 0){
	   return e.Error(e.ErrInsufficientFunds)
   }

	if err := tx.data.Modifier.ensureValid(); err != nil {
		return err
	}

	if err := tx.data.Modified.ensureValid(); err != nil {
		return err
	}

	if tx.data.Modified.Address == crypto.GlobalAddress {
		return fmt.Errorf("You can not change global account's permission")
	}

	return nil
}

/// ----------
/// MARSHALING

func (tx PermissionsTx) MarshalAmino() ([]byte, error) {
	return cdc.MarshalBinaryLengthPrefixed(tx.data)
}

func (tx *PermissionsTx) UnmarshalAmino(bs []byte) error {
	return cdc.UnmarshalBinaryLengthPrefixed(bs, &tx.data)
}

func (tx PermissionsTx) MarshalJSON() ([]byte, error) {
	return json.Marshal(tx.data)
}

func (tx *PermissionsTx) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &tx.data)
}
