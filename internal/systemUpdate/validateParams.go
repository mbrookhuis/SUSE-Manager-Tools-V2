package systemUpdate

import (
	"fmt"
	"os"

	"SUSE-Manager-Tools-V2/internal/vars"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// validateParams
//
// param: params
// return: error or nil
func validateParams(params vars.Params, logger *zap.Logger) error {
	// Check if the given config file exists.
	if _, err := os.Stat(params.ConfigFile); errors.Is(err, os.ErrNotExist) {
		logger.Error("the given configuration file is not present.", zap.Any("configfile", params.ConfigFile))
		return fmt.Errorf("configfile %s does not exist", params.ConfigFile)
	}
	// Check if there is a server given that needs to be updated
	if params.Server == "not_set" {
		logger.Error("Server to be updated is not given", zap.Any("given server", params.Server))
		return fmt.Errorf("server to be updated is not given")
	}
	return nil
}
