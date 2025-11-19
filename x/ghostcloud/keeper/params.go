package keeper

import (
	"github.com/liftedinit/ghostcloud/x/ghostcloud/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetParams(ctx sdk.Context) (types.Params, error) { return k.params.Get(ctx) }
func (k Keeper) SetParams(ctx sdk.Context, p types.Params) error { return k.params.Set(ctx, p) }
