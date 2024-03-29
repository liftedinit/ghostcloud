package keeper

import (
	"testing"

	"github.com/liftedinit/ghostcloud/testutil/sample"
	"github.com/liftedinit/ghostcloud/x/ghostcloud/keeper"
	"github.com/liftedinit/ghostcloud/x/ghostcloud/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/stretchr/testify/require"
)

const (
	NUM_DEPLOYMENT = 10
	DATASET_SIZE   = 5
)

type MsgServerTestCase struct {
	Name     string
	Metas    []*types.Meta
	Payloads []*types.Payload
	Err      error
}

func GhostcloudKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"GhostcloudParams",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}

func setDeployments(ctx sdk.Context, k *keeper.Keeper, metas []*types.Meta, datasets []*types.Dataset) {
	for i := 0; i < len(metas); i++ {
		addr := sdk.MustAccAddressFromBech32(metas[i].Creator)
		k.SetDeployment(ctx, addr, metas[i], datasets[i])
	}
}
func CreateAndSetNDeployments(ctx sdk.Context, k *keeper.Keeper, numDeployment int, datasetSize int) ([]*types.Meta, []*types.Dataset) {
	metas, datasets := sample.CreateNMetaDataset(numDeployment, datasetSize)
	setDeployments(ctx, k, metas, datasets)
	return metas, datasets
}

func CreateAndSetNDeploymentsWithAddr(ctx sdk.Context, k *keeper.Keeper, numDeployment int, datasetSize int, addr string) ([]*types.Meta, []*types.Dataset) {
	metas, datasets := sample.CreateNMetaDatasetWithAddr(addr, numDeployment, datasetSize)
	setDeployments(ctx, k, metas, datasets)
	return metas, datasets
}
