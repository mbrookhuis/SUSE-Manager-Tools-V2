package consts

import (
	_formUsed "ecp-golang-cm/pkg/models/susemanager"
	_utilhn "ecp-golang-cm/pkg/util/hostname"
)

const DefaultAutoyastDir string = "/opt/autoyast"
const AutoyastTypes string = "SL_SERVER-s154 SL_SERVER-s155 K3S_MGMT-mi52 K3S_MGMT-mi55 K3S_SERVER-mi52 K3S_SERVER-mi55 POD_SERVER-mi52 POD_SERVER-mi55 SUMAS_SERVER-sm43 dtag_server"
const VirtType string = "none"
const AutoyastDistribution string = "installFirstRun"

var dtagServerVars = map[string]interface{}{
	"org":          1,
	"SUMAN_SERVER": _utilhn.GetHostnameFqdn(),
}

var ProfileVariables = []_formUsed.KickstartProfileVar{
	{
		ProfileName: "dtag_server",
		ProfileVars: dtagServerVars,
	},
}
