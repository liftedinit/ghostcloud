package keeper_test

import (
	"testing"

	testkeeper "github.com/liftedinit/ghostcloud/testutil/keeper"
	"github.com/liftedinit/ghostcloud/x/ghostcloud/types"

	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.GhostcloudKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
