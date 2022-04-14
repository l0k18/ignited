package ignited_test

import (
	"testing"

	keepertest "github.com/l0k18/ignited/testutil/keeper"
	"github.com/l0k18/ignited/testutil/nullify"
	"github.com/l0k18/ignited/x/ignited"
	"github.com/l0k18/ignited/x/ignited/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IgnitedKeeper(t)
	ignited.InitGenesis(ctx, *k, genesisState)
	got := ignited.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
