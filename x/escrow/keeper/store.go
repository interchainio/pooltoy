package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/escrow/types"
	"strconv"
)

func (k Keeper) InsertOffer(ctx sdk.Context, offer types.OfferRequest) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	addr, err := sdk.AccAddressFromBech32(offer.Sender)
	if err != nil {
		return []byte{}, err
	}
	storeK := []byte{}
	storeK = append(storeK, addr...)
	storeK = append(storeK, []byte(strconv.FormatInt(*k.index, 10))...)
		// todo maybe we do not need int64 for index
	storeV, err := offer.Marshal()
	if err != nil {
		return nil, err
	}
	fmt.Println("store key is", storeK)
	store.Set([]byte(storeK), storeV)

	return storeV, nil
}

func (k Keeper) ListOffer(ctx sdk.Context, offer types.OfferListAllRequest) (types.OfferListResponse, error) {
	store := ctx.KVStore(k.storeKey)
	offers := []*types.Offer{}

	iterator := sdk.KVStorePrefixIterator(store, []byte(types.OfferPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var offer types.Offer
		fmt.Println("key is!!!!", string(iterator.Key()))
		k.cdc.MustUnmarshalBinaryBare(store.Get(iterator.Key()), &offer)
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
	fmt.Println("does the store has the addr!!!!", addrStore.Has(addr))
	iterator := addrStore.Iterator(nil, nil)
	defer iterator.Close()

	offers := []*types.Offer{}
	 i :=0
	 fmt.Println("valid iterator!!!!", iterator.Valid())
	for ; iterator.Valid(); iterator.Next() {
		fmt.Println("i is!!!!", i)
		i ++
		fmt.Println("substore key is!!!!", string(iterator.Key()))
		var offer types.Offer
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &offer)
		fmt.Println("offer unmarshal is!!!!", offer.Sender, offer.Amount, offer.Request)
		offers = append(offers, &offer)
	}

	return types.OfferListResponse{offers}, nil
}

func (k Keeper) getAddressPrefixStore(ctx sdk.Context, addr sdk.AccAddress) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	ref :="offer-"
	b := []byte(ref)
	fmt.Println("ref b is!!!!",b)
	addr, _ = sdk.AccAddressFromBech32("cosmos126c2ak7k5qw8rhhha60kgksjrjf5jvkgtzlklw")
	fmt.Println("ref addr !!!!", addr)
	fmt.Println("the new prefix is!!!!", CreateAddrPrefix(addr))
	addrStore := prefix.NewStore(store, CreateAddrPrefix(addr))

	return addrStore
}


func CreateAddrPrefix(addr []byte) []byte {
	return append([]byte(types.OfferPrefix), addr...)
}

// todo cannot find the pakcage: address.MustLengthPrefix(addr)...
