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

	err := k.SetParams(ctx, params)
	require.NoError(t, err)

	params, err = k.GetParams(ctx)
	require.NoError(t, err)

	require.EqualValues(t, params, params)
}
