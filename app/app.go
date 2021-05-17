package app

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	transfer "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer"
	ibctransferkeeper "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/keeper"
	ibctransfertypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	ibc "github.com/cosmos/cosmos-sdk/x/ibc/core"
	ibcclient "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client"
	porttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/05-port/types"
	ibchost "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	ibckeeper "github.com/cosmos/cosmos-sdk/x/ibc/core/keeper"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	appparams "github.com/interchainberlin/pooltoy/app/params"
	"github.com/interchainberlin/pooltoy/x/faucet"
	faucetkeeper "github.com/interchainberlin/pooltoy/x/faucet/keeper"
	faucettypes "github.com/interchainberlin/pooltoy/x/faucet/types"
	"github.com/interchainberlin/pooltoy/x/pooltoy"
	pooltoykeeper "github.com/interchainberlin/pooltoy/x/pooltoy/keeper"
	pooltoytypes "github.com/interchainberlin/pooltoy/x/pooltoy/types"
	"github.com/spf13/cast"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

const Name = "pooltoy"

// this line is used by starport scaffolding # stargate/wasm/app/enabledProposals

func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		distrclient.ProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)

	return govProposalHandlers
}

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(getGovProposalHandlers()...),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		ibc.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
		pooltoy.AppModuleBasic{},
		faucet.AppModule{},
	// this line is used by starport scaffolding # stargate/app/moduleBasic
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		faucettypes.ModuleName:         {authtypes.Minter},
	}

	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{
		distrtypes.ModuleName: true,
	}
)

var (
	_ simapp.App              = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Marshaler
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	MintKeeper       mintkeeper.Keeper
	DistrKeeper      distrkeeper.Keeper
	GovKeeper        govkeeper.Keeper
	CrisisKeeper     crisiskeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper   evidencekeeper.Keeper
	TransferKeeper   ibctransferkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper        capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper   capabilitykeeper.ScopedKeeper
	ScopedIbcAccountKeeper capabilitykeeper.ScopedKeeper

	PooltoyKeeper pooltoykeeper.Keeper
	FaucetKeeper  faucetkeeper.Keeper
	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	// the module manager
	mm *module.Manager
	sm *module.SimulationManager
}

