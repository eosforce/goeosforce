package system

import "github.com/eosforce/goeosforce"

// NewNonce returns a `nonce` action that lives on the
// `eosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `eosio.system` contract.
func NewVoteProducer(voter eos.AccountName, bpName eos.AccountName, amount eos.Asset) *eos.Action {
	a := &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("vote"),
		Authorization: []eos.PermissionLevel{
			{Actor: voter, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(
			VoteProducer{
				Voter:     voter,
				BpName:     bpName,
				Amount: amount,
			},
		),
	}
	return a
}

// VoteProducer represents the `eosio.system::voteproducer` action
type VoteProducer struct {
	Voter     eos.AccountName   `json:"voter"`
	BpName    eos.AccountName   `json:"proxy"`
	Amount 	  eos.Asset `json:"amount"`
}
