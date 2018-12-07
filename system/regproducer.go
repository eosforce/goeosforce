package system

import (
	eos "github.com/eosforce/goeosforce"
	"github.com/eosforce/goeosforce/ecc"
)

// NewRegProducer returns a `regproducer` action that lives on the
// `eosio.system` contract.
func NewRegProducer(producer eos.AccountName, producerKey ecc.PublicKey, url string, commissionRate int32) *eos.Action {
	return &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("updatebp"),
		Authorization: []eos.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(RegProducer{
			Producer:    producer,
			ProducerKey: producerKey,
			CommissionRate: commissionRate,
			URL:         url,
		}),
	}
}

// RegProducer represents the `eosio.system::regproducer` action
type RegProducer struct {
	Producer    eos.AccountName `json:"producer"`
	ProducerKey ecc.PublicKey   `json:"producer_key"`
	CommissionRate int32		`json:"commission_rate"`
	URL         string          `json:"url"`
}
