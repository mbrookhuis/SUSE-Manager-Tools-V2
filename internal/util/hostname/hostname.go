package hostname

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetHostnameFqdn() string {
	cmd := exec.Command("/bin/hostname", "-f")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Sprintf("Error retrieving get_hostname_fqdn: %v", err)
		return ""
	}
	fqdn := out.String()
	fqdn = fqdn[:len(fqdn)-1] // removing EOL
	return fqdn
}
