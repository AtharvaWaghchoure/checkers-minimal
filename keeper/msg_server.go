package keeper

import (
	"context"
	"cosmossdk.io/collections"
	"errors"
	"fmt"
	"github.com/alice/checkers"
	"github.com/alice/checkers/rules"
)

type msgServer struct {
	keeper Keeper
}

var _ checkers.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) checkers.MsgServer {
	return &msgServer{keeper}
	// so the MsgServer is an interface through which we can use the CreateGame
	// rpc method
}

// CreateGame defines the handler for the MsgCreateGame message.
func (ms msgServer) CreateGame(ctx context.Context, msg *checkers.MsgCreateGame) (*checkers.MsgCreateGameResponse, error) {
	if length := len([]byte(msg.Index)); checkers.MaxIndexLength < length || length < 1 {
		return nil, checkers.ErrIndexTooLong
	}
	if _, err := ms.keeper.StoredGames.Get(ctx, msg.Index); err == nil || errors.Is(err, collections.ErrEncoding) {
		return nil, fmt.Errorf("game already exists at index: %s", msg.Index)
	}

	newBoard := rules.New()
	storedGame := checkers.StoredGame{
		Board: newBoard.String(),
		Turn:  rules.PieceStrings[newBoard.Turn],
		Black: msg.Black,
		Red:   msg.Red,
	}
	if err := storedGame.Validate(); err != nil {
		return nil, err
	}
	if err := ms.keeper.StoredGames.Set(ctx, msg.Index, storedGame); err != nil {
		return nil, err
	}

	return &checkers.MsgCreateGameResponse{}, nil

}
