package msig

import eos "github.com/eosforce/goeosforce"

type ProposalRow struct {
	ProposalName       eos.Name              `json:"proposal_name"`
	RequestedApprovals []eos.PermissionLevel `json:"requested_approvals"`
	ProvidedApprovals  []eos.PermissionLevel `json:"provided_approvals"`
	PackedTransaction  eos.HexBytes          `json:"packed_transaction"`
}
