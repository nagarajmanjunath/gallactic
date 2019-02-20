package executors

import (
	"fmt"

	"github.com/gallactic/gallactic/core/account"
	"github.com/gallactic/gallactic/core/account/permission"
	"github.com/gallactic/gallactic/core/blockchain"
	"github.com/gallactic/gallactic/core/evm/sputnikvm"
	"github.com/gallactic/gallactic/core/state"
	e "github.com/gallactic/gallactic/errors"
	"github.com/gallactic/gallactic/txs"
	"github.com/gallactic/gallactic/txs/tx"

	"github.com/hyperledger/burrow/logging"
)

type CallContext struct {
	Committing bool
	BC         *blockchain.Blockchain
	Cache      *state.Cache
	Logger     *logging.Logger
}

func (ctx *CallContext) Execute(txEnv *txs.Envelope, txRec *txs.Receipt) error {
	tx, ok := txEnv.Tx.(*tx.CallTx)
	if !ok {
		return e.Error(e.ErrInvalidTxType)
	}

	caller, err := getInputAccount(ctx.Cache, tx.Caller(), permission.Call)
	if err != nil {
		return err
	}

	var callee *account.Account
	if tx.CreateContract() {
		if !ctx.Cache.HasPermissions(caller, permission.CreateContract) {
			return e.Errorf(e.ErrPermissionDenied, "%s has %s but needs %s", caller.Address(), caller.Permissions(), permission.CreateContract)
		}

		// In case of create contract we must pass nil as callee
		// sputnik vm will create the account and returns the code
		callee = nil

		ctx.Logger.TraceMsg("Creating new contract", "init_code", tx.Data())
	} else {
		callee, _ = ctx.Cache.GetAccount(tx.Callee().Address)
		if callee == nil {
			return e.Errorf(e.ErrInvalidAddress, "attempt to call a non-existing account: %s", tx.Callee().Address)
		}
	}

	if ctx.Committing {
		ret := ctx.Deliver(tx, caller, callee)
		// we get the latest changes after sputnik modified caller account info
		// scenario: sputnik modified caller balance, then at this line, we need to get the latest info
		caller, _ = ctx.Cache.GetAccount(caller.Address())

		//Here we can acquire sputnikVM result
		if ret.Failed {
			txRec.Status = txs.Failed
		} else {
			txRec.Status = txs.Ok
		}
		txRec.GasUsed = ret.UsedGas
		txRec.GasWanted = tx.GasLimit()
		txRec.Logs = ret.Logs
		txRec.Output = ret.Output
		txRec.ContractAddress = ret.ContractAddress
	}

	err = caller.SubtractFromBalance(tx.Fee())
	if err != nil {
		return err
	}

	/// Update state cache
	ctx.Cache.UpdateAccount(caller)

	return nil
}

func (ctx *CallContext) Deliver(tx *tx.CallTx, caller, callee *account.Account) (ret sputnikvm.Output) {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("recovered from panic on calling sputnikVM: %v", r)
			ret.Failed = true
			ret.Output = []byte(err.Error())
		}
	}()

	adapter := sputnikvm.GallacticAdapter{ctx.BC, ctx.Cache, caller,
		callee, tx.GasLimit(), tx.Amount(), tx.Data(), caller.Sequence()}

	ret = sputnikvm.Execute(adapter)

	return ret
}
