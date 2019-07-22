package rex

import (
	eos "github.com/eosforce/goeosforce"
)

func NewREXExec(user eos.AccountName, max uint16) *eos.Action {
	return &eos.Action{
		Account: REXAN,
		Name:    ActN("rexexec"),
		Authorization: []eos.PermissionLevel{
			{Actor: user, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(REXExec{
			User: user,
			Max:  max,
		}),
	}
}

type REXExec struct {
	User eos.AccountName
	Max  uint16
}
