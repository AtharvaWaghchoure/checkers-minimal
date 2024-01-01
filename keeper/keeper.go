package keeper

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"fmt"
	"github.com/alice/checkers"
	"github.com/cosmos/cosmos-sdk/codec"
)

type Keeper struct {
	coderdecoder codec.BinaryCodec
	// Address convertor from address format to normal format
	addressCodec address.Codec

	// authorizedAddress is the address capable of executing a MsgUdpateParams and other authority-gated message.
	// typically, this should be the x/gov module account
	authorizedAddress string

	Schema      collections.Schema
	Params      collections.Item[checkers.Params]
	StoredGames collections.Map[string, checkers.StoredGame]
}

func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService storetypes.KVStoreService, authorizedAddress string) Keeper {
	if _, err := addressCodec.StringToBytes(authorizedAddress); err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}

	schemeBuilder := collections.NewSchemaBuilder(storeService) // we created the schema for the keeper storing the keys i think
	keeper := Keeper{
		// giving the encoding and decoding type
		coderdecoder: cdc,
		// address conversion from the address format to normal format and vise versa
		addressCodec: addressCodec,
		// Address capable of performing state changing task
		authorizedAddress: authorizedAddress,
		Params:            collections.NewItem(schemeBuilder, checkers.ParamsKey, "params", codec.CollValue[checkers.Params](cdc)),
		StoredGames:       collections.NewMap(schemeBuilder, checkers.StoredGamesKey, "storedGames", collections.StringKey, codec.CollValue[checkers.StoredGame](cdc)),
	}

	schema, err := schemeBuilder.Build()
	if err != nil {
		panic(err)
	}

	keeper.Schema = schema

	return keeper
}

// GetAuthority returns the module's authority.
func (keeper Keeper) GetAuthority() string {
	return keeper.authorizedAddress
}
