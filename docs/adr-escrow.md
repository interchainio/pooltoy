# ADR-0002: Escrow

## Changelog

- 08.2021: Adding escrow module to pooltoy

## Abstract

 This ADR will discuss the design of the escrow module, with the focus on the message types, the escrow process, and data store.


## Context

The escrow module provides the "middleman" function when two pooltoy users swap emojis.
The escrow ensures the two participants agree on the emojis exchange and get the promised emojis.

### Prospective chain-to-chain flow

- Chain A is connected to the Cosmos Hub via IBC and now wants to perform a token swap with the Cosmos Hub
- A user on chain A constructs a governance proposal that does the following:
  - Creates an interchain account in which chain A's gov module is the controller, with the Cosmos Hub as the host chain
  - Community spend proposal of X (+ some extra for fees) tokens from chain A distribution module to chain A gov module
  - IBC send X (+ some extra for fees) tokens to the interchain account host address on the Cosmos Hub
  - Send escrow offer message via interchain account: I wish to offer X chain A tokens for Y ATOM, which will transfer X chain A tokens to the escrow module
- A user on the Cosmos Hub must then make a proposal which does the following:
  - Community spend proposal for Y ATOM (+ some extra for fees) going to the gov module
  - Escrow fill message, accepting the offer
- Once this proposal passes and the fill offer is sent, X chain A tokens will be moved to the Cosmos Hub gov module account and Y ATOMs will be moved to the host interchain account controlled by chain A gov module account
- A user on chain A may then make a gov proposal which sends and interchain account message to IBC send Y ATOMs to chain A's gov module


## Decision: escrow process

We decide to make the escrow process into 3 steps: offer &#8594; response &#8594; exchange. There are also two extra auxiliary
steps: offer query and offer cancel to help escrow participants finish the escrow process.  
 
### Offer

> Offerer: *"Hi, I have 2🍎, and want to use them to exchange for 1🍕. I will send my 2🍎 to escrow!"*

A pooltoy user can send some number of emojis to the module account, with an offer description including what they wish to swap for with those emojis, and optionally with whom they would like to swap.


```go
type Offer struct {
	Sender  string  
	// Offer.Sender is an address initiates the escrow process
	Amount  github_com_cosmos_cosmos_sdk_types.Coins 
	// Offer.Amount is the list of emojis Offer.Sender offers at escrow.
	Request github_com_cosmos_cosmos_sdk_types.Coins 
	// Offer.Request is the list of emojis Offer.Sender plans to use Offer.Amount to exchange for.
}

```
The sender must make sure that they have enough tokens to make an offer of size `Offer.Amount`.
When the sender sends the `Offer` message to the escrow module, the `Offer.Amount` will be transferred from this user to the escrow module account (a public account that no nodes hold a private key), then this `Offer` is recorded in the store with a self-increased ID (Details please see **escrow store** section.), waiting for a responser to finish the emoji exchange.

Please note the `Offer.Amount` is bonded token after offering successfully.`Offer.Sender` cannot spend emojis which are at escrow, but `Offer.Sender` can make them unbonded again through `CancelOffer`.


### Response (should rename to fill)

> Resonser: *"Hi, I have 1🍕, I want those 2🍎!"*

The `Resonse.Sender` is the person who is willing to provide `Offer.Request` of an `Offer` and to receive this `Offer.Sender`'s `Offer.Amount`.
The resonser has to make sure that he has enough balances for `Offer.Request` and send the Resonse message to escrow module:
```go
type Response struct {
	Sender string
	// Response.Sender is an address, this address holder responses to the Offer with ID = Response.ID
	ID     int64
	// Offer ID in the store
}
```


### Swap

> Escrow: *"Hi responser, the 2🍎 is at my hand now, I can help send your 1🍕 to the offerer, and send those 2🍎 to you!"*

At response step, if the offer is still valid, the escrow module will do an atomic operation consisting of 3 steps:
- module account sends the `Offer.Amount` to the `Response.Sender`,
- the `Response.Sender` sends directly the `Offer.Request` to the `Offer.Sender`.
- escrow module deletes this offer record from the escrow store.


### Cancel offer

The `Offer.Sender` can cancel the offer anytime before a  `Response.Sender` responses to this offer. By cancellation, the module account will send the `Offer.Amount` back to the `Offer.Sender` and delete this offer record from the escrow store. 
```go
type CancelOffer struct {
	Sender string
	ID     int64
}
```


### Query offer(s)

To help the responsers check which offers are in escrow now, escrow module provides query offers in 3 ways: `offer-list-all`, `offer-by-addr`, and `offer-by-ID`.

```go
// query all the offers
type OfferListAllRequest struct {}

// query an offer by its ID
type QueryOfferByIDRequest struct {
	Querier string 
	Id      int64  
}

// query offers from a certain address
type QueryOfferByAddrRequest struct {
    Querier string 
    Offerer string
}

type OfferListResponse struct {
    OfferList []Offer
}
```


## Decision: escrow store

The escrow database stores only the valid, active `Offer`s. The canceled or successfully responsed offers are deleted from the store.  
The `store` consists of key-value pairs.The key is bytes of `[store_prefix][offer_sender_address][ID]`, value is serialized `Offer`.

An address store is the subset of the parent store created by `prefix.NewStore()`, which uses `[store_prefix][address]` as the new prefix.
```go
addrStore := prefix.NewStore(store, CreateAddrPrefix(addr))
```


## Discussion: edge cases

#### Empty offers

> Offerer 1: *"If anyone want to give me 1🍕?"*
> Offerer 2: *"I just want to give away my 4🌲!"*

In an `Offer`, empty `Offer.Amount` is allowed, empty `Offer.Request` is allowed. Those two cases can be the wishes of the offerers. However,  `Offer.Amount == nil && Offer.Request == nil` will throw an error, because malicious node can take advantage of pooltoy's zero gas fee policy (writing to DB does not consume gas) to overwhelmingly writing to our store without essential contents.


## Further Discussions

- `Offer` is stored with `ID` (int64). There will be one day the `ID` is too big for int64, shall we use big int or reuse the deleted `Offer` `ID`?
- `CancelOffer` message only needs ID ? or sender plus ID?
- Do we need querier address in `QueryOffer` messages such as in `QueryOfferByIDRequest` and `QueryOfferByAddrRequest`?
- Should there be time limit for an active `Offer`?
- If the `OfferListResponse` is long, Should `Offer`s being ordered by creating time when query all offers?

