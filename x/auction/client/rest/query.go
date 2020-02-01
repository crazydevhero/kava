package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/kava-labs/kava/x/auction/types"
)

const RestAuctionID = "auction-id"

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/auction/auctions/{%s}", RestAuctionID), queryAuctionHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/auction/auctions", queryAuctionsHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/auction/parameters", getParamsHandlerFn(cliCtx)).Methods("GET")
}

func queryAuctionHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the query height
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		// Prepare params for querier
		vars := mux.Vars(r)
		if len(vars[RestAuctionID]) == 0 {
			err := fmt.Errorf("%s required but not specified", RestAuctionID)
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		auctionID, ok := rest.ParseUint64OrReturnBadRequest(w, vars[RestAuctionID])
		if !ok {
			return
		}
		bz, err := cliCtx.Codec.MarshalJSON(types.QueryAuctionParams{AuctionID: auctionID})
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Query
		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("/custom/%s/%s", types.ModuleName, types.QueryGetAuction), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		var auction types.Auction
		err = cliCtx.Codec.UnmarshalJSON(res, &auction)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		// Decode and return results
		cliCtx = cliCtx.WithHeight(height)

		auctionWithPhase := types.NewAuctionWithPhase(auction)

		rest.PostProcessResponse(w, cliCtx, cliCtx.Codec.MustMarshalJSON(auctionWithPhase))
	}
}

func queryAuctionsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the query height
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}
		// Get all auctions
		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("/custom/%s/%s", types.ModuleName, types.QueryGetAuctions), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		// Return auctions
		cliCtx = cliCtx.WithHeight(height)

		var auctions types.Auctions
		err = cliCtx.Codec.UnmarshalJSON(res, &auctions)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		var auctionsWithPhase []types.AuctionWithPhase
		for _, a := range auctions {
			auctionsWithPhase = append(auctionsWithPhase, types.NewAuctionWithPhase(a))
		}

		rest.PostProcessResponse(w, cliCtx, cliCtx.Codec.MustMarshalJSON(auctionsWithPhase))
	}
}

func getParamsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the query height
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}
		// Get the params
		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryGetParams), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		// Return the params
		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
