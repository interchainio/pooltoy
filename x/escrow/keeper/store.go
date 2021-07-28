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

	fmt.Println("storeK is!!!!",storeK)
	storeK = append(storeK, []byte(strconv.FormatInt(*k.index, 10))...)
		// todo maybe we do not need int64 for index
	storeV, err := offer.Marshal()
	if err != nil {
		return nil, err
	}
	fmt.Println("store key is", storeK)
	store.Set(storeK, storeV)

	return storeV, nil
}

func (k Keeper) ListOffer(ctx sdk.Context, offer types.OfferListAllRequest) (types.OfferListResponse, error) {
	store := ctx.KVStore(k.storeKey)
	offers := []*types.Offer{}

	iterator := sdk.KVStorePrefixIterator(store, []byte(types.OfferPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var offer types.Offer
		fmt.Println("key is!!!!", string(iterator.Key()))
	//	k.cdc.MustUnmarshalBinaryBare(store.Get(iterator.Key()), &offer)
	offer.Unmarshal(iterator.Value())
		offers = append(offers, &offer)
	}

	return types.OfferListResponse{OfferList: offers}, nil
}

func (k Keeper) ListOfferByAddr(ctx sdk.Context, offer types.QueryOfferByAddrRequest) (types.OfferListResponse, error) {

	addr, err := sdk.AccAddressFromBech32(offer.Offerer)
	fmt.Println("addr is!!!!", addr.String())
	if err != nil {
		return types.OfferListResponse{}, err
	}
	addrStore := k.getAddressPrefixStore(ctx, addr)
	//fmt.Println("does the store has the addr!!!!", addrStore.Has(addrStore.key()))
	iterator := addrStore.Iterator(nil, nil)
	defer iterator.Close()

	offers := []*types.Offer{}
	 //i :=0
	 fmt.Println("valid iterator!!!!", iterator.Valid())
	fmt.Println("valid iterator err!!!!", iterator.Error())
	for ; iterator.Valid(); iterator.Next() {
		fmt.Println("start iterate sub store!!!!")
		fmt.Println("substore key is!!!!", string(iterator.Key()))
		var offer types.Offer
		k.cdc.MustUnmarshalBinaryBare(addrStore.Get(iterator.Key()), &offer)
		offers = append(offers, &offer)
	}

	return types.OfferListResponse{offers}, nil
}

func (k Keeper) getAddressPrefixStore(ctx sdk.Context, addr sdk.AccAddress) prefix.Store {
	store := ctx.KVStore(k.storeKey)
//	ref :="offer-"
	b := []byte(types.OfferPrefix)
	fmt.Println("ref b is!!!!",b)
	addr, _ = sdk.AccAddressFromBech32("cosmos126c2ak7k5qw8rhhha60kgksjrjf5jvkgtzlklw")
	fmt.Println("ref addr !!!!", addr)
	fmt.Println("the new prefix is!!!!", CreateAddrPrefix(addr))
	addrStore := prefix.NewStore(store, CreateAddrPrefix(addr))

	return addrStore
}


func CreateAddrPrefix(addr []byte) []byte {
//todo change to  MustLengthPrefix(addr)...) for v0.43.0 release
	//prefix := append([]byte(types.OfferPrefix),[]byte{byte(len(addr))}...)
	prefix := append([]byte(types.OfferPrefix), addr...)
	fmt.Println("the prefix is!!!!", prefix)
	return prefix
}

// todo cannot find the pakcage: address.MustLengthPrefix(addr)...
