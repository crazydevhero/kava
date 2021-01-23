package hard

// DO NOT EDIT - generated by aliasgen tool (github.com/rhuairahrighairidh/aliasgen)

import (
	"github.com/kava-labs/kava/x/hard/keeper"
	"github.com/kava-labs/kava/x/hard/types"
)

const (
	AttributeKeyBlockHeight            = types.AttributeKeyBlockHeight
	AttributeKeyBorrow                 = types.AttributeKeyBorrow
	AttributeKeyBorrowCoins            = types.AttributeKeyBorrowCoins
	AttributeKeyBorrower               = types.AttributeKeyBorrower
	AttributeKeyDeposit                = types.AttributeKeyDeposit
	AttributeKeyDepositCoins           = types.AttributeKeyDepositCoins
	AttributeKeyDepositDenom           = types.AttributeKeyDepositDenom
	AttributeKeyDepositor              = types.AttributeKeyDepositor
	AttributeKeyRepayCoins             = types.AttributeKeyRepayCoins
	AttributeKeyRewardsDistribution    = types.AttributeKeyRewardsDistribution
	AttributeKeySender                 = types.AttributeKeySender
	AttributeValueCategory             = types.AttributeValueCategory
	DefaultParamspace                  = types.DefaultParamspace
	DelegatorAccount                   = types.DelegatorAccount
	EventTypeDeleteHardDeposit         = types.EventTypeDeleteHardDeposit
	EventTypeDepositLiquidation        = types.EventTypeDepositLiquidation
	EventTypeHardBorrow                = types.EventTypeHardBorrow
	EventTypeHardDelegatorDistribution = types.EventTypeHardDelegatorDistribution
	EventTypeHardDeposit               = types.EventTypeHardDeposit
	EventTypeHardLPDistribution        = types.EventTypeHardLPDistribution
	EventTypeHardRepay                 = types.EventTypeHardRepay
	EventTypeHardWithdrawal            = types.EventTypeHardWithdrawal
	LPAccount                          = types.LPAccount
	LiquidatorAccount                  = types.LiquidatorAccount
	ModuleAccountName                  = types.ModuleAccountName
	ModuleName                         = types.ModuleName
	QuerierRoute                       = types.QuerierRoute
	QueryGetBorrows                    = types.QueryGetBorrows
	QueryGetDeposits                   = types.QueryGetDeposits
	QueryGetModuleAccounts             = types.QueryGetModuleAccounts
	QueryGetParams                     = types.QueryGetParams
	QueryGetTotalBorrowed              = types.QueryGetTotalBorrowed
	QueryGetTotalDeposited             = types.QueryGetTotalDeposited
	RouterKey                          = types.RouterKey
	StoreKey                           = types.StoreKey
)

