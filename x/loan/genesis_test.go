package loan_test

import (
	"testing"

	keepertest "github.com/jensoncollins/loan/testutil/keeper"
	"github.com/jensoncollins/loan/testutil/nullify"
	"github.com/jensoncollins/loan/x/loan"
	"github.com/jensoncollins/loan/x/loan/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LoanList: []types.Loan{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		LoanCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LoanKeeper(t)
	loan.InitGenesis(ctx, *k, genesisState)
	got := loan.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LoanList, got.LoanList)
	require.Equal(t, genesisState.LoanCount, got.LoanCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
