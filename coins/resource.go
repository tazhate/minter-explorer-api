package coins

import (
	"github.com/MinterTeam/minter-explorer-api/helpers"
	"github.com/MinterTeam/minter-explorer-api/resource"
	"github.com/MinterTeam/minter-explorer-tools/models"
)

type Resource struct {
	Crr            uint64 `json:"crr"`
	Volume         string `json:"volume"`
	ReserveBalance string `json:"reserveBalance"`
	MaxSupply      string `json:"max_supply"`
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
}

func (Resource) Transform(model resource.ItemInterface, params ...resource.ParamInterface) resource.Interface {
	coin := model.(models.Coin)
	return Resource{
		Crr:            coin.Crr,
		Volume:         helpers.PipStr2Bip(coin.Volume),
		ReserveBalance: helpers.PipStr2Bip(coin.ReserveBalance),
		MaxSupply:      helpers.PipStr2Bip(coin.MaxSupply),
		Name:           coin.Name,
		Symbol:         coin.Symbol,
	}
}
