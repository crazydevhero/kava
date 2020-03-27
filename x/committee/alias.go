// nolint
// DO NOT EDIT - generated by aliasgen tool (github.com/rhuairahrighairidh/aliasgen)
package committee

import (
	"github.com/kava-labs/kava/x/committee/client"
	"github.com/kava-labs/kava/x/committee/keeper"
	"github.com/kava-labs/kava/x/committee/types"
)

const (
	AttributeKeyProposalID      = types.AttributeKeyProposalID
	DefaultCodespace            = types.DefaultCodespace
	DefaultNextProposalID       = types.DefaultNextProposalID
	DefaultParamspace           = types.DefaultParamspace
	EventTypeSubmitProposal     = types.EventTypeSubmitProposal
	ModuleName                  = types.ModuleName
	ProposalTypeCommitteeChange = types.ProposalTypeCommitteeChange
	ProposalTypeCommitteeDelete = types.ProposalTypeCommitteeDelete
	QuerierRoute                = types.QuerierRoute
	QueryCommittee              = types.QueryCommittee
	QueryCommittees             = types.QueryCommittees
	QueryProposal               = types.QueryProposal
	QueryProposals              = types.QueryProposals
	QueryTally                  = types.QueryTally
	QueryVote                   = types.QueryVote
	QueryVotes                  = types.QueryVotes
	RouterKey                   = types.RouterKey
	StoreKey                    = types.StoreKey
	TypeMsgSubmitProposal       = types.TypeMsgSubmitProposal
	TypeMsgVote                 = types.TypeMsgVote
)

var (
	// function aliases
	NewKeeper                  = keeper.NewKeeper
	NewQuerier                 = keeper.NewQuerier
	DefaultGenesisState        = types.DefaultGenesisState
	GetKeyFromID               = types.GetKeyFromID
	GetVoteKey                 = types.GetVoteKey
	NewCommittee               = types.NewCommittee
	NewCommitteeChangeProposal = types.NewCommitteeChangeProposal
	NewCommitteeDeleteProposal = types.NewCommitteeDeleteProposal
	NewGenesisState            = types.NewGenesisState
	NewMsgSubmitProposal       = types.NewMsgSubmitProposal
	NewMsgVote                 = types.NewMsgVote
	NewQueryCommitteeParams    = types.NewQueryCommitteeParams
	NewQueryProposalParams     = types.NewQueryProposalParams
	NewQueryVoteParams         = types.NewQueryVoteParams
	RegisterAppCodec           = types.RegisterAppCodec
	RegisterModuleCodec        = types.RegisterModuleCodec
	RegisterProposalTypeCodec  = types.RegisterProposalTypeCodec
	Uint64FromBytes            = types.Uint64FromBytes

	// variable aliases
	ProposalHandler     = client.ProposalHandler
	CommitteeKeyPrefix  = types.CommitteeKeyPrefix
	MaxProposalDuration = types.MaxProposalDuration
	ModuleCdc           = types.ModuleCdc
	NextProposalIDKey   = types.NextProposalIDKey
	ProposalKeyPrefix   = types.ProposalKeyPrefix
	VoteKeyPrefix       = types.VoteKeyPrefix
	VoteThreshold       = types.VoteThreshold
)

type (
	Keeper                  = keeper.Keeper
	AllowedParam            = types.AllowedParam
	AllowedParams           = types.AllowedParams
	Committee               = types.Committee
	CommitteeChangeProposal = types.CommitteeChangeProposal
	CommitteeDeleteProposal = types.CommitteeDeleteProposal
	GenesisState            = types.GenesisState
	GodPermission           = types.GodPermission
	MsgSubmitProposal       = types.MsgSubmitProposal
	MsgVote                 = types.MsgVote
	ParamChangePermission   = types.ParamChangePermission
	Permission              = types.Permission
	Proposal                = types.Proposal
	PubProposal             = types.PubProposal
	QueryCommitteeParams    = types.QueryCommitteeParams
	QueryProposalParams     = types.QueryProposalParams
	QueryVoteParams         = types.QueryVoteParams
	ShutdownPermission      = types.ShutdownPermission
	Vote                    = types.Vote
)
