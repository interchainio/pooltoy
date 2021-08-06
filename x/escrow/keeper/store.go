package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"

	//	types2 "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/interchainberlin/pooltoy/x/escrow/types"
	"strconv"
	//	"github.com/cosmos/cosmos-sdk/types/address"
)

func (k Keeper) InsertOffer(ctx sdk.Context, offerReq types.OfferRequest) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	addr, err := sdk.AccAddressFromBech32(offerReq.Sender)
	if err != nil {
		return []byte{}, err
	}
	storeK := []byte(types.OfferPrefix)
	storeK = append(storeK, addr...)
	id := k.GetUpdatedID(ctx)
	offer := types.Offer{
		Sender: offerReq.Sender,
		Amount: offerReq.Amount,
		Request: offerReq.Request,
		Id: id,
	}
	storeK = append(storeK, []byte(strconv.FormatInt(id, 10))...)

	// todo maybe we do not need int64 for index
	storeV, err := offer.Marshal()
	if err != nil {
		return nil, err
	}

	store.Set(storeK, storeV)
	return storeV, nil
}

func (k Keeper) DeleteOffer(ctx sdk.Context, id int64) error {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.OfferPrefix))
	defer iterator.Close()

	foundID := false
	for ; iterator.Valid(); iterator.Next() {
		idBytes := iterator.Key()[types.AddrPrefixLen:]
		i, err := BytesToInt64(idBytes)
		if err != nil {
			return errors.New("convert ID failed")
		}
		if i == id {
			foundID = true
			store.Delete(iterator.Key())
		}
	}
	// the ID not found
	if foundID == false {
		return errors.New("ID not found")
	}
	return nil
}

func (k Keeper) ListOffer(ctx sdk.Context, offer types.OfferListAllRequest) (types.OfferListResponse, error) {
	store := ctx.KVStore(k.storeKey)
	offers := []*types.Offer{}

	iterator := sdk.KVStorePrefixIterator(store, []byte(types.OfferPrefix))
	defer iterator.Close()
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

	return prefix
}

func (k Keeper) ListOfferByID(ctx sdk.Context, offerReq types.QueryOfferByIDRequest) (types.Offer, error) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.OfferPrefix))
	defer iterator.Close()

	var offer types.Offer
	for ; iterator.Valid(); iterator.Next() {
		idBytes := iterator.Key()[types.AddrPrefixLen:]

		i, err := BytesToInt64(idBytes)
		if err != nil {
			return offer, errors.New("convert ID failed")
		}
		if i == offerReq.Id {
			k.cdc.MustUnmarshalBinaryBare(store.Get(iterator.Key()), &offer)
		}
	}
	// the ID not found
	if offer.Sender == "" {
		return offer, errors.New("ID not found")
	}

	return offer, nil
}

func BytesToInt64(b []byte) (int64, error) {
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return 0, err
	}
	return int64(i), nil
}

// store for record latest ID
func (k Keeper) GetLatestID(ctx sdk.Context) int64 {
	idStore := ctx.KVStore(k.idStoreKey)
	if !idStore.Has([]byte(types.IDStoreKey)) {
		// store is empty
		return int64(-1)
	}

	b := idStore.Get([]byte(types.IDStoreKey))
	i, _ := BytesToInt64(b)
	return i
}

// store id + 1
func (k Keeper) GetUpdatedID(ctx sdk.Context) int64 {
	id := k.GetLatestID(ctx)
	idStore := ctx.KVStore(k.idStoreKey)
	idStore.Set([]byte(types.IDStoreKey), []byte(strconv.FormatInt(id+1, 10)))

	return id + 1
}
