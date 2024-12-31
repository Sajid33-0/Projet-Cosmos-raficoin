package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "RafiCoinNetwork/testutil/keeper"
	"RafiCoinNetwork/x/raficoinnetwork/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.RaficoinnetworkKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