// New returns a reference to an initialized Gaia.
// NewSimApp returns a reference to an initialized SimApp.
func New(
	logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool,
	homePath string, invCheckPeriod uint, encodingConfig appparams.EncodingConfig,
	// this line is used by starport scaffolding # stargate/app/newArgument
	appOpts servertypes.AppOptions, baseAppOptions ...func(*baseapp.BaseApp),
) *App {

	appCodec := encodingConfig.Marshaler
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(Name, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, ibchost.StoreKey,
		evidencetypes.StoreKey, ibctransfertypes.StoreKey, capabilitytypes.StoreKey,
		pooltoytypes.StoreKey,
		faucettypes.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &App{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// set the BaseApp's parameter store
	bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()))

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey], app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
	)
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec, keys[banktypes.StoreKey], app.AccountKeeper, app.GetSubspace(banktypes.ModuleName), app.BlockedAddrs(),
	)
	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec, keys[stakingtypes.StoreKey], app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName),
	)
	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec, keys[minttypes.StoreKey], app.GetSubspace(minttypes.ModuleName), &stakingKeeper,
		app.AccountKeeper, app.BankKeeper, authtypes.FeeCollectorName,
	)
	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec, keys[distrtypes.StoreKey], app.GetSubspace(distrtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, authtypes.FeeCollectorName, app.ModuleAccountAddrs(),
	)
	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec, keys[slashingtypes.StoreKey], &stakingKeeper, app.GetSubspace(slashingtypes.ModuleName),
	)
	app.CrisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName), invCheckPeriod, app.BankKeeper, authtypes.FeeCollectorName,
	)
	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks()),
	)

	// ... other modules keepers

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, keys[ibchost.StoreKey], app.GetSubspace(ibchost.ModuleName), app.StakingKeeper, scopedIBCKeeper,
	)

	// register the proposal types
	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(ibchost.RouterKey, ibcclient.NewClientUpdateProposalHandler(app.IBCKeeper.ClientKeeper))

	// Create Transfer Keepers
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec, keys[ibctransfertypes.StoreKey], app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
		app.AccountKeeper, app.BankKeeper, scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(app.TransferKeeper)

	app.PooltoyKeeper = pooltoykeeper.NewKeeper(
		appCodec,
		keys[pooltoytypes.StoreKey],
		app.AccountKeeper,
	)

	app.FaucetKeeper = faucetkeeper.NewKeeper(
		app.BankKeeper,
		app.StakingKeeper,
		app.AccountKeeper,
		1,            // amount for mint
		24*time.Hour, // rate limit by time
		keys[faucettypes.StoreKey],
		app.appCodec,
	)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := porttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferModule)
	app.IBCKeeper.SetRouter(ibcRouter)

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, keys[evidencetypes.StoreKey], &app.StakingKeeper, app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper

	// this line is used by starport scaffolding # stargate/app/keeperDefinition

	app.GovKeeper = govkeeper.NewKeeper(
		appCodec, keys[govtypes.StoreKey], app.GetSubspace(govtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, govRouter,
	)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		params.NewAppModule(app.ParamsKeeper),
		transferModule,
		pooltoy.NewAppModule(
			appCodec,
			app.PooltoyKeeper,
			app.BankKeeper,
		),

		faucet.NewAppModule(
			app.FaucetKeeper,
		),
		// this line is used by starport scaffolding # stargate/app/appModule
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		distrtypes.ModuleName, slashingtypes.ModuleName,
		evidencetypes.ModuleName, stakingtypes.ModuleName, minttypes.ModuleName, ibchost.ModuleName,
	)

	app.mm.SetOrderEndBlockers(crisistypes.ModuleName, govtypes.ModuleName, stakingtypes.ModuleName, pooltoytypes.ModuleName, faucettypes.ModuleName)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		ibchost.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		ibctransfertypes.ModuleName,
		minttypes.ModuleName,
		pooltoytypes.ModuleName,
		faucettypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.mm.RegisterServices(module.NewConfigurator(app.MsgServiceRouter(), app.GRPCQueryRouter()))

	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		params.NewAppModule(app.ParamsKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		transferModule,
	)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(
		NewAnteHandler(
			app.AccountKeeper, app.BankKeeper, ante.DefaultSigVerificationGasConsumer,
			encodingConfig.TxConfig.SignModeHandler(),
		),
	)
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}

		// Initialize and seal the capability keeper so all persistent capabilities
		// are loaded in-memory and prevent any further modules from creating scoped
		// sub-keepers.
		// This must be done during creation of baseapp rather than in InitChain so
		// that in-memory capabilities get regenerated on app restart.
		// Note that since this reads from the store, we can only perform it when
		// `loadLatest` is set to true.
		ctx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})
		app.CapabilityKeeper.InitializeAndSeal(ctx)
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper

	return app
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}

	// https://github.com/cosmos/cosmos-sdk/pull/7998
	newDnmRegex := `[a-z][a-z0-9]{2,15}|(?:\x{1F469}\x{200D}\x{2764}\x{FE0F}\x{200D}\x{1F48B}\x{200D}\x{1F468})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{2764}\x{FE0F}\x{200D}\x{1F48B}\x{200D}[\x{1F468}-\x{1F469}])|(?:\x{1F9D1}\x{1F3FB}\x{200D}\x{1F91D}\x{200D}\x{1F9D1}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9D1}\x{1F3FC}\x{200D}\x{1F91D}\x{200D}\x{1F9D1}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9D1}\x{1F3FD}\x{200D}\x{1F91D}\x{200D}\x{1F9D1}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9D1}\x{1F3FE}\x{200D}\x{1F91D}\x{200D}\x{1F9D1}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9D1}\x{1F3FF}\x{200D}\x{1F91D}\x{200D}\x{1F9D1}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F469}\x{1F3FB}\x{200D}\x{1F91D}\x{200D}\x{1F469}[\x{1F3FC}-\x{1F3FF}])|(?:\x{1F469}\x{1F3FC}\x{200D}\x{1F91D}\x{200D}\x{1F469}\x{1F3FB})|(?:\x{1F469}\x{1F3FC}\x{200D}\x{1F91D}\x{200D}\x{1F469}[\x{1F3FD}-\x{1F3FF}])|(?:\x{1F469}\x{1F3FD}\x{200D}\x{1F91D}\x{200D}\x{1F469}[\x{1F3FB}-\x{1F3FC}])|(?:\x{1F469}\x{1F3FD}\x{200D}\x{1F91D}\x{200D}\x{1F469}[\x{1F3FE}-\x{1F3FF}])|(?:\x{1F469}\x{1F3FE}\x{200D}\x{1F91D}\x{200D}\x{1F469}[\x{1F3FB}-\x{1F3FD}])|(?:\x{1F469}\x{1F3FE}\x{200D}\x{1F91D}\x{200D}\x{1F469}\x{1F3FF})|(?:\x{1F469}\x{1F3FF}\x{200D}\x{1F91D}\x{200D}\x{1F469}[\x{1F3FB}-\x{1F3FE}])|(?:\x{1F469}\x{1F3FB}\x{200D}\x{1F91D}\x{200D}\x{1F468}[\x{1F3FC}-\x{1F3FF}])|(?:\x{1F469}\x{1F3FC}\x{200D}\x{1F91D}\x{200D}\x{1F468}\x{1F3FB})|(?:\x{1F469}\x{1F3FC}\x{200D}\x{1F91D}\x{200D}\x{1F468}[\x{1F3FD}-\x{1F3FF}])|(?:\x{1F469}\x{1F3FD}\x{200D}\x{1F91D}\x{200D}\x{1F468}[\x{1F3FB}-\x{1F3FC}])|(?:\x{1F469}\x{1F3FD}\x{200D}\x{1F91D}\x{200D}\x{1F468}[\x{1F3FE}-\x{1F3FF}])|(?:\x{1F469}\x{1F3FE}\x{200D}\x{1F91D}\x{200D}\x{1F468}[\x{1F3FB}-\x{1F3FD}])|(?:\x{1F469}\x{1F3FE}\x{200D}\x{1F91D}\x{200D}\x{1F468}\x{1F3FF})|(?:\x{1F469}\x{1F3FF}\x{200D}\x{1F91D}\x{200D}\x{1F468}[\x{1F3FB}-\x{1F3FE}])|(?:\x{1F468}\x{1F3FB}\x{200D}\x{1F91D}\x{200D}\x{1F468}[\x{1F3FC}-\x{1F3FF}])|(?:\x{1F468}\x{1F3FC}\x{200D}\x{1F91D}\x{200D}\x{1F468}\x{1F3FB})|(?:\x{1F468}\x{1F3FC}\x{200D}\x{1F91D}\x{200D}\x{1F468}[\x{1F3FD}-\x{1F3FF}])|(?:\x{1F468}\x{1F3FD}\x{200D}\x{1F91D}\x{200D}\x{1F468}[\x{1F3FB}-\x{1F3FC}])|(?:\x{1F468}\x{1F3FD}\x{200D}\x{1F91D}\x{200D}\x{1F468}[\x{1F3FE}-\x{1F3FF}])|(?:\x{1F468}\x{1F3FE}\x{200D}\x{1F91D}\x{200D}\x{1F468}[\x{1F3FB}-\x{1F3FD}])|(?:\x{1F468}\x{1F3FE}\x{200D}\x{1F91D}\x{200D}\x{1F468}\x{1F3FF})|(?:\x{1F468}\x{1F3FF}\x{200D}\x{1F91D}\x{200D}\x{1F468}[\x{1F3FB}-\x{1F3FE}])|(?:\x{1F469}\x{200D}\x{2764}\x{200D}\x{1F48B}\x{200D}\x{1F468})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{2764}\x{200D}\x{1F48B}\x{200D}[\x{1F468}-\x{1F469}])|(?:\x{1F468}\x{200D}\x{1F469}\x{200D}\x{1F467}\x{200D}\x{1F466})|(?:\x{1F468}\x{200D}\x{1F469}\x{200D}[\x{1F466}-\x{1F467}]\x{200D}[\x{1F466}-\x{1F467}])|(?:\x{1F468}\x{200D}\x{1F468}\x{200D}\x{1F467}\x{200D}\x{1F466})|(?:\x{1F468}\x{200D}\x{1F468}\x{200D}[\x{1F466}-\x{1F467}]\x{200D}[\x{1F466}-\x{1F467}])|(?:\x{1F469}\x{200D}\x{1F469}\x{200D}\x{1F467}\x{200D}\x{1F466})|(?:\x{1F469}\x{200D}\x{1F469}\x{200D}[\x{1F466}-\x{1F467}]\x{200D}[\x{1F466}-\x{1F467}])|(?:\x{1F3F4}\x{E0067}\x{E0062}\x{E0065}\x{E006E}\x{E0067}\x{E007F})|(?:\x{1F3F4}\x{E0067}\x{E0062}\x{E0073}\x{E0063}\x{E0074}\x{E007F})|(?:\x{1F3F4}\x{E0067}\x{E0062}\x{E0077}\x{E006C}\x{E0073}\x{E007F})|(?:\x{1F469}\x{200D}\x{2764}\x{FE0F}\x{200D}\x{1F468})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{2764}\x{FE0F}\x{200D}[\x{1F468}-\x{1F469}])|(?:\x{1F441}\x{FE0F}\x{200D}\x{1F5E8}\x{FE0F})|(?:\x{1F471}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F471}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F64D}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F64D}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F64E}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F64E}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F645}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F645}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F646}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F646}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F481}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F481}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F64B}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F64B}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9CF}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9CF}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F647}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F647}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F926}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F926}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F937}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F937}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2695}\x{FE0F})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2695}\x{FE0F})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2695}\x{FE0F})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2696}\x{FE0F})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2696}\x{FE0F})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2696}\x{FE0F})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2708}\x{FE0F})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2708}\x{FE0F})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2708}\x{FE0F})|(?:\x{1F46E}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F46E}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F575}\x{FE0F}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F575}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F575}\x{FE0F}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F575}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F482}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F482}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F477}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F477}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F473}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F473}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9B8}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9B8}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9B9}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9B9}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9D9}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9D9}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9DA}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9DA}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9DB}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9DB}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9DC}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9DC}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9DD}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9DD}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F486}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F486}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F487}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F487}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F6B6}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F6B6}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9CD}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9CD}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9CE}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9CE}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F3C3}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F3C3}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9D6}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9D6}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9D7}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9D7}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F3CC}\x{FE0F}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F3CC}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F3CC}\x{FE0F}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F3CC}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F3C4}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F3C4}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F6A3}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F6A3}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F3CA}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F3CA}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{26F9}\x{FE0F}\x{200D}\x{2642}\x{FE0F})|(?:\x{26F9}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{26F9}\x{FE0F}\x{200D}\x{2640}\x{FE0F})|(?:\x{26F9}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F3CB}\x{FE0F}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F3CB}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F3CB}\x{FE0F}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F3CB}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F6B4}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F6B4}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F6B5}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F6B5}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F938}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F938}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F93D}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F93D}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F93E}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F93E}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F939}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F939}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9D8}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9D8}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9D1}\x{200D}\x{1F91D}\x{200D}\x{1F9D1})|(?:\x{1F469}\x{200D}\x{2764}\x{200D}\x{1F468})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{2764}\x{200D}[\x{1F468}-\x{1F469}])|(?:\x{1F468}\x{200D}\x{1F469}\x{200D}[\x{1F466}-\x{1F467}])|(?:\x{1F468}\x{200D}\x{1F468}\x{200D}[\x{1F466}-\x{1F467}])|(?:\x{1F469}\x{200D}\x{1F469}\x{200D}[\x{1F466}-\x{1F467}])|(?:\x{1F468}\x{200D}[\x{1F466}-\x{1F467}]\x{200D}\x{1F466})|(?:\x{1F468}\x{200D}\x{1F467}\x{200D}\x{1F467})|(?:\x{1F469}\x{200D}[\x{1F466}-\x{1F467}]\x{200D}\x{1F466})|(?:\x{1F469}\x{200D}\x{1F467}\x{200D}\x{1F467})|(?:\x{1F441}\x{200D}\x{1F5E8}\x{FE0F})|(?:\x{1F441}\x{FE0F}\x{200D}\x{1F5E8})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9B0})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9B1})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9B3})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9B2})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9B0})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9B0})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9B1})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9B1})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9B3})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9B3})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9B2})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9B2})|(?:\x{1F471}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F471}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F471}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F471}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F64D}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F64D}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F64D}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F64D}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F64E}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F64E}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F64E}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F64E}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F645}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F645}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F645}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F645}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F646}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F646}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F646}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F646}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F481}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F481}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F481}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F481}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F64B}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F64B}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F64B}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F64B}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F9CF}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9CF}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F9CF}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9CF}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F647}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F647}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F647}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F647}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F926}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F926}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F926}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F926}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F937}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F937}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F937}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F937}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F9D1}\x{200D}\x{2695}\x{FE0F})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2695})|(?:\x{1F468}\x{200D}\x{2695}\x{FE0F})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2695})|(?:\x{1F469}\x{200D}\x{2695}\x{FE0F})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2695})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F393})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F393})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F393})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F3EB})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F3EB})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F3EB})|(?:\x{1F9D1}\x{200D}\x{2696}\x{FE0F})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2696})|(?:\x{1F468}\x{200D}\x{2696}\x{FE0F})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2696})|(?:\x{1F469}\x{200D}\x{2696}\x{FE0F})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2696})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F33E})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F33E})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F33E})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F373})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F373})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F373})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F527})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F527})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F527})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F3ED})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F3ED})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F3ED})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F4BC})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F4BC})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F4BC})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F52C})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F52C})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F52C})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F4BB})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F4BB})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F4BB})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F3A4})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F3A4})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F3A4})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F3A8})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F3A8})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F3A8})|(?:\x{1F9D1}\x{200D}\x{2708}\x{FE0F})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2708})|(?:\x{1F468}\x{200D}\x{2708}\x{FE0F})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2708})|(?:\x{1F469}\x{200D}\x{2708}\x{FE0F})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2708})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F680})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F680})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F680})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F692})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F692})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F692})|(?:\x{1F46E}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F46E}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F46E}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F46E}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F575}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F575}\x{FE0F}\x{200D}\x{2642})|(?:\x{1F575}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F575}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F575}\x{FE0F}\x{200D}\x{2640})|(?:\x{1F575}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F482}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F482}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F482}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F482}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F477}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F477}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F477}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F477}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F473}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F473}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F473}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F473}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F9B8}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9B8}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F9B8}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9B8}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F9B9}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9B9}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F9B9}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9B9}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F9D9}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9D9}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F9D9}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9D9}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F9DA}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9DA}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F9DA}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9DA}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F9DB}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9DB}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F9DB}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9DB}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F9DC}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9DC}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F9DC}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9DC}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F9DD}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9DD}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F9DD}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9DD}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F9DE}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9DE}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9DF}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9DF}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F486}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F486}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F486}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F486}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F487}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F487}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F487}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F487}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F6B6}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F6B6}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F6B6}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F6B6}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F9CD}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9CD}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F9CD}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9CD}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F9CE}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9CE}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F9CE}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9CE}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9AF})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9AF})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9AF})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9BC})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9BC})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9BC})|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9BD})|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9BD})|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9BD})|(?:\x{1F3C3}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F3C3}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F3C3}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F3C3}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F46F}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F46F}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9D6}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9D6}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F9D6}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9D6}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F9D7}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9D7}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F9D7}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9D7}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F3CC}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F3CC}\x{FE0F}\x{200D}\x{2642})|(?:\x{1F3CC}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F3CC}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F3CC}\x{FE0F}\x{200D}\x{2640})|(?:\x{1F3CC}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F3C4}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F3C4}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F3C4}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F3C4}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F6A3}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F6A3}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F6A3}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F6A3}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F3CA}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F3CA}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F3CA}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F3CA}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{26F9}\x{200D}\x{2642}\x{FE0F})|(?:\x{26F9}\x{FE0F}\x{200D}\x{2642})|(?:\x{26F9}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{26F9}\x{200D}\x{2640}\x{FE0F})|(?:\x{26F9}\x{FE0F}\x{200D}\x{2640})|(?:\x{26F9}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F3CB}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F3CB}\x{FE0F}\x{200D}\x{2642})|(?:\x{1F3CB}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F3CB}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F3CB}\x{FE0F}\x{200D}\x{2640})|(?:\x{1F3CB}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F6B4}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F6B4}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F6B4}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F6B4}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F6B5}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F6B5}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F6B5}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F6B5}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F938}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F938}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F938}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F938}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F93C}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F93C}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F93D}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F93D}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F93D}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F93D}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F93E}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F93E}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F93E}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F93E}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F939}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F939}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F939}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F939}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F9D8}\x{200D}\x{2642}\x{FE0F})|(?:\x{1F9D8}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2642})|(?:\x{1F9D8}\x{200D}\x{2640}\x{FE0F})|(?:\x{1F9D8}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{2640})|(?:\x{1F3F3}\x{FE0F}\x{200D}\x{1F308})|(?:\x{1F3F4}\x{200D}\x{2620}\x{FE0F})|(?:\x{1F441}\x{200D}\x{1F5E8})|(?:\x{1F468}\x{200D}[\x{1F9B0}-\x{1F9B1}])|(?:\x{1F468}\x{200D}\x{1F9B3})|(?:\x{1F468}\x{200D}\x{1F9B2})|(?:\x{1F469}\x{200D}\x{1F9B0})|(?:\x{1F9D1}\x{200D}\x{1F9B0})|(?:\x{1F469}\x{200D}\x{1F9B1})|(?:\x{1F9D1}\x{200D}\x{1F9B1})|(?:\x{1F469}\x{200D}\x{1F9B3})|(?:\x{1F9D1}\x{200D}\x{1F9B3})|(?:\x{1F469}\x{200D}\x{1F9B2})|(?:\x{1F9D1}\x{200D}\x{1F9B2})|(?:\x{1F471}\x{200D}\x{2640})|(?:\x{1F471}\x{200D}\x{2642})|(?:\x{1F64D}\x{200D}\x{2642})|(?:\x{1F64D}\x{200D}\x{2640})|(?:\x{1F64E}\x{200D}\x{2642})|(?:\x{1F64E}\x{200D}\x{2640})|(?:\x{1F645}\x{200D}\x{2642})|(?:\x{1F645}\x{200D}\x{2640})|(?:\x{1F646}\x{200D}\x{2642})|(?:\x{1F646}\x{200D}\x{2640})|(?:\x{1F481}\x{200D}\x{2642})|(?:\x{1F481}\x{200D}\x{2640})|(?:\x{1F64B}\x{200D}\x{2642})|(?:\x{1F64B}\x{200D}\x{2640})|(?:\x{1F9CF}\x{200D}\x{2642})|(?:\x{1F9CF}\x{200D}\x{2640})|(?:\x{1F647}\x{200D}\x{2642})|(?:\x{1F647}\x{200D}\x{2640})|(?:\x{1F926}\x{200D}\x{2642})|(?:\x{1F926}\x{200D}\x{2640})|(?:\x{1F937}\x{200D}\x{2642})|(?:\x{1F937}\x{200D}\x{2640})|(?:\x{1F9D1}\x{200D}\x{2695})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{2695})|(?:\x{1F9D1}\x{200D}\x{1F393})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F393})|(?:\x{1F9D1}\x{200D}\x{1F3EB})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F3EB})|(?:\x{1F9D1}\x{200D}\x{2696})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{2696})|(?:\x{1F9D1}\x{200D}\x{1F33E})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F33E})|(?:\x{1F9D1}\x{200D}\x{1F373})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F373})|(?:\x{1F9D1}\x{200D}\x{1F527})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F527})|(?:\x{1F9D1}\x{200D}\x{1F3ED})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F3ED})|(?:\x{1F9D1}\x{200D}\x{1F4BC})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F4BC})|(?:\x{1F9D1}\x{200D}\x{1F52C})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F52C})|(?:\x{1F9D1}\x{200D}\x{1F4BB})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F4BB})|(?:\x{1F9D1}\x{200D}\x{1F3A4})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F3A4})|(?:\x{1F9D1}\x{200D}\x{1F3A8})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F3A8})|(?:\x{1F9D1}\x{200D}\x{2708})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{2708})|(?:\x{1F9D1}\x{200D}\x{1F680})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F680})|(?:\x{1F9D1}\x{200D}\x{1F692})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F692})|(?:\x{1F46E}\x{200D}\x{2642})|(?:\x{1F46E}\x{200D}\x{2640})|(?:\x{1F575}\x{200D}\x{2642})|(?:\x{1F575}\x{200D}\x{2640})|(?:\x{1F482}\x{200D}\x{2642})|(?:\x{1F482}\x{200D}\x{2640})|(?:\x{1F477}\x{200D}\x{2642})|(?:\x{1F477}\x{200D}\x{2640})|(?:\x{1F473}\x{200D}\x{2642})|(?:\x{1F473}\x{200D}\x{2640})|(?:\x{1F9B8}\x{200D}\x{2642})|(?:\x{1F9B8}\x{200D}\x{2640})|(?:\x{1F9B9}\x{200D}\x{2642})|(?:\x{1F9B9}\x{200D}\x{2640})|(?:\x{1F9D9}\x{200D}\x{2642})|(?:\x{1F9D9}\x{200D}\x{2640})|(?:\x{1F9DA}\x{200D}\x{2642})|(?:\x{1F9DA}\x{200D}\x{2640})|(?:\x{1F9DB}\x{200D}\x{2642})|(?:\x{1F9DB}\x{200D}\x{2640})|(?:\x{1F9DC}\x{200D}\x{2642})|(?:\x{1F9DC}\x{200D}\x{2640})|(?:\x{1F9DD}\x{200D}\x{2642})|(?:\x{1F9DD}\x{200D}\x{2640})|(?:\x{1F9DE}\x{200D}\x{2642})|(?:\x{1F9DE}\x{200D}\x{2640})|(?:\x{1F9DF}\x{200D}\x{2642})|(?:\x{1F9DF}\x{200D}\x{2640})|(?:\x{1F486}\x{200D}\x{2642})|(?:\x{1F486}\x{200D}\x{2640})|(?:\x{1F487}\x{200D}\x{2642})|(?:\x{1F487}\x{200D}\x{2640})|(?:\x{1F6B6}\x{200D}\x{2642})|(?:\x{1F6B6}\x{200D}\x{2640})|(?:\x{1F9CD}\x{200D}\x{2642})|(?:\x{1F9CD}\x{200D}\x{2640})|(?:\x{1F9CE}\x{200D}\x{2642})|(?:\x{1F9CE}\x{200D}\x{2640})|(?:\x{1F9D1}\x{200D}\x{1F9AF})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F9AF})|(?:\x{1F9D1}\x{200D}\x{1F9BC})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F9BC})|(?:\x{1F9D1}\x{200D}\x{1F9BD})|(?:[\x{1F468}-\x{1F469}]\x{200D}\x{1F9BD})|(?:\x{1F3C3}\x{200D}\x{2642})|(?:\x{1F3C3}\x{200D}\x{2640})|(?:\x{1F46F}\x{200D}\x{2642})|(?:\x{1F46F}\x{200D}\x{2640})|(?:\x{1F9D6}\x{200D}\x{2642})|(?:\x{1F9D6}\x{200D}\x{2640})|(?:\x{1F9D7}\x{200D}\x{2642})|(?:\x{1F9D7}\x{200D}\x{2640})|(?:\x{1F3CC}\x{200D}\x{2642})|(?:\x{1F3CC}\x{200D}\x{2640})|(?:\x{1F3C4}\x{200D}\x{2642})|(?:\x{1F3C4}\x{200D}\x{2640})|(?:\x{1F6A3}\x{200D}\x{2642})|(?:\x{1F6A3}\x{200D}\x{2640})|(?:\x{1F3CA}\x{200D}\x{2642})|(?:\x{1F3CA}\x{200D}\x{2640})|(?:\x{26F9}\x{200D}\x{2642})|(?:\x{26F9}\x{200D}\x{2640})|(?:\x{1F3CB}\x{200D}\x{2642})|(?:\x{1F3CB}\x{200D}\x{2640})|(?:\x{1F6B4}\x{200D}\x{2642})|(?:\x{1F6B4}\x{200D}\x{2640})|(?:\x{1F6B5}\x{200D}\x{2642})|(?:\x{1F6B5}\x{200D}\x{2640})|(?:\x{1F938}\x{200D}\x{2642})|(?:\x{1F938}\x{200D}\x{2640})|(?:\x{1F93C}\x{200D}\x{2642})|(?:\x{1F93C}\x{200D}\x{2640})|(?:\x{1F93D}\x{200D}\x{2642})|(?:\x{1F93D}\x{200D}\x{2640})|(?:\x{1F93E}\x{200D}\x{2642})|(?:\x{1F93E}\x{200D}\x{2640})|(?:\x{1F939}\x{200D}\x{2642})|(?:\x{1F939}\x{200D}\x{2640})|(?:\x{1F9D8}\x{200D}\x{2642})|(?:\x{1F9D8}\x{200D}\x{2640})|(?:\x{1F468}\x{200D}[\x{1F466}-\x{1F467}])|(?:\x{1F469}\x{200D}[\x{1F466}-\x{1F467}])|(?:\x{1F415}\x{200D}\x{1F9BA})|(?:\x{0023}\x{FE0F}\x{20E3})|(?:\x{002A}\x{FE0F}\x{20E3})|(?:[\x{0030}-\x{0039}]\x{FE0F}\x{20E3})|(?:\x{1F3F3}\x{200D}\x{1F308})|(?:\x{1F3F4}\x{200D}\x{2620})|(?:\x{263A}\x{FE0F})|(?:\x{2639}\x{FE0F})|(?:\x{2620}\x{FE0F})|(?:[\x{2763}-\x{2764}]\x{FE0F})|(?:\x{1F573}\x{FE0F})|(?:\x{1F5E8}\x{FE0F})|(?:\x{1F5EF}\x{FE0F})|(?:\x{1F44B}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F91A}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F590}\x{FE0F})|(?:\x{1F590}[\x{1F3FB}-\x{1F3FF}])|(?:\x{270B}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F596}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F44C}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F90F}[\x{1F3FB}-\x{1F3FF}])|(?:\x{270C}\x{FE0F})|(?:\x{270C}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F91E}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F91F}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F918}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F919}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F448}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F449}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F446}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F595}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F447}[\x{1F3FB}-\x{1F3FF}])|(?:\x{261D}\x{FE0F})|(?:\x{261D}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F44D}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F44E}[\x{1F3FB}-\x{1F3FF}])|(?:\x{270A}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F44A}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F91B}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F91C}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F44F}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F64C}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F450}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F932}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F64F}[\x{1F3FB}-\x{1F3FF}])|(?:\x{270D}\x{FE0F})|(?:\x{270D}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F485}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F933}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F4AA}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9B5}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9B6}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F442}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9BB}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F443}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F441}\x{FE0F})|(?:\x{1F476}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9D2}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F466}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F467}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9D1}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F471}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9D4}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F469}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9D3}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F474}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F475}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F64D}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F64E}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F645}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F646}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F481}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F64B}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9CF}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F647}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F926}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F937}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F46E}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F575}\x{FE0F})|(?:\x{1F575}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F482}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F477}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F934}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F478}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F473}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F472}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9D5}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F935}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F470}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F930}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F931}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F47C}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F385}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F936}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9B8}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9B9}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9D9}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9DA}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9DB}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9DC}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9DD}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F486}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F487}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F6B6}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9CD}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9CE}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F3C3}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F483}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F57A}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F574}\x{FE0F})|(?:\x{1F574}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9D6}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9D7}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F3C7}[\x{1F3FB}-\x{1F3FF}])|(?:\x{26F7}\x{FE0F})|(?:\x{1F3C2}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F3CC}\x{FE0F})|(?:\x{1F3CC}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F3C4}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F6A3}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F3CA}[\x{1F3FB}-\x{1F3FF}])|(?:\x{26F9}\x{FE0F})|(?:\x{26F9}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F3CB}\x{FE0F})|(?:\x{1F3CB}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F6B4}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F6B5}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F938}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F93D}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F93E}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F939}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F9D8}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F6C0}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F6CC}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F46D}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F46B}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F46C}[\x{1F3FB}-\x{1F3FF}])|(?:\x{1F5E3}\x{FE0F})|(?:\x{1F43F}\x{FE0F})|(?:\x{1F54A}\x{FE0F})|(?:[\x{1F577}-\x{1F578}]\x{FE0F})|(?:\x{1F3F5}\x{FE0F})|(?:\x{2618}\x{FE0F})|(?:\x{1F336}\x{FE0F})|(?:\x{1F37D}\x{FE0F})|(?:\x{1F5FA}\x{FE0F})|(?:\x{1F3D4}\x{FE0F})|(?:\x{26F0}\x{FE0F})|(?:[\x{1F3D5}-\x{1F3D6}]\x{FE0F})|(?:[\x{1F3DC}-\x{1F3DF}]\x{FE0F})|(?:\x{1F3DB}\x{FE0F})|(?:[\x{1F3D7}-\x{1F3D8}]\x{FE0F})|(?:\x{1F3DA}\x{FE0F})|(?:\x{26E9}\x{FE0F})|(?:\x{1F3D9}\x{FE0F})|(?:\x{2668}\x{FE0F})|(?:\x{1F3CE}\x{FE0F})|(?:\x{1F3CD}\x{FE0F})|(?:[\x{1F6E3}-\x{1F6E4}]\x{FE0F})|(?:\x{1F6E2}\x{FE0F})|(?:\x{1F6F3}\x{FE0F})|(?:\x{26F4}\x{FE0F})|(?:\x{1F6E5}\x{FE0F})|(?:\x{2708}\x{FE0F})|(?:\x{1F6E9}\x{FE0F})|(?:\x{1F6F0}\x{FE0F})|(?:\x{1F6CE}\x{FE0F})|(?:[\x{23F1}-\x{23F2}]\x{FE0F})|(?:\x{1F570}\x{FE0F})|(?:\x{1F321}\x{FE0F})|(?:[\x{2600}-\x{2601}]\x{FE0F})|(?:\x{26C8}\x{FE0F})|(?:[\x{1F324}-\x{1F32C}]\x{FE0F})|(?:\x{2602}\x{FE0F})|(?:\x{26F1}\x{FE0F})|(?:\x{2744}\x{FE0F})|(?:[\x{2603}-\x{2604}]\x{FE0F})|(?:\x{1F397}\x{FE0F})|(?:\x{1F39F}\x{FE0F})|(?:\x{1F396}\x{FE0F})|(?:\x{26F8}\x{FE0F})|(?:\x{1F579}\x{FE0F})|(?:\x{2660}\x{FE0F})|(?:[\x{2665}-\x{2666}]\x{FE0F})|(?:\x{2663}\x{FE0F})|(?:\x{265F}\x{FE0F})|(?:\x{1F5BC}\x{FE0F})|(?:\x{1F576}\x{FE0F})|(?:\x{1F6CD}\x{FE0F})|(?:\x{26D1}\x{FE0F})|(?:[\x{1F399}-\x{1F39B}]\x{FE0F})|(?:\x{260E}\x{FE0F})|(?:\x{1F5A5}\x{FE0F})|(?:\x{1F5A8}\x{FE0F})|(?:\x{2328}\x{FE0F})|(?:[\x{1F5B1}-\x{1F5B2}]\x{FE0F})|(?:\x{1F39E}\x{FE0F})|(?:\x{1F4FD}\x{FE0F})|(?:\x{1F56F}\x{FE0F})|(?:\x{1F5DE}\x{FE0F})|(?:\x{1F3F7}\x{FE0F})|(?:\x{2709}\x{FE0F})|(?:\x{1F5F3}\x{FE0F})|(?:\x{270F}\x{FE0F})|(?:\x{2712}\x{FE0F})|(?:\x{1F58B}\x{FE0F})|(?:\x{1F58A}\x{FE0F})|(?:[\x{1F58C}-\x{1F58D}]\x{FE0F})|(?:\x{1F5C2}\x{FE0F})|(?:[\x{1F5D2}-\x{1F5D3}]\x{FE0F})|(?:\x{1F587}\x{FE0F})|(?:\x{2702}\x{FE0F})|(?:[\x{1F5C3}-\x{1F5C4}]\x{FE0F})|(?:\x{1F5D1}\x{FE0F})|(?:\x{1F5DD}\x{FE0F})|(?:\x{26CF}\x{FE0F})|(?:\x{2692}\x{FE0F})|(?:\x{1F6E0}\x{FE0F})|(?:\x{1F5E1}\x{FE0F})|(?:\x{2694}\x{FE0F})|(?:\x{1F6E1}\x{FE0F})|(?:\x{2699}\x{FE0F})|(?:\x{1F5DC}\x{FE0F})|(?:\x{2696}\x{FE0F})|(?:\x{26D3}\x{FE0F})|(?:\x{2697}\x{FE0F})|(?:\x{1F6CF}\x{FE0F})|(?:\x{1F6CB}\x{FE0F})|(?:[\x{26B0}-\x{26B1}]\x{FE0F})|(?:\x{26A0}\x{FE0F})|(?:[\x{2622}-\x{2623}]\x{FE0F})|(?:\x{2B06}\x{FE0F})|(?:\x{2197}\x{FE0F})|(?:\x{27A1}\x{FE0F})|(?:\x{2198}\x{FE0F})|(?:\x{2B07}\x{FE0F})|(?:\x{2199}\x{FE0F})|(?:\x{2B05}\x{FE0F})|(?:\x{2196}\x{FE0F})|(?:\x{2195}\x{FE0F})|(?:\x{2194}\x{FE0F})|(?:[\x{21A9}-\x{21AA}]\x{FE0F})|(?:[\x{2934}-\x{2935}]\x{FE0F})|(?:\x{269B}\x{FE0F})|(?:\x{1F549}\x{FE0F})|(?:\x{2721}\x{FE0F})|(?:\x{2638}\x{FE0F})|(?:\x{262F}\x{FE0F})|(?:\x{271D}\x{FE0F})|(?:\x{2626}\x{FE0F})|(?:\x{262A}\x{FE0F})|(?:\x{262E}\x{FE0F})|(?:\x{25B6}\x{FE0F})|(?:\x{23ED}\x{FE0F})|(?:\x{23EF}\x{FE0F})|(?:\x{25C0}\x{FE0F})|(?:\x{23EE}\x{FE0F})|(?:[\x{23F8}-\x{23FA}]\x{FE0F})|(?:\x{23CF}\x{FE0F})|(?:\x{2640}\x{FE0F})|(?:\x{2642}\x{FE0F})|(?:\x{2695}\x{FE0F})|(?:\x{267E}\x{FE0F})|(?:\x{267B}\x{FE0F})|(?:\x{269C}\x{FE0F})|(?:\x{2611}\x{FE0F})|(?:\x{2714}\x{FE0F})|(?:\x{2716}\x{FE0F})|(?:\x{303D}\x{FE0F})|(?:[\x{2733}-\x{2734}]\x{FE0F})|(?:\x{2747}\x{FE0F})|(?:\x{203C}\x{FE0F})|(?:\x{2049}\x{FE0F})|(?:\x{3030}\x{FE0F})|(?:\x{00A9}\x{FE0F})|(?:\x{00AE}\x{FE0F})|(?:\x{2122}\x{FE0F})|(?:\x{0023}\x{20E3})|(?:\x{002A}\x{20E3})|(?:[\x{0030}-\x{0039}]\x{20E3})|(?:[\x{1F170}-\x{1F171}]\x{FE0F})|(?:\x{2139}\x{FE0F})|(?:\x{24C2}\x{FE0F})|(?:[\x{1F17E}-\x{1F17F}]\x{FE0F})|(?:\x{1F202}\x{FE0F})|(?:\x{1F237}\x{FE0F})|(?:\x{3297}\x{FE0F})|(?:\x{3299}\x{FE0F})|(?:\x{25FC}\x{FE0F})|(?:\x{25FB}\x{FE0F})|(?:[\x{25AA}-\x{25AB}]\x{FE0F})|(?:\x{1F3F3}\x{FE0F})|(?:\x{1F1E6}[\x{1F1E8}-\x{1F1EC}])|(?:\x{1F1E6}\x{1F1EE})|(?:\x{1F1E6}[\x{1F1F1}-\x{1F1F2}])|(?:\x{1F1E6}\x{1F1F4})|(?:\x{1F1E6}[\x{1F1F6}-\x{1F1FA}])|(?:\x{1F1E6}[\x{1F1FC}-\x{1F1FD}])|(?:\x{1F1E6}\x{1F1FF})|(?:\x{1F1E7}[\x{1F1E6}-\x{1F1E7}])|(?:\x{1F1E7}[\x{1F1E9}-\x{1F1EF}])|(?:\x{1F1E7}[\x{1F1F1}-\x{1F1F4}])|(?:\x{1F1E7}[\x{1F1F6}-\x{1F1F9}])|(?:\x{1F1E7}[\x{1F1FB}-\x{1F1FC}])|(?:\x{1F1E7}[\x{1F1FE}-\x{1F1FF}])|(?:\x{1F1E8}\x{1F1E6})|(?:\x{1F1E8}[\x{1F1E8}-\x{1F1E9}])|(?:\x{1F1E8}[\x{1F1EB}-\x{1F1EE}])|(?:\x{1F1E8}[\x{1F1F0}-\x{1F1F5}])|(?:\x{1F1E8}\x{1F1F7})|(?:\x{1F1E8}[\x{1F1FA}-\x{1F1FF}])|(?:\x{1F1E9}\x{1F1EA})|(?:\x{1F1E9}\x{1F1EC})|(?:\x{1F1E9}[\x{1F1EF}-\x{1F1F0}])|(?:\x{1F1E9}\x{1F1F2})|(?:\x{1F1E9}\x{1F1F4})|(?:\x{1F1E9}\x{1F1FF})|(?:\x{1F1EA}\x{1F1E6})|(?:\x{1F1EA}\x{1F1E8})|(?:\x{1F1EA}\x{1F1EA})|(?:\x{1F1EA}[\x{1F1EC}-\x{1F1ED}])|(?:\x{1F1EA}[\x{1F1F7}-\x{1F1FA}])|(?:\x{1F1EB}[\x{1F1EE}-\x{1F1F0}])|(?:\x{1F1EB}\x{1F1F2})|(?:\x{1F1EB}\x{1F1F4})|(?:\x{1F1EB}\x{1F1F7})|(?:\x{1F1EC}[\x{1F1E6}-\x{1F1E7}])|(?:\x{1F1EC}[\x{1F1E9}-\x{1F1EE}])|(?:\x{1F1EC}[\x{1F1F1}-\x{1F1F3}])|(?:\x{1F1EC}[\x{1F1F5}-\x{1F1FA}])|(?:\x{1F1EC}\x{1F1FC})|(?:\x{1F1EC}\x{1F1FE})|(?:\x{1F1ED}\x{1F1F0})|(?:\x{1F1ED}[\x{1F1F2}-\x{1F1F3}])|(?:\x{1F1ED}\x{1F1F7})|(?:\x{1F1ED}[\x{1F1F9}-\x{1F1FA}])|(?:\x{1F1EE}[\x{1F1E8}-\x{1F1EA}])|(?:\x{1F1EE}[\x{1F1F1}-\x{1F1F4}])|(?:\x{1F1EE}[\x{1F1F6}-\x{1F1F9}])|(?:\x{1F1EF}\x{1F1EA})|(?:\x{1F1EF}\x{1F1F2})|(?:\x{1F1EF}[\x{1F1F4}-\x{1F1F5}])|(?:\x{1F1F0}\x{1F1EA})|(?:\x{1F1F0}[\x{1F1EC}-\x{1F1EE}])|(?:\x{1F1F0}[\x{1F1F2}-\x{1F1F3}])|(?:\x{1F1F0}\x{1F1F5})|(?:\x{1F1F0}\x{1F1F7})|(?:\x{1F1F0}\x{1F1FC})|(?:\x{1F1F0}[\x{1F1FE}-\x{1F1FF}])|(?:\x{1F1F1}[\x{1F1E6}-\x{1F1E8}])|(?:\x{1F1F1}\x{1F1EE})|(?:\x{1F1F1}\x{1F1F0})|(?:\x{1F1F1}[\x{1F1F7}-\x{1F1FB}])|(?:\x{1F1F1}\x{1F1FE})|(?:\x{1F1F2}\x{1F1E6})|(?:\x{1F1F2}[\x{1F1E8}-\x{1F1ED}])|(?:\x{1F1F2}[\x{1F1F0}-\x{1F1FF}])|(?:\x{1F1F3}\x{1F1E6})|(?:\x{1F1F3}\x{1F1E8})|(?:\x{1F1F3}[\x{1F1EA}-\x{1F1EC}])|(?:\x{1F1F3}\x{1F1EE})|(?:\x{1F1F3}\x{1F1F1})|(?:\x{1F1F3}[\x{1F1F4}-\x{1F1F5}])|(?:\x{1F1F3}\x{1F1F7})|(?:\x{1F1F3}\x{1F1FA})|(?:\x{1F1F3}\x{1F1FF})|(?:\x{1F1F4}\x{1F1F2})|(?:\x{1F1F5}\x{1F1E6})|(?:\x{1F1F5}[\x{1F1EA}-\x{1F1ED}])|(?:\x{1F1F5}[\x{1F1F0}-\x{1F1F3}])|(?:\x{1F1F5}[\x{1F1F7}-\x{1F1F9}])|(?:\x{1F1F5}\x{1F1FC})|(?:\x{1F1F5}\x{1F1FE})|(?:\x{1F1F6}\x{1F1E6})|(?:\x{1F1F7}\x{1F1EA})|(?:\x{1F1F7}\x{1F1F4})|(?:\x{1F1F7}\x{1F1F8})|(?:\x{1F1F7}\x{1F1FA})|(?:\x{1F1F7}\x{1F1FC})|(?:\x{1F1F8}[\x{1F1E6}-\x{1F1EA}])|(?:\x{1F1F8}[\x{1F1EC}-\x{1F1F4}])|(?:\x{1F1F8}[\x{1F1F7}-\x{1F1F9}])|(?:\x{1F1F8}\x{1F1FB})|(?:\x{1F1F8}[\x{1F1FD}-\x{1F1FF}])|(?:\x{1F1F9}\x{1F1E6})|(?:\x{1F1F9}[\x{1F1E8}-\x{1F1E9}])|(?:\x{1F1F9}[\x{1F1EB}-\x{1F1ED}])|(?:\x{1F1F9}[\x{1F1EF}-\x{1F1F4}])|(?:\x{1F1F9}\x{1F1F7})|(?:\x{1F1F9}\x{1F1F9})|(?:\x{1F1F9}[\x{1F1FB}-\x{1F1FC}])|(?:\x{1F1F9}\x{1F1FF})|(?:\x{1F1FA}\x{1F1E6})|(?:\x{1F1FA}\x{1F1EC})|(?:\x{1F1FA}[\x{1F1F2}-\x{1F1F3}])|(?:\x{1F1FA}\x{1F1F8})|(?:\x{1F1FA}[\x{1F1FE}-\x{1F1FF}])|(?:\x{1F1FB}\x{1F1E6})|(?:\x{1F1FB}\x{1F1E8})|(?:\x{1F1FB}\x{1F1EA})|(?:\x{1F1FB}\x{1F1EC})|(?:\x{1F1FB}\x{1F1EE})|(?:\x{1F1FB}\x{1F1F3})|(?:\x{1F1FB}\x{1F1FA})|(?:\x{1F1FC}\x{1F1EB})|(?:\x{1F1FC}\x{1F1F8})|(?:\x{1F1FD}\x{1F1F0})|(?:\x{1F1FE}\x{1F1EA})|(?:\x{1F1FE}\x{1F1F9})|(?:\x{1F1FF}\x{1F1E6})|(?:\x{1F1FF}\x{1F1F2})|(?:\x{1F1FF}\x{1F1FC})|\x{00A9}|\x{00AE}|\x{1F004}|\x{1F0CF}|[\x{1F170}-\x{1F171}]|[\x{1F17E}-\x{1F17F}]|\x{1F18E}|[\x{1F191}-\x{1F19A}]|[\x{1F201}-\x{1F202}]|\x{1F21A}|\x{1F22F}|[\x{1F232}-\x{1F23A}]|[\x{1F250}-\x{1F251}]|[\x{1F300}-\x{1F321}]|[\x{1F324}-\x{1F393}]|[\x{1F396}-\x{1F397}]|[\x{1F399}-\x{1F39B}]|[\x{1F39E}-\x{1F3F0}]|[\x{1F3F3}-\x{1F3F5}]|[\x{1F3F7}-\x{1F4FD}]|[\x{1F4FF}-\x{1F53D}]|[\x{1F549}-\x{1F54E}]|[\x{1F550}-\x{1F567}]|[\x{1F56F}-\x{1F570}]|[\x{1F573}-\x{1F57A}]|\x{1F587}|[\x{1F58A}-\x{1F58D}]|\x{1F590}|[\x{1F595}-\x{1F596}]|[\x{1F5A4}-\x{1F5A5}]|\x{1F5A8}|[\x{1F5B1}-\x{1F5B2}]|\x{1F5BC}|[\x{1F5C2}-\x{1F5C4}]|[\x{1F5D1}-\x{1F5D3}]|[\x{1F5DC}-\x{1F5DE}]|\x{1F5E1}|\x{1F5E3}|\x{1F5E8}|\x{1F5EF}|\x{1F5F3}|[\x{1F5FA}-\x{1F64F}]|[\x{1F680}-\x{1F6C5}]|[\x{1F6CB}-\x{1F6D2}]|\x{1F6D5}|[\x{1F6E0}-\x{1F6E5}]|\x{1F6E9}|[\x{1F6EB}-\x{1F6EC}]|\x{1F6F0}|[\x{1F6F3}-\x{1F6FA}]|[\x{1F7E0}-\x{1F7EB}]|[\x{1F90D}-\x{1F93A}]|[\x{1F93C}-\x{1F945}]|[\x{1F947}-\x{1F971}]|[\x{1F973}-\x{1F976}]|[\x{1F97A}-\x{1F9A2}]|[\x{1F9A5}-\x{1F9AA}]|[\x{1F9AE}-\x{1F9CA}]|[\x{1F9CD}-\x{1F9FF}]|[\x{1FA70}-\x{1FA73}]|[\x{1FA78}-\x{1FA7A}]|[\x{1FA80}-\x{1FA82}]|[\x{1FA90}-\x{1FA95}]|\x{203C}|\x{2049}|\x{2122}|\x{2139}|[\x{2194}-\x{2199}]|[\x{21A9}-\x{21AA}]|[\x{231A}-\x{231B}]|\x{2328}|\x{23CF}|[\x{23E9}-\x{23F3}]|[\x{23F8}-\x{23FA}]|\x{24C2}|[\x{25AA}-\x{25AB}]|\x{25B6}|\x{25C0}|[\x{25FB}-\x{25FE}]|[\x{2600}-\x{2604}]|\x{260E}|\x{2611}|[\x{2614}-\x{2615}]|\x{2618}|\x{261D}|\x{2620}|[\x{2622}-\x{2623}]|\x{2626}|\x{262A}|[\x{262E}-\x{262F}]|[\x{2638}-\x{263A}]|\x{2640}|\x{2642}|[\x{2648}-\x{2653}]|[\x{265F}-\x{2660}]|\x{2663}|[\x{2665}-\x{2666}]|\x{2668}|\x{267B}|[\x{267E}-\x{267F}]|[\x{2692}-\x{2697}]|\x{2699}|[\x{269B}-\x{269C}]|[\x{26A0}-\x{26A1}]|[\x{26AA}-\x{26AB}]|[\x{26B0}-\x{26B1}]|[\x{26BD}-\x{26BE}]|[\x{26C4}-\x{26C5}]|\x{26C8}|[\x{26CE}-\x{26CF}]|\x{26D1}|[\x{26D3}-\x{26D4}]|[\x{26E9}-\x{26EA}]|[\x{26F0}-\x{26F5}]|[\x{26F7}-\x{26FA}]|\x{26FD}|\x{2702}|\x{2705}|[\x{2708}-\x{270D}]|\x{270F}|\x{2712}|\x{2714}|\x{2716}|\x{271D}|\x{2721}|\x{2728}|[\x{2733}-\x{2734}]|\x{2744}|\x{2747}|\x{274C}|\x{274E}|[\x{2753}-\x{2755}]|\x{2757}|[\x{2763}-\x{2764}]|[\x{2795}-\x{2797}]|\x{27A1}|\x{27B0}|\x{27BF}|[\x{2934}-\x{2935}]|[\x{2B05}-\x{2B07}]|[\x{2B1B}-\x{2B1C}]|\x{2B50}|\x{2B55}|\x{3030}|\x{303D}|\x{3297}|\x{3299}`
	sdk.SetCoinDenomRegex(func() string {
		return newDnmRegex
	})

	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// BlockedAddrs returns all the app's module account addresses that are not
// allowed to receive external tokens.
func (app *App) BlockedAddrs() map[string]bool {
	blockedAddrs := make(map[string]bool)
	for acc := range maccPerms {
		blockedAddrs[authtypes.NewModuleAddress(acc).String()] = !allowedReceivingModAcc[acc]
	}

	return blockedAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns Gaia's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Marshaler {
	return app.appCodec
}

// InterfaceRegistry returns Gaia's InterfaceRegistry
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *App) GetMemKey(storeKey string) *sdk.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Register legacy tx routes.
	authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
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
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	// this line is used by starport scaffolding # stargate/app/paramSubspace

	return paramsKeeper
}

func (*App) OnTxSucceeded(ctx sdk.Context, sourcePort, sourceChannel string, txHash []byte, txBytes []byte) {
}

func (*App) OnTxFailed(ctx sdk.Context, sourcePort, sourceChannel string, txHash []byte, txBytes []byte) {
}
