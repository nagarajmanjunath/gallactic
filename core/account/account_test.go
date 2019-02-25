package account

import (
	"encoding/json"
	"fmt"
	"testing"
	"math/big"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshaling(t *testing.T) {
	acc1 := NewAccountFromSecret("Secret")
	newBigInt := new(big.Int)
	newBigInt.SetString("1000000000000000000", 10)
	acc1.AddToBalance(newBigInt)
	acc1.SetPermissions(0x77)
	acc1.IncSequence()
	acc1.SetCode([]byte{60, 23, 45})

	/// test amino encoding/decoding
	bs, err := acc1.Encode()
	require.NoError(t, err)
	acc2 := new(Account)
	err = acc2.Decode(bs)
	acc2.AddToBalance(newBigInt)
	fmt.Println("Account2Balance",acc2.Balance())
	require.NoError(t, err)
	assert.Equal(t, acc1, acc2)

	acc3, err := AccountFromBytes(bs)
	acc3.AddToBalance(newBigInt)
	require.NoError(t, err)
	assert.Equal(t, acc2, acc3)

	// /// test json marshaing/unmarshaling
	js, err := json.Marshal(acc1)
	require.NoError(t, err)
	fmt.Println(string(js))
	acc4 := new(Account)
	require.NoError(t, json.Unmarshal(js, acc4))
     assert.Equal(t, acc3, acc4)

	// /// should fail
	acc5, err := AccountFromBytes([]byte("asdfghjkl"))
	require.Error(t, err)
	assert.Nil(t, acc5)

}