var (
	// function aliases
	APYToSPY                      = keeper.APYToSPY
	CalculateBorrowInterestFactor = keeper.CalculateBorrowInterestFactor
	CalculateBorrowRate           = keeper.CalculateBorrowRate
	CalculateSupplyInterestFactor = keeper.CalculateSupplyInterestFactor
	CalculateUtilizationRatio     = keeper.CalculateUtilizationRatio
	NewKeeper                     = keeper.NewKeeper
	NewQuerier                    = keeper.NewQuerier
	DefaultGenesisState           = types.DefaultGenesisState
	DefaultParams                 = types.DefaultParams
	DepositTypeIteratorKey        = types.DepositTypeIteratorKey
	GetBorrowByLtvKey             = types.GetBorrowByLtvKey
	GetTotalVestingPeriodLength   = types.GetTotalVestingPeriodLength
	NewBorrow                     = types.NewBorrow
	NewBorrowInterestFactor       = types.NewBorrowInterestFactor
	NewBorrowLimit                = types.NewBorrowLimit
	NewDeposit                    = types.NewDeposit
	NewGenesisAccumulationTime    = types.NewGenesisAccumulationTime
	NewGenesisState               = types.NewGenesisState
	NewInterestRateModel          = types.NewInterestRateModel
	NewMoneyMarket                = types.NewMoneyMarket
	NewMsgBorrow                  = types.NewMsgBorrow
	NewMsgDeposit                 = types.NewMsgDeposit
	NewMsgLiquidate               = types.NewMsgLiquidate
	NewMsgRepay                   = types.NewMsgRepay
	NewMsgWithdraw                = types.NewMsgWithdraw
	NewMultiHARDHooks             = types.NewMultiHARDHooks
	NewParams                     = types.NewParams
	NewPeriod                     = types.NewPeriod
	NewQueryAccountParams         = types.NewQueryAccountParams
	NewQueryBorrowsParams         = types.NewQueryBorrowsParams
	NewQueryDepositsParams        = types.NewQueryDepositsParams
	NewQueryTotalBorrowedParams   = types.NewQueryTotalBorrowedParams
	NewQueryTotalDepositedParams  = types.NewQueryTotalDepositedParams
	NewSupplyInterestFactor       = types.NewSupplyInterestFactor
	NewValuationMap               = types.NewValuationMap
	ParamKeyTable                 = types.ParamKeyTable
	RegisterCodec                 = types.RegisterCodec

	// variable aliases
	BorrowInterestFactorPrefix       = types.BorrowInterestFactorPrefix
	BorrowedCoinsPrefix              = types.BorrowedCoinsPrefix
	BorrowsKeyPrefix                 = types.BorrowsKeyPrefix
	DefaultAccumulationTimes         = types.DefaultAccumulationTimes
	DefaultBorrows                   = types.DefaultBorrows
	DefaultCheckLtvIndexCount        = types.DefaultCheckLtvIndexCount
	DefaultDeposits                  = types.DefaultDeposits
	DefaultMoneyMarkets              = types.DefaultMoneyMarkets
	DefaultPreviousBlockTime         = types.DefaultPreviousBlockTime
	DefaultTotalBorrowed             = types.DefaultTotalBorrowed
	DefaultTotalReserves             = types.DefaultTotalReserves
	DefaultTotalSupplied             = types.DefaultTotalSupplied
	DepositsKeyPrefix                = types.DepositsKeyPrefix
	ErrAccountNotFound               = types.ErrAccountNotFound
	ErrBorrowEmptyCoins              = types.ErrBorrowEmptyCoins
	ErrBorrowExceedsAvailableBalance = types.ErrBorrowExceedsAvailableBalance
	ErrBorrowNotFound                = types.ErrBorrowNotFound
	ErrBorrowNotLiquidatable         = types.ErrBorrowNotLiquidatable
	ErrBorrowedCoinsNotFound         = types.ErrBorrowedCoinsNotFound
	ErrDepositNotFound               = types.ErrDepositNotFound
	ErrDepositsNotFound              = types.ErrDepositsNotFound
	ErrGreaterThanAssetBorrowLimit   = types.ErrGreaterThanAssetBorrowLimit
	ErrInsufficientBalanceForBorrow  = types.ErrInsufficientBalanceForBorrow
	ErrInsufficientBalanceForRepay   = types.ErrInsufficientBalanceForRepay
	ErrInsufficientCoins             = types.ErrInsufficientCoins
	ErrInsufficientLoanToValue       = types.ErrInsufficientLoanToValue
	ErrInsufficientModAccountBalance = types.ErrInsufficientModAccountBalance
	ErrInvalidAccountType            = types.ErrInvalidAccountType
	ErrInvalidDepositDenom           = types.ErrInvalidDepositDenom
	ErrInvalidReceiver               = types.ErrInvalidReceiver
	ErrInvalidRepaymentDenom         = types.ErrInvalidRepaymentDenom
	ErrInvalidWithdrawAmount         = types.ErrInvalidWithdrawAmount
	ErrInvalidWithdrawDenom          = types.ErrInvalidWithdrawDenom
	ErrMarketNotFound                = types.ErrMarketNotFound
	ErrMoneyMarketNotFound           = types.ErrMoneyMarketNotFound
	ErrNegativeBorrowedCoins         = types.ErrNegativeBorrowedCoins
	ErrNegativeSuppliedCoins         = types.ErrNegativeSuppliedCoins
	ErrPreviousAccrualTimeNotFound   = types.ErrPreviousAccrualTimeNotFound
	ErrPriceNotFound                 = types.ErrPriceNotFound
	ErrSuppliedCoinsNotFound         = types.ErrSuppliedCoinsNotFound
	GovDenom                         = types.GovDenom
	KeyCheckLtvIndexCount            = types.KeyCheckLtvIndexCount
	KeyMoneyMarkets                  = types.KeyMoneyMarkets
	LtvIndexPrefix                   = types.LtvIndexPrefix
	ModuleCdc                        = types.ModuleCdc
	MoneyMarketsPrefix               = types.MoneyMarketsPrefix
	PreviousAccrualTimePrefix        = types.PreviousAccrualTimePrefix
	PreviousBlockTimeKey             = types.PreviousBlockTimeKey
	SuppliedCoinsPrefix              = types.SuppliedCoinsPrefix
	SupplyInterestFactorPrefix       = types.SupplyInterestFactorPrefix
	TotalReservesPrefix              = types.TotalReservesPrefix
)

type (
	Keeper                    = keeper.Keeper
	LiqData                   = keeper.LiqData
	AccountKeeper             = types.AccountKeeper
	AuctionKeeper             = types.AuctionKeeper
	Borrow                    = types.Borrow
	BorrowInterestFactor      = types.BorrowInterestFactor
	BorrowInterestFactors     = types.BorrowInterestFactors
	BorrowLimit               = types.BorrowLimit
	Borrows                   = types.Borrows
	Deposit                   = types.Deposit
	Deposits                  = types.Deposits
	GenesisAccumulationTime   = types.GenesisAccumulationTime
	GenesisAccumulationTimes  = types.GenesisAccumulationTimes
	GenesisState              = types.GenesisState
	HARDHooks                 = types.HARDHooks
	InterestRateModel         = types.InterestRateModel
	InterestRateModels        = types.InterestRateModels
	MoneyMarket               = types.MoneyMarket
	MoneyMarkets              = types.MoneyMarkets
	MsgBorrow                 = types.MsgBorrow
	MsgDeposit                = types.MsgDeposit
	MsgLiquidate              = types.MsgLiquidate
	MsgRepay                  = types.MsgRepay
	MsgWithdraw               = types.MsgWithdraw
	MultiHARDHooks            = types.MultiHARDHooks
	Params                    = types.Params
	PricefeedKeeper           = types.PricefeedKeeper
	QueryAccountParams        = types.QueryAccountParams
	QueryBorrowsParams        = types.QueryBorrowsParams
	QueryDepositsParams       = types.QueryDepositsParams
	QueryTotalBorrowedParams  = types.QueryTotalBorrowedParams
	QueryTotalDepositedParams = types.QueryTotalDepositedParams
	StakingKeeper             = types.StakingKeeper
	SupplyInterestFactor      = types.SupplyInterestFactor
	SupplyInterestFactors     = types.SupplyInterestFactors
	SupplyKeeper              = types.SupplyKeeper
	ValuationMap              = types.ValuationMap
)
