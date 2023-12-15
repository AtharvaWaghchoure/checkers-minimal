package module

import (
	autocli "cosmossdk.io/api/cosmos/autocli/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocli.ModuleOptions {
	return &autocli.ModuleOptions{
		Query: nil,
		Tx:    nil,
	}
}
