package keeper

import (
	"testing"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	storemetrics "cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/liftedinit/ghostcloud/testutil/sample"
	"github.com/liftedinit/ghostcloud/x/ghostcloud/keeper"
	"github.com/liftedinit/ghostcloud/x/ghostcloud/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	logger := log.NewNopLogger()

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, logger, storemetrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	k := keeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		logger,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	_ = k.SetParams(ctx, types.DefaultParams())

	return &k, ctx
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
