package keeper

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/liftedinit/ghostcloud/x/ghostcloud/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetAllDeployments(ctx sdk.Context, k Keeper) (deployments []*types.Deployment) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	iterator := storetypes.KVStorePrefixIterator(store, types.DeploymentMetaKeyPrefix)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var meta types.Meta
		k.cdc.MustUnmarshal(iterator.Value(), &meta)

		creator := sdk.MustAccAddressFromBech32(meta.GetCreator())
		dataset := k.GetDataset(ctx, creator, meta.GetName())

		deployments = append(deployments, &types.Deployment{
			Meta:    &meta,
			Dataset: dataset,
		})
	}

	return
}
