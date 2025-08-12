package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	keepertest "github.com/liftedinit/ghostcloud/testutil/keeper"
	"github.com/liftedinit/ghostcloud/testutil/sample"
	"github.com/liftedinit/ghostcloud/x/ghostcloud/keeper"
	"github.com/liftedinit/ghostcloud/x/ghostcloud/types"
)

const (
	OpWeightMsgCreateDeployment      = "op_weight_msg_ghostcloud_create_deployment" // nolint: gosec
	OpWeightMsgRemoveDeployment      = "op_weight_msg_ghostcloud_remove_deployment" // nolint: gosec
	OpWeightMsgUpdateDeployment      = "op_weight_msg_ghostcloud_update_deployment" // nolint: gosec
	DefaultWeightMsgCreateDeployment = 100
	DefaultWeightMsgRemoveDeployment = 20
	DefaultWeightMsgUpdateDeployment = 80
)

// WeightedOperations returns the all the gov module operations with their respective weights.
func WeightedOperations(appParams simtypes.AppParams,
	_ codec.JSONCodec,
	txGen client.TxConfig,
	k keeper.Keeper,
) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateDeployment int
	appParams.GetOrGenerate(OpWeightMsgCreateDeployment, &weightMsgCreateDeployment, nil, func(_ *rand.Rand) {
		weightMsgCreateDeployment = DefaultWeightMsgCreateDeployment
	})

	var weightMsgRemoveDeployment int
	appParams.GetOrGenerate(OpWeightMsgRemoveDeployment, &weightMsgRemoveDeployment, nil, func(_ *rand.Rand) {
		weightMsgRemoveDeployment = DefaultWeightMsgRemoveDeployment
	})

	var weightMsgUpdateDeployment int
	appParams.GetOrGenerate(OpWeightMsgUpdateDeployment, &weightMsgUpdateDeployment, nil, func(_ *rand.Rand) {
		weightMsgUpdateDeployment = DefaultWeightMsgUpdateDeployment
	})

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDeployment,
		SimulateMsgCreateDeployment(txGen, k),
	))

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveDeployment,
		SimulateMsgRemoveDeployment(txGen, k),
	))

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDeployment,
		SimulateMsgUpdateDeployment(txGen, k),
	))

	return operations
}

func SimulateMsgCreateDeployment(
	txGen client.TxConfig,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Intn(5)
		meta, payload := sample.CreateDatasetPayloadWithAddrAndIndexHtml(simAccount.Address.String(), i, keepertest.DATASET_SIZE)
		msg := types.MsgCreateDeploymentRequest{
			Meta:    meta,
			Payload: payload,
		}

		found := k.HasDeployment(ctx, simAccount.Address, meta.GetName())
		if found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "deployment already exist"), nil, nil
		}

		return genAndDeliverTxWithRandFees(r, app, ctx, txGen, simAccount, &msg, k)
	}
}

func SimulateMsgUpdateDeployment(
	txGen client.TxConfig,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Intn(5)
		meta, payload := sample.CreateDatasetPayloadWithAddrAndIndexHtml(simAccount.Address.String(), i, keepertest.DATASET_SIZE)
		msg := types.MsgUpdateDeploymentRequest{
			Meta:    meta,
			Payload: payload,
		}

		found := k.HasDeployment(ctx, simAccount.Address, meta.GetName())
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "deployment doesn't exist"), nil, nil
		}

		return genAndDeliverTxWithRandFees(r, app, ctx, txGen, simAccount, &msg, k)
	}
}

func SimulateMsgRemoveDeployment(
	txGen client.TxConfig,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Intn(5)
		msg := types.MsgRemoveDeploymentRequest{
			Creator: simAccount.Address.String(),
			Name:    strconv.Itoa(i),
		}

		found := k.HasDeployment(ctx, simAccount.Address, strconv.Itoa(i))
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "deployment doesn't exist"), nil, nil
		}

		return genAndDeliverTxWithRandFees(r, app, ctx, txGen, simAccount, &msg, k)
	}
}

func newOperationInput(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, txGen client.TxConfig, simAccount simtypes.Account, msg sdk.Msg, k keeper.Keeper) simulation.OperationInput {
	return simulation.OperationInput{
		R:             r,
		App:           app,
		TxGen:         txGen,
		Cdc:           nil,
		Msg:           msg,
		Context:       ctx,
		SimAccount:    simAccount,
		AccountKeeper: k.GetTestAccountKeeper(),
		Bankkeeper:    k.GetTestBankKeeper(),
		ModuleName:    types.ModuleName,
	}
}

func genAndDeliverTxWithRandFees(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, txGen client.TxConfig, simAccount simtypes.Account, msg sdk.Msg, k keeper.Keeper) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	return simulation.GenAndDeliverTxWithRandFees(newOperationInput(r, app, ctx, txGen, simAccount, msg, k))
}

func genAndDeliverTx(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, txGen client.TxConfig, simAccount simtypes.Account, msg sdk.Msg, k keeper.Keeper, fees sdk.Coins) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	return simulation.GenAndDeliverTx(newOperationInput(r, app, ctx, txGen, simAccount, msg, k), fees)
}
