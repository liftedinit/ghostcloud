package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	storecore "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/liftedinit/ghostcloud/x/ghostcloud/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService storecore.KVStoreService
		logger       log.Logger
		Schema       collections.Schema
		params       collections.Item[types.Params]
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService storecore.KVStoreService,
	logger log.Logger,

) Keeper {
	logger = logger.With(log.ModuleKey, "x/"+types.ModuleName)

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:          cdc,
		storeService: storeService,
		logger:       logger,
	}
	k.params = collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc))

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) HasDeployment(ctx sdk.Context, creator sdk.AccAddress, name string) bool {
	base := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(base, types.DeploymentMetaKeyPrefix)
	return store.Has(types.DeploymentKey(creator, name))
}

func (k Keeper) SetDeployment(ctx sdk.Context, addr sdk.AccAddress, meta *types.Meta, dataset *types.Dataset) {
	k.SetMeta(ctx, addr, meta)
	k.SetDataset(ctx, addr, meta.GetName(), dataset)
}

func (k Keeper) SetMeta(ctx sdk.Context, addr sdk.AccAddress, meta *types.Meta) {
	base := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(base, types.DeploymentMetaKeyPrefix)
	b := k.cdc.MustMarshal(meta)
	store.Set(types.DeploymentKey(addr, meta.GetName()), b)
}

func (k Keeper) SetDataset(ctx sdk.Context, addr sdk.AccAddress, name string, dataset *types.Dataset) {
	// NOTE: Safe to ignore the error here because the caller ensures that
	for _, item := range dataset.GetItems() {
		k.SetItem(ctx, addr, name, item)
	}
}

func (k Keeper) Remove(ctx sdk.Context, addr sdk.AccAddress, name string) {
	k.RemoveDataset(ctx, addr, name)

	base := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(base, types.DeploymentMetaKeyPrefix)
	store.Delete(types.DeploymentKey(addr, name))
}

func (k Keeper) RemoveDataset(ctx sdk.Context, addr sdk.AccAddress, name string) {
	base := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	// Delete item metas
	metaStore := prefix.NewStore(base, types.DeploymentItemMetaPrefix)
	it := storetypes.KVStorePrefixIterator(metaStore, types.DeploymentKey(addr, name))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		metaStore.Delete(it.Key())
	}

	// Delete item contents
	contentStore := prefix.NewStore(base, types.DeploymentItemContentPrefix)
	it2 := storetypes.KVStorePrefixIterator(contentStore, types.DeploymentKey(addr, name))
	defer it2.Close()

	for ; it2.Valid(); it2.Next() {
		contentStore.Delete(it2.Key())
	}
}

func (k Keeper) SetItem(ctx sdk.Context, addr sdk.AccAddress, name string, item *types.Item) {
	base := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	// Set Item meta
	metaStore := prefix.NewStore(base, types.DeploymentItemMetaPrefix)
	meta := item.GetMeta()
	path := meta.GetPath()

	b := k.cdc.MustMarshal(meta)
	metaStore.Set(types.DeploymentItemKey(addr, name, path), b)

	// Set Item content
	contentStore := prefix.NewStore(base, types.DeploymentItemContentPrefix)
	b = k.cdc.MustMarshal(item.GetContent())
	contentStore.Set(types.DeploymentItemKey(addr, name, path), b)
}

func (k Keeper) GetDataset(ctx sdk.Context, addr sdk.AccAddress, name string) (dataset *types.Dataset) {
	base := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	metaStore := prefix.NewStore(base, types.DeploymentItemMetaPrefix)
	it := storetypes.KVStorePrefixIterator(metaStore, types.DeploymentKey(addr, name))
	defer it.Close()

	items := make([]*types.Item, 0)
	for ; it.Valid(); it.Next() {
		var meta types.ItemMeta
		k.cdc.MustUnmarshal(it.Value(), &meta)

		contentStore := prefix.NewStore(base, types.DeploymentItemContentPrefix)
		b := contentStore.Get(types.DeploymentItemKey(addr, name, meta.GetPath()))
		if b == nil {
			continue
		}

		var content types.ItemContent
		k.cdc.MustUnmarshal(b, &content)

		items = append(items, &types.Item{
			Meta:    &meta,
			Content: &content,
		})
	}

	return &types.Dataset{Items: items}
}

func (k Keeper) GetMeta(ctx sdk.Context, addr sdk.AccAddress, name string) (meta types.Meta, found bool) {
	base := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(base, types.DeploymentMetaKeyPrefix)
	b := store.Get(types.DeploymentKey(addr, name))
	if b == nil {
		return meta, false
	}

	k.cdc.MustUnmarshal(b, &meta)
	return meta, true
}

func (k Keeper) GetItemContent(ctx sdk.Context, addr sdk.AccAddress, name string, path string) (content types.ItemContent, found bool) {
	base := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(base, types.DeploymentItemContentPrefix)
	b := store.Get(types.DeploymentItemKey(addr, name, path))
	if b == nil {
		return content, false
	}

	k.cdc.MustUnmarshal(b, &content)
	return content, true
}

func (k Keeper) getDeploymentMetaStore(ctx sdk.Context) prefix.Store {
	base := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(base, types.DeploymentMetaKeyPrefix)
}

func (k Keeper) GetAllMeta(ctx sdk.Context) (metas []*types.Meta) {
	base := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	it := storetypes.KVStorePrefixIterator(base, types.DeploymentMetaKeyPrefix)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		var meta types.Meta
		k.cdc.MustUnmarshal(it.Value(), &meta)
		metas = append(metas, &meta)
	}

	return metas
}
