package rex

import (
	eos "github.com/eosforce/goeosforce"
)

func NewBuyREX(from eos.AccountName, amount eos.Asset) *eos.Action {
	return &eos.Action{
		Account: REXAN,
		Name:    ActN("buyrex"),
		Authorization: []eos.PermissionLevel{
			{Actor: from, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(BuyREX{
			From:   from,
			Amount: amount,
		}),
	}
}

type BuyREX struct {
	From   eos.AccountName
	Amount eos.Asset
}
