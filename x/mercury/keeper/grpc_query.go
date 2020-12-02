package keeper

import (
	"github.com/fadeev/mercury/x/mercury/types"
)

var _ types.QueryServer = Keeper{}
