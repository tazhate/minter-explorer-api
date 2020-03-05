package data_resources

import (
	"github.com/MinterTeam/minter-explorer-api/helpers"
	"github.com/MinterTeam/minter-explorer-api/resource"
	"github.com/MinterTeam/minter-explorer-tools/v4/models"
	"github.com/MinterTeam/minter-go-sdk/transaction"
)

type RedeemCheck struct {
	RawCheck string    `json:"raw_check"`
	Proof    string    `json:"proof"`
	Check    CheckData `json:"check"`
}

type CheckData struct {
	Coin     string `json:"coin"`
	GasCoin  string `json:"gas_coin"`
	Nonce    string `json:"nonce"`
	Value    string `json:"value"`
	Sender   string `json:"sender"`
	DueBlock uint64 `json:"due_block"`
}

func (RedeemCheck) Transform(txData resource.ItemInterface, params ...resource.ParamInterface) resource.Interface {
	data := txData.(*models.RedeemCheckTxData)

	//TODO: handle error
	check, _ := TransformCheckData(data.RawCheck)

	return RedeemCheck{
		RawCheck: data.RawCheck,
		Proof:    data.Proof,
		Check:    check,
	}
}

func TransformCheckData(raw string) (CheckData, error) {
	data, err := transaction.DecodeCheck(raw)
	if err != nil {
		return CheckData{}, err
	}

	sender, err := data.Sender()
	if err != nil {
		return CheckData{}, err
	}

	return CheckData{
		Coin:     string(data.Coin[:]),
		GasCoin:  string(data.GasCoin[:]),
		Nonce:    string(data.Nonce[:]),
		Value:    helpers.PipStr2Bip(data.Value.String()),
		Sender:   sender,
		DueBlock: data.DueBlock,
	}, nil
}

