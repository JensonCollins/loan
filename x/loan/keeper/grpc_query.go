package keeper

import (
	"github.com/jensoncollins/loan/x/loan/types"
)

var _ types.QueryServer = Keeper{}
