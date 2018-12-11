package system

import (
	eos "github.com/eosforce/goeosforce"
//	"github.com/eosforce/goeosforce/ecc"
)

// NewRegProducer returns a `regproducer` action that lives on the
// `eosio.system` contract.
func NewSetemergency(producer eos.AccountName,Emergency bool) *eos.Action {
	return &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("setemergency"),
		Authorization: []eos.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Setemergency{
			Producer:    producer,
			Emergency: Emergency,
		}),
	}
}

// RegProducer represents the `eosio.system::regproducer` action
type Setemergency struct {
	Producer    eos.AccountName `json:"producer"`
	Emergency	bool	`json:"emergency"`
}
