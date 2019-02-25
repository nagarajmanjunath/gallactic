package account

import (
	"math/big"
	"encoding/json"
	"fmt"

	"github.com/gallactic/gallactic/common/binary"
	"github.com/gallactic/gallactic/crypto"
	"github.com/gallactic/gallactic/errors"
	amino "github.com/tendermint/go-amino"
)

// Account structure
type Account struct {
	data accountData
}

type accountData struct {
	Address     crypto.Address  `json:"address"`
	Sequence    uint64          `json:"sequence"`
	Balance     *big.Int        `json:"balance"`
	Code        binary.HexBytes `json:"code"`
	Permissions Permissions     `json:"permissions"`
}

///---- Constructors
func NewAccount(addr crypto.Address) (*Account, error) {
	if err := addr.EnsureValid(); err != nil {
		return nil, err
	}

	if !addr.IsAccountAddress() {
		return nil, e.Errorf(e.ErrInvalidAddress, "This is not a valid acccount address: %s", addr.String())
	}

	return &Account{
		data: accountData{
			Address: addr,
		    Balance: new(big.Int),
		},
	}, nil
}

func NewContractAccount(addr crypto.Address) (*Account, error) {
	if err := addr.EnsureValid(); err != nil {
		return nil, err
	}

	if !addr.IsContractAddress() {
		return nil, e.Errorf(e.ErrInvalidAddress, "This is not a valid contract address: %s", addr.String())
	}

	return &Account{
		data: accountData{
			Address: addr,
			Balance: new(big.Int),
		},
	}, nil
}

/// For tests
func NewAccountFromSecret(secret string) *Account {
	pb, _ := crypto.GenerateKeyFromSecret(secret)
	acc, _ := NewAccount(pb.AccountAddress())
	return acc
}

func AccountFromBytes(bs []byte) (*Account, error) {
	var acc Account
	if err := acc.Decode(bs); err != nil {
		return nil, err
	}
	return &acc, nil
}

func (acc Account) Address() crypto.Address  { return acc.data.Address }
func (acc Account) Balance() *big.Int		 {return acc.data.Balance}
func (acc Account) Sequence() uint64         { return acc.data.Sequence }
func (acc Account) Code() []byte             { return acc.data.Code }
func (acc Account) Permissions() Permissions { return acc.data.Permissions }

func (acc Account) HasPermissions(perm Permissions) bool {
	return acc.data.Permissions.IsSet(perm)
}

func (acc *Account) SetBalance(bal *big.Int) error {
	acc.data.Balance = bal
	return nil
}

func (acc *Account) SubtractFromBalance(amt *big.Int) error {

	result := amt.Cmp(acc.Balance())
	if result > 0 {
		return e.Errorf(e.ErrInsufficientFunds, "Attempt to subtract %v from the balance of %s", amt, acc.Address())
	}
    acc.Balance().Sub(acc.data.Balance,amt)
	return nil
}

func (acc *Account) AddToBalance(amt *big.Int) error {
	acc.Balance().Add(acc.Balance(), amt)
	return nil
}

func (acc *Account) SetCode(code []byte) error {
	acc.data.Code = code
	return nil
}

func (acc *Account) SetSequence(seq uint64) {
	acc.data.Sequence = seq
}

func (acc *Account) IncSequence() {
	acc.data.Sequence++
}

func (acc *Account) SetPermissions(perm Permissions) error {
	acc.data.Permissions.Set(perm)
	return nil
}

func (acc *Account) UnsetPermissions(perm Permissions) error {
	acc.data.Permissions.Unset(perm)
	return nil
}

//protobuf marshal,unmarshal and size
var cdc = amino.NewCodec()

func (acc *Account) Encode() ([]byte, error) {
	return cdc.MarshalBinaryLengthPrefixed(&acc.data)
}

func (acc *Account) Decode(bs []byte) error {
	return cdc.UnmarshalBinaryLengthPrefixed(bs, &acc.data)
}

func (acc *Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(acc.data)
}

func (acc *Account) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &acc.data)
}

func AccountFromJSON(bs []byte) (*Account, error) {
	var acc Account
	if err := acc.UnmarshalJSON(bs); err != nil {
		return nil, err
	}
	return &acc, nil
}

func (acc Account) String() string {
	b, _ := acc.MarshalJSON()
	return fmt.Sprintf("Account%s", string(b))
}

func (acc *Account) Unmarshal(bs []byte) error {
	return acc.Decode(bs)
}

func (acc *Account) Marshal() ([]byte, error) {
	return acc.Encode()
}

func (acc *Account) MarshalTo(data []byte) (int, error) {
	bs, err := acc.Encode()
	if err != nil {
		return -1, err
	}
	return copy(data, bs), nil
}

func (acc *Account) Size() int {
	bs, _ := acc.Encode()
	return len(bs)
}
