package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/daniel-farina/cryptstartr/testutil/keeper"
	"github.com/daniel-farina/cryptstartr/x/cryptstartr/keeper"
	"github.com/daniel-farina/cryptstartr/x/cryptstartr/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CryptstartrKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
