package sortition

import (
	"math/big"
	"testing"

	"github.com/gallactic/gallactic/crypto"
	"github.com/stretchr/testify/assert"
)

func TestVRF(t *testing.T) {
	for i := 0; i < 100; i++ {
		pk, pv := crypto.GenerateKey(nil)
		signer := crypto.NewValidatorSigner(pv)
		m := []byte{byte(i)}
		vrf := NewVRF(signer)
		max :=(i + 1*1000)
	    m1  := big.NewInt(int64(max))
		vrf.max = m1
		index, proof := vrf.Evaluate(m)

		//fmt.Printf("%x\n", index)
		i := index.Cmp(m1)
		assert.Equal(t, i <= 0, true)
		index2, result := vrf.Verify(m, pk, proof)

		assert.Equal(t, result, true)
		assert.Equal(t, index, index2)
	}
}
