package sputnikvm

import (
	"fmt"
	// "math/big"

	// "github.com/gallactic/gallactic/core/account"

	// ETCCommon "github.com/ethereumproject/go-ethereum/common"
	// "github.com/gallactic/gallactic/core/evm"
	// "github.com/gallactic/sputnikvm-ffi/go/sputnikvm"

	// tmRPC "github.com/tendermint/tendermint/rpc/core"
)

func Execute(adapter GallacticAdapter) Output {
	fmt.Printf("SputnikVM called.\n")

	var out Output

// 	transaction := sputnikvm.Transaction{
// 		Caller:   adapter.callerAddress(),
// 		Address:  adapter.calleeAddress(),
// 		GasPrice: adapter.GetGasPrice(),
// 		GasLimit: adapter.GetGasLimit(),
// 		Value:    adapter.GetAmount(),
// 		Input:    adapter.GetData(),
// 		Nonce:    adapter.GetNonce(),
// 	}

// 	beneficiary := ETCCommon.HexToAddress("ffffffffffffffffffffffffffffffffffffffff")
// 	header := sputnikvm.HeaderParams{
// 		Timestamp: adapter.TimeStamp(),
// 		Number:    adapter.LastBlockNumber(),
// 		GasLimit:  adapter.GetGasLimit(),
// 		// Required by PoW block info
// 		Difficulty:  new(big.Int).SetUint64(0),
// 		Beneficiary: beneficiary,
// 	}

// 	vm := sputnikvm.NewGallactic(&transaction, &header)

// Loop:
// 	for {
// 		require := vm.Fire()

// 		switch require.Typ() {

// 		case sputnikvm.RequireNone:
// 			break Loop

// 		case sputnikvm.RequireAccount:
// 			acc := adapter.getAccount(require.Address())
// 			if acc != nil {
// 				vm.CommitAccount(require.Address(), new(big.Int).SetUint64(acc.Sequence()),
// 					(acc.Balance()), acc.Code())
// 			} else {
// 				vm.CommitNonexist(require.Address())
// 			}

// 		case sputnikvm.RequireAccountCode:
// 			acc := adapter.getAccount(require.Address())
// 			if acc != nil {
// 				vm.CommitAccountCode(require.Address(), acc.Code())
// 			} else {
// 				vm.CommitNonexist(require.Address())
// 			}

// 		case sputnikvm.RequireAccountStorage:
// 			storage := adapter.getStorage(require.Address(), require.StorageKey())
// 			vm.CommitAccountStorage(require.Address(), require.StorageKey(), storage)

// 		case sputnikvm.RequireBlockhash:
// 			var blockHash ETCCommon.Hash

// 			reqblockNumber := require.BlockNumber().Int64()
// 			block, err := tmRPC.Block(&reqblockNumber)
// 			if err == nil {
// 				hash := block.Block.Hash().Bytes()
// 				blockHash.SetBytes(hash)
// 			}
// 			vm.CommitBlockhash(require.BlockNumber(), blockHash)

// 		default:
// 			/// The cache is irreversible, during delivering call transaction
// 			/// Better panic in case of unexpected error happens
// 			/// rather than changing the state which the tx is not delivered.
// 			panic("unreachable")
// 		}
// 	}

// 	// HACKING SPUTNIKVM:
// 	// If contract A creates contract B, the byte code of B will always be less than A.
// 	// AccountChange are not always synchronize.
// 	var contractCodeLength = 0

// 	changedAccs := vm.AccountChanges()
// 	accLen := len(changedAccs)

// 	for i := 0; i < accLen; i++ {
// 		changedAcc := changedAccs[i]
// 		// We don't support beneficiary account
// 		if beneficiary == changedAcc.Address() {
// 			continue
// 		}

// 		switch changedAcc.Typ() {

// 		case sputnikvm.AccountChangeIncreaseBalance:
// 			//Increase Balance
// 			amount := changedAcc.ChangedAmount()
// 			adapter.addBalance(changedAcc.Address(), amount)

// 		case sputnikvm.AccountChangeDecreaseBalance:
// 			//Decrease Balance
// 			amount := changedAcc.ChangedAmount()
// 			adapter.subBalance(changedAcc.Address(), amount)

// 		case sputnikvm.AccountChangeRemoved:
// 			//Removing Account
// 			adapter.removeAccount(changedAcc.Address())

// 		case sputnikvm.AccountChangeFull:
// 			acc := adapter.getAccount(changedAcc.Address())

// 			changeStorage := changedAcc.ChangedStorage()
// 			if len(changeStorage) > 0 {
// 				for i := 0; i < len(changeStorage); i++ {
// 					key := changeStorage[i].Key
// 					value := changeStorage[i].Value
// 					adapter.updateStorage(changedAcc.Address(), key, value)
// 				}
// 			}
// 			acc.SetBalance(changedAcc.Balance())
// 			acc.SetSequence(changedAcc.Nonce().Uint64())
// 			// After deploying the contract, code can't be changed
// 			// https://github.com/ethereumproject/go-ethereum/issues/696
// 			// acc.SetCode(changedAcc.Code())

// 			adapter.updateAccount(acc)

// 		case sputnikvm.AccountChangeCreate:
// 			var acc *account.Account
// 			if len(changedAcc.Code()) == 0 {
// 				acc = adapter.createAccount(changedAcc.Address())

// 				acc.SetBalance(changedAcc.Balance())
// 				acc.SetSequence(changedAcc.Nonce().Uint64())
// 			} else {
// 				acc = adapter.createContractAccount(changedAcc.Address())

// 				changeStorage := changedAcc.Storage()
// 				if len(changeStorage) > 0 {
// 					for i := 0; i < len(changeStorage); i++ {
// 						key := changeStorage[i].Key
// 						value := changeStorage[i].Value
// 						adapter.updateStorage(changedAcc.Address(), key, value)
// 					}
// 				}
// 				acc.SetBalance(changedAcc.Balance())
// 				acc.SetSequence(changedAcc.Nonce().Uint64())
// 				acc.SetCode(changedAcc.Code())

// 				if contractCodeLength <= len(acc.Code()) {
// 					addr := acc.Address()
// 					out.ContractAddress = &addr
// 					//
// 					contractCodeLength = len(acc.Code())
// 				}
// 			}

// 			adapter.updateAccount(acc)

// 		default:
// 			panic("unreachable")
// 		}
// 	}

// 	out.UsedGas = vm.UsedGas().Uint64()

// 	if vm.Failed() {
// 		fmt.Println("SputnikVm Failed")
// 		out.Failed = true
// 	} else {
// 		out.Failed = false
// 	}

// 	//Extract logs and events
// 	var logs []evm.Log
// 	for _, log := range vm.Logs() {
// 		logs = append(logs, adapter.ConvertLog(log))
// 	}

// 	out.Logs = logs
// 	out.UsedGas = vm.UsedGas().Uint64()
// 	out.Output = make([]uint8, vm.OutLen())
// 	copy(out.Output, vm.Output())

// 	vm.Free()

	return out
}
