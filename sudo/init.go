package sudo

import eos "github.com/eosforce/goeosforce"

func init() {
	eos.RegisterAction(AN("eosio.sudo"), ActN("exec"), Exec{})
}

var AN = eos.AN
var ActN = eos.ActN
