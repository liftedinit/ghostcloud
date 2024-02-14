package ghostcloud_test

import (
	"testing"

	"github.com/liftedinit/ghostcloud/testutil/sample"

	keepertest "github.com/liftedinit/ghostcloud/testutil/keeper"
	"github.com/liftedinit/ghostcloud/testutil/nullify"
	"github.com/liftedinit/ghostcloud/x/ghostcloud"
	"github.com/liftedinit/ghostcloud/x/ghostcloud/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:      types.DefaultParams(),
		Deployments: sample.CreateNDeployments(keepertest.NUM_DEPLOYMENT, keepertest.DATASET_SIZE),
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GhostcloudKeeper(t)
	ghostcloud.InitGenesis(ctx, *k, genesisState)
	got := ghostcloud.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.Deployments, got.Deployments)
	// this line is used by starport scaffolding # genesis/test/assert
}
