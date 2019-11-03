package rex

import (
	eos "github.com/eosforce/goeosforce"
)

func NewUpdateREX(owner eos.AccountName) *eos.Action {
	return &eos.Action{
		Account: REXAN,
		Name:    ActN("updaterex"),
		Authorization: []eos.PermissionLevel{
			{Actor: owner, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(UpdateREX{
			Owner: owner,
		}),
	}
}

type UpdateREX struct {
	Owner eos.AccountName
}
