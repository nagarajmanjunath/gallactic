package sortition

import (
	"math/big"
	"github.com/gallactic/gallactic/core/state"
	"github.com/gallactic/gallactic/core/validator"
	"github.com/gallactic/gallactic/crypto"
	"github.com/gallactic/gallactic/txs"
	"github.com/gallactic/gallactic/txs/tx"

	"github.com/hyperledger/burrow/logging"
	tmRPC "github.com/tendermint/tendermint/rpc/core"
)

type Sortition struct {
	//transactor ITransactor
	state        *state.State
	signer       crypto.Signer
	vrf          VRF
	chainID      string
	sortitionFee *big.Int
	logger       *logging.Logger
}

func NewSortition(state *state.State, signer crypto.Signer, chainID string, logger *logging.Logger) *Sortition {
	return &Sortition{
		signer:  signer,
		state:   state,
		chainID: chainID,
		vrf:     NewVRF(signer),
		logger:  logger,
	}
}

// Evaluate return the vrf for self choosing to be a validator
func (s *Sortition) Evaluate(blockHeight uint64, blockHash []byte) {
	totalStake, valStake := s.getTotalStake(s.signer.Address())

	s.vrf.SetMax(totalStake)
	index, proof := s.vrf.Evaluate(blockHash)
     iresult := index.Cmp(valStake)
	if iresult < 0 {
		s.logger.InfoMsg("This validator is chosen to be in set at height %v", blockHeight)

		/// TODO: better way????
		val, err := s.state.GetValidator(s.Address())
		if err != nil {
			return
		}
		tx, _ := tx.NewSortitionTx(
			s.signer.Address(),
			blockHeight,
			val.Sequence()+1,
			s.sortitionFee,
			index,
			proof)

		txEnv := txs.Enclose(s.chainID, tx)
		err = txEnv.Sign(s.signer)
		if err != nil {
			return
		}

		// TODO:: better way?????
		codec := txs.NewAminoCodec()
		bs, err := codec.MarshalBinaryLengthPrefixed(txEnv)
		if err != nil {
			return
		}

		res, err := tmRPC.BroadcastTxAsync(bs)
		if err != nil {
			return
		}

		if res != nil {
			/// TODO: log result
		}
	}
}

func (s *Sortition) Verify(blockHash []byte, pb crypto.PublicKey, index *big.Int, proof []byte) bool {

	totalStake, valStake := s.getTotalStake(pb.ValidatorAddress())

	// Note: totalStake can be changed by time on verifying
	// So we calculate the index again

	s.vrf.SetMax(totalStake)
	index2, result := s.vrf.Verify(blockHash, pb, proof)
	if !result {
		return false
	}
	var iresult = index2.Cmp(valStake)
	if (iresult < 0){
		return true
	}else {
	return false
   }


	// return index2 < valStake
}

func (s *Sortition) Address() crypto.Address {
	return s.signer.Address()
}

func (s *Sortition) getTotalStake(addr crypto.Address) (totalStake *big.Int, validatorStake *big.Int) {
	totalStake = big.NewInt(0)
	validatorStake = big.NewInt(0)
	s.state.IterateValidators(func(validator *validator.Validator) (stop bool) {
		totalStake.Add(totalStake,validator.Stake())
		if addr == validator.Address() {
			validatorStake = validator.Stake()
		}
		return false
	})

	return
}
