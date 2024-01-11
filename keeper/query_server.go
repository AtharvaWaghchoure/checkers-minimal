package keeper

import (
	"context"
	"cosmossdk.io/collections"
	"errors"
	"github.com/alice/checkers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ checkers.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(keeper Keeper) checkers.QueryServer {
	return queryServer{keeper}
}

type queryServer struct {
	keeper Keeper
}

// GetGame defines the handler for the Query/GetGame RPC method.
func (qs queryServer) GetGame(ctx context.Context, req *checkers.QueryGetGameRequest) (*checkers.QueryGetGameResponse, error) {
	game, err := qs.keeper.StoredGames.Get(ctx, req.Index)

	if err == nil {
		return &checkers.QueryGetGameResponse{Game: &game}, nil
	}
	if errors.Is(err, collections.ErrNotFound) {
		return &checkers.QueryGetGameResponse{Game: nil}, nil
	}
	return nil, status.Error(codes.Internal, err.Error())
}
