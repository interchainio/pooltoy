package app

import (
	"io"
	"os"
	"time"

	"github.com/charleenfei/modules/incubator/faucet"
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/capability"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	ibc "github.com/cosmos/cosmos-sdk/x/ibc/core"
	"github.com/cosmos/cosmos-sdk/x/supply"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"

	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	appparams "github.com/interchainberlin/pooltoy/app/params"
	"github.com/interchainberlin/pooltoy/x/pooltoy"
	pooltoykeeper "github.com/interchainberlin/pooltoy/x/pooltoy/keeper"

	ante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	pooltoytypes "github.com/interchainberlin/pooltoy/x/pooltoy/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
)

const appName = "app"

var (
	DefaultCLIHome  = os.ExpandEnv("$HOME/.pooltoycli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.pooltoyd")

	// defines basic modules
	ModuleBasics = module.NewBasicManager(
		genutil.AppModuleBasic{},
		capability.AppModuleBasic{},
		ibc.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler,
			distrclient.ProposalHandler,
			upgradeclient.ProposalHandler,
		),
		params.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		slashing.AppModuleBasic{},
		// supply.AppModuleBasic{},
		pooltoy.AppModuleBasic{},
		faucet.AppModule{},
	)

	// module acct permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		faucet.ModuleName:              {authtypes.Minter},
	}

	// new Bank module update: module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{
		distrtypes.ModuleName: true,
	}
)

// func MakeCodec() *codec.Codec {
// 	var cdc = codec.New()

// 	ModuleBasics.RegisterCodec(cdc)
// 	vesting.RegisterCodec(cdc)
// 	sdk.RegisterCodec(cdc)
// 	codec.RegisterCrypto(cdc)

// 	return cdc.Seal()
// }

type NewApp struct {
	*bam.BaseApp
	appName string

	cdc               *codec.LegacyAmino
	appCodec          codec.Marshaler
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys
	keys  map[string]*sdk.KVStoreKey
	tKeys map[string]*sdk.TransientStoreKey

	// subspaces map[string]params.Subspace

	// keepers
	ParamsKeeper   paramskeeper.Keeper
	AccountKeeper  authkeeper.AccountKeeper
	BankKeeper     bankkeeper.Keeper
	StakingKeeper  stakingkeeper.Keeper
	SlashingKeeper slashingkeeper.Keeper
	GovKeeper      govkeeper.Keeper
	UpgradeKeeper  upgradekeeper.Keeper
	DistrKeeper    distrkeeper.Keeper
	PooltoyKeeper  pooltoykeeper.Keeper
	FaucetKeeper   faucet.Keeper

	mm *module.Manager

	sm *module.SimulationManager
}

var _ simapp.App = (*NewApp)(nil)

