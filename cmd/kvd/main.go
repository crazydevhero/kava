// Copyright 2016 All in Bits, inc
// Modifications copyright 2019 Kava Labs

package main

import (
	"encoding/json"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/genaccounts"
	genaccscli "github.com/cosmos/cosmos-sdk/x/auth/genaccounts/client/cli"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/kava-labs/_/app"
)

// gaiad custom flags
const flagInvCheckPeriod = "inv-check-period"

var invCheckPeriod uint

func main() {
	cdc := app.MakeCodec()

	config := sdk.GetConfig()
	app.SetAddressPrefixes(config)
	config.Seal()

	ctx := server.NewDefaultContext()
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:               "kvd",
		Short:             "Kava Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(genutilcli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.CollectGenTxsCmd(ctx, cdc, genaccounts.AppModuleBasic{}, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.GenTxCmd(ctx, cdc, app.ModuleBasics, staking.AppModuleBasic{},
		genaccounts.AppModuleBasic{}, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(genutilcli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics))
	rootCmd.AddCommand(genaccscli.AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(client.NewCompletionCmd(rootCmd, true))
	rootCmd.AddCommand(testnetCmd(ctx, cdc, app.ModuleBasics, genaccounts.AppModuleBasic{}))
	rootCmd.AddCommand(replayCmd())

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "KA", app.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewApp(
		logger, db, traceStore, true, invCheckPeriod,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
		baseapp.SetHaltHeight(uint64(viper.GetInt(server.FlagHaltHeight))),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {

	if height != -1 {
		appStruct := app.NewApp(logger, db, traceStore, false, uint(1))
		err := appStruct.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return appStruct.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}
	appStruct := app.NewApp(logger, db, traceStore, true, uint(1))
	return appStruct.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
