package systemUpdate

import (
	"fmt"

	"SUSE-Manager-Tools-V2/internal/vars"
	"go.uber.org/zap"
)

// validateParams
//
// param: params
// return: error or nil
func validateParams(params vars.Params, logger *zap.Logger) error {
	// Check if there is a server given that needs to be updated
	if params.Server == "not_set" {
		logger.Error("Server to be updated is not given", zap.Any("given server", params.Server))
		return fmt.Errorf("server to be updated is not given")
	}
	return nil
}