func NewInitApp(
	encodingConfig appparams.EncodingConfig,
	homePath string,
	logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, skipUpgradeHeights map[int64]bool, baseAppOptions ...func(*bam.BaseApp),
) *NewApp {

	appCodec := encodingConfig.Marshaler
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := bam.NewBaseApp(
		appName,
		logger, db,
		encodingConfig.TxConfig.TxDecoder(),
		baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey,
		stakingtypes.StoreKey,
		distrtypes.StoreKey,
		slashingtypes.StoreKey,
		paramstypes.StoreKey,
		govtypes.StoreKey,
		upgradetypes.StoreKey,
		pooltoytypes.StoreKey,
		faucet.StoreKey)

	tKeys := sdk.NewTransientStoreKeys(
		// TODO: stakingtype transient store key??
		// stakingtypes.TStoreKey,
		paramstypes.TStoreKey,
	)

	app := &NewApp{
		BaseApp:        bApp,
		cdc:            cdc,
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		tKeys:          tKeys,
		// subspaces:      make(map[string]params.Subspace),
	}

	app.ParamsKeeper = initParamsKeeper(
		appCodec,
		cdc,
		keys[paramstypes.StoreKey],
		tKeys[paramstypes.TStoreKey],
	)

	// set the BaseApp's parameter store
	bApp.SetParamStore(
		app.ParamsKeeper.Subspace(bam.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()),
	)

	// app.paramsKeeper = params.NewKeeper(app.cdc, keys[params.StoreKey], tKeys[params.TStoreKey])
	// app.subspaces[auth.ModuleName] = app.paramsKeeper.Subspace(auth.DefaultParamspace)
	// app.subspaces[bank.ModuleName] = app.paramsKeeper.Subspace(bank.DefaultParamspace)
	// app.subspaces[staking.ModuleName] = app.paramsKeeper.Subspace(staking.DefaultParamspace)
	// app.subspaces[distr.ModuleName] = app.paramsKeeper.Subspace(distr.DefaultParamspace)
	// app.subspaces[slashing.ModuleName] = app.paramsKeeper.Subspace(slashing.DefaultParamspace)

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		app.GetSubspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms,
	)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.BlockedAddrs(),
	)

	// app.supplyKeeper = supply.NewKeeper(
	// 	app.cdc,
	// 	keys[supply.StoreKey],
	// 	app.accountKeeper,
	// 	app.bankKeeper,
	// 	maccPerms,
	// )

	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(stakingtypes.ModuleName),
	)

	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		keys[distrtypes.StoreKey],
		app.GetSubspace(distrtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&stakingKeeper,
		authtypes.FeeCollectorName,
		app.ModuleAccountAddrs(),
	)

	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		keys[slashingtypes.StoreKey],
		&stakingKeeper,
		app.GetSubspace(slashingtypes.ModuleName),
	)

	// app.takingKeeper = *stakingKeeper.SetHooks(
	// 	staking.NewMultiStakingHooks(
	// 		app.distrKeeper.Hooks(),
	// 		app.slashingKeeper.Hooks()),
	// )

	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
	)
	// no-op handler for "nextversion" upgrade
	// app.upgradeKeeper.SetUpgradeHandler("nextversion", func(ctx sdk.Context, plan upgrade.Plan) {
	// 	// faucetAddr, _ := sdk.AccAddressFromBech32("akash1dz70tfxxrsh8fned6len7feu3atz7k59zgz77n")
	// 	// personalAcc1, _ := sdk.AccAddressFromBech32("akash1dz70tfxxrsh8fned6len7feu3atz7k59zgz77n")
	// 	// personalAcc2, _ := sdk.AccAddressFromBech32("akash1y4uskp4v6s086pdg6phufqcg82g9xd494kdwmn")
	// 	// valAddr, _ := sdk.ValAddressFromBech32("akashvaloper1753ew9z7cfhu6awyp6ff7qtm9ex30kg3r7zd7j")

	// 	// _, _ = app.bankKeeper.AddCoins(ctx, faucetAddr, sdk.Coins{sdk.Coin{Denom: "uakt", Amount: sdk.NewInt(10000000000)}})
	// 	// _, _ = app.bankKeeper.AddCoins(ctx, personalAcc1, sdk.Coins{sdk.Coin{Denom: "uakt", Amount: sdk.NewInt(10000000000)}})
	// 	// _, _ = app.bankKeeper.AddCoins(ctx, personalAcc2, sdk.Coins{sdk.Coin{Denom: "uakt", Amount: sdk.NewInt(10000000000)}})

	// 	// delegation := stakingTypes.Delegation{
	// 	// 	DelegatorAddress: faucetAddr,
	// 	// 	ValidatorAddress: valAddr,
	// 	// 	Shares:           sdk.NewDec(100000000000),
	// 	// }

	// 	// app.stakingKeeper.SetDelegation(ctx, delegation)

	// 	votingParams := gov.NewVotingParams(6000000000)
	// 	app.paramsKeeper.Subspace(gov.DefaultParamspace).Set(ctx, gov.ParamStoreKeyVotingParams, &votingParams)
	// })

	app.PooltoyKeeper = pooltoykeeper.NewKeeper(
		appCodec,
		keys[pooltoytypes.StoreKey],
		app.BankKeeper,
		app.AccountKeeper,
	)

	app.FaucetKeeper = faucet.NewKeeper(
		// TODO: look into Supply Keeper!
		app.SupplyKeeper,
		app.StakingKeeper,
		app.AccountKeeper,
		1,            // amount for mint
		24*time.Hour, // rate limit by time
		keys[faucet.StoreKey],
		app.cdc,
	)

	// register the proposal types
	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper))

	app.GovKeeper = govkeeper.NewKeeper(
		appCodec,
		keys[govtypes.StoreKey],
		app.GetSubspace(govtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&stakingKeeper,
		govRouter,
	)

	bankModule := bank.NewAppModule(
		appCodec,
		app.BankKeeper,
		app.AccountKeeper,
	)
	restrictedBank := NewRestrictedBankModule(
		bankModule,
		app.BankKeeper,
		app.AccountKeeper,
	)

	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper,
			app.StakingKeeper,
			app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(
			appCodec,
			app.AccountKeeper,
			nil,
		),
		restrictedBank,
		// TODO: why two banks?
		// bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		// supply.NewAppModule(app.supplyKeeper, app.accountKeeper),
		distr.NewAppModule(
			appCodec,
			app.DistrKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.StakingKeeper,
		),
		slashing.NewAppModule(
			appCodec,
			app.SlashingKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.StakingKeeper,
		),
		staking.NewAppModule(
			appCodec,
			app.StakingKeeper,
			app.AccountKeeper,
			app.BankKeeper,
		),
		gov.NewAppModule(
			appCodec,
			app.GovKeeper,
			app.AccountKeeper,
			app.BankKeeper,
		),
		pooltoy.NewAppModule(
			appCodec,
			app.PooltoyKeeper,
			app.BankKeeper,
		),

		faucet.NewAppModule(
			appCodec,
			app.FaucetKeeper,
		),
		upgrade.NewAppModule(app.UpgradeKeeper),
	)

	app.mm.SetOrderBeginBlockers(
		upgradetypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		stakingtypes.ModuleName,
		govtypes.ModuleName,
	)

	app.mm.SetOrderInitGenesis(
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		pooltoytypes.ModuleName,
		// supplytypes.ModuleName,
		genutiltypes.ModuleName,
	)

	app.mm.RegisterRoutes(
		app.Router(),
		app.QueryRouter(),
		encodingConfig.Amino,
	)
	app.mm.RegisterServices(
		module.NewConfigurator(app.MsgServiceRouter(), app.GRPCQueryRouter()),
	)

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tKeys)

	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.SetAnteHandler(
		NewAnteHandler(
			app.AccountKeeper,
			app.BankKeeper,
			ante.DefaultSigVerificationGasConsumer,
			encodingConfig.TxConfig.SignModeHandler(),
		),
	)

	app.MountKVStores(keys)
	app.MountTransientStores(tKeys)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}
	return app
}

func (app *NewApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

func (app *NewApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

func (app *NewApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

func (app *NewApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

func (app *NewApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[supply.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// LegacyAmino returns SimApp's amino codec.
func (app *NewApp) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// func (app *NewApp) Codec() *codec.Codec {
// 	return app.cdc
// }

// GetSubspace returns a param subspace for a given module name.
func (app *NewApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

func (app *NewApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *NewApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Register legacy tx routes.
	authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *NewApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *NewApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

func GetMaccPerms() map[string][]string {
	modAccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		modAccPerms[k] = v
	}
	return modAccPerms
}

// BlockedAddrs returns all the app's module account addresses that are not
// allowed to receive external tokens.
func (app *NewApp) BlockedAddrs() map[string]bool {
	blockedAddrs := make(map[string]bool)
	for acc := range maccPerms {
		blockedAddrs[authtypes.NewModuleAddress(acc).String()] = !allowedReceivingModAcc[acc]
	}

	return blockedAddrs
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryMarshaler, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)

	return paramsKeeper
}
