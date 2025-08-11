package ghostcloud

import (
	"fmt"

	"github.com/liftedinit/ghostcloud/x/ghostcloud/keeper"
	"github.com/liftedinit/ghostcloud/x/ghostcloud/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, deployment := range genState.Deployments {
		addr := sdk.MustAccAddressFromBech32(deployment.Meta.Creator)
		k.SetDeployment(ctx, addr, deployment.Meta, deployment.Dataset)
	}
	// this line is used by starport scaffolding # genesis/module/init
	err := k.SetParams(ctx, genState.Params)
	if err != nil {
		fmt.Printf("Error setting ghostcloud params %+v", err)
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	params, err := k.GetParams(ctx)
	if err != nil {
		panic(err)
	}
	genesis.Params = params

	genesis.Deployments = keeper.GetAllDeployments(ctx, k)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
