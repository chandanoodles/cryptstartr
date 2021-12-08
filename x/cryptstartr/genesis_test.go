package cryptstartr_test

import (
	"testing"

	keepertest "github.com/daniel-farina/cryptstartr/testutil/keeper"
	"github.com/daniel-farina/cryptstartr/x/cryptstartr"
	"github.com/daniel-farina/cryptstartr/x/cryptstartr/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CryptstartrKeeper(t)
	cryptstartr.InitGenesis(ctx, *k, genesisState)
	got := cryptstartr.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
