package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//	types2 "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/interchainberlin/pooltoy/x/escrow/types"
	"strconv"
	//	"github.com/cosmos/cosmos-sdk/types/address"
)

func (k Keeper) InsertOffer(ctx sdk.Context, offer types.Offer) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	addr, err := sdk.AccAddressFromBech32(offer.Sender)
	if err != nil {
		return []byte{}, err
	}
	storeK := []byte(types.OfferPrefix)
	storeK = append(storeK, addr...)
	storeK = append(storeK, []byte(strconv.FormatInt(*k.index, 10))...)
	// todo maybe we do not need int64 for index
	storeV, err := offer.Marshal()
	if err != nil {
		return nil, err
	}

	store.Set(storeK, storeV)

	return storeV, nil
}

func (k Keeper) ListOffer(ctx sdk.Context, offer types.OfferListAllRequest) (types.OfferListResponse, error) {
	store := ctx.KVStore(k.storeKey)
	offers := []*types.Offer{}

	iterator := sdk.KVStorePrefixIterator(store, []byte(types.OfferPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var offer types.Offer
		offer.Unmarshal(iterator.Value())
		offers = append(offers, &offer)
	}

	return types.OfferListResponse{OfferList: offers}, nil
}

func (k Keeper) ListOfferByAddr(ctx sdk.Context, offer types.QueryOfferByAddrRequest) (types.OfferListResponse, error) {

	addr, err := sdk.AccAddressFromBech32(offer.Offerer)
	if err != nil {
		return types.OfferListResponse{}, err
	}
	addrStore := k.getAddressPrefixStore(ctx, addr)

	iterator := addrStore.Iterator(nil, nil)
	defer iterator.Close()
	offers := []*types.Offer{}
	for ; iterator.Valid(); iterator.Next() {
		var offer types.Offer
		k.cdc.MustUnmarshalBinaryBare(addrStore.Get(iterator.Key()), &offer)
		offers = append(offers, &offer)
	}

	return types.OfferListResponse{offers}, nil
}

func (k Keeper) getAddressPrefixStore(ctx sdk.Context, addr sdk.AccAddress) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	addrStore := prefix.NewStore(store, CreateAddrPrefix(addr))

	return addrStore
}

func CreateAddrPrefix(addr []byte) []byte {
	//todo change to  MustLengthPrefix(addr)...) for v0.43.0 release
	prefix := append([]byte(types.OfferPrefix), addr...)
	fmt.Println("the prefix len is!!!!", len(prefix))
	return prefix
}
