package keeper

import (
	"github.com/daniel-farina/cryptstartr/x/cryptstartr/types"
)

var _ types.QueryServer = Keeper{}
