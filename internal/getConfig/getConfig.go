package getConfig

import (
	"os"
)

// ReadYamlConfigFile
//
// param: filename
// return: ConfigInfo or error
func ReadYamlConfigFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return data, err
	}
	return data, nil
}
