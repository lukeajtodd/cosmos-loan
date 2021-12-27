package keeper

import (
	"github.com/lukeajtodd/loan/x/loan/types"
)

var _ types.QueryServer = Keeper{}
