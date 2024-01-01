package keeper

import (
	"context"

	"github.com/alice/checkers"
)

// InitGenesis initializes the module state from a genesis state.
func (keeper *Keeper) InitGenesis(ctx context.Context, data *checkers.GenesisState) error {
	if err := keeper.Params.Set(ctx, data.Params); err != nil {
		return err
	}

	for _, indexedStoredGame := range data.IndexedStoredGameList {
		if err := keeper.StoredGames.Set(ctx, indexedStoredGame.Index, indexedStoredGame.StoredGame); err != nil {
			return err
		}
	}
	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (keeper *Keeper) ExportGenesis(ctx context.Context) (*checkers.GenesisState, error) {
	params, err := keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	var indexedStoredGames []checkers.IndexedStoredGame
	if err := keeper.StoredGames.Walk(ctx, nil, func(index string, storedGame checkers.StoredGame) (bool, error) {
		// we add the current game to the indexed game list
		indexedStoredGames = append(indexedStoredGames, checkers.IndexedStoredGame{
			Index:      index,
			StoredGame: storedGame,
		})
		return false, nil
	}); err != nil {
		return nil, err
	}

	return &checkers.GenesisState{
		Params:                params,
		IndexedStoredGameList: indexedStoredGames,
	}, nil
}
