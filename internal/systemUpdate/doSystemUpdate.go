package systemUpdate

import (
	"fmt"

	_sumanUseCase "SUSE-Manager-Tools-V2/internal/susemanager"
	returnCodes "SUSE-Manager-Tools-V2/internal/util/returnCodes"
	util "SUSE-Manager-Tools-V2/internal/util/uuid"
	"SUSE-Manager-Tools-V2/internal/vars"
	"go.uber.org/zap"
	zapCore "go.uber.org/zap/zapcore"
)

type SystemUpdate struct {
	sumanProxy           _sumanUseCase.IProxy
	suse                 _sumanUseCase.ISuseManager
	suseOperationTimeout int
	logger               *zap.Logger
	params               vars.Params
	configFile           vars.ConfigInfo
	requestID            string
}

// NewSystemUpdate Perform system update
//
// param: sumanProxy
// param: suse
// param: suseOperationTimeout
// param: logger
// param: params
// param: configFile
// return:
func NewSystemUpdate(sumanProxy _sumanUseCase.IProxy, suse _sumanUseCase.ISuseManager, suseOperationTimeout int, logger *zap.Logger, params vars.Params, configFile vars.ConfigInfo, requestID string) ISystemUpdate {
	return &SystemUpdate{
		sumanProxy:           sumanProxy,
		suse:                 suse,
		suseOperationTimeout: suseOperationTimeout,
		logger:               logger,
		params:               params,
		configFile:           configFile,
		requestID:            requestID,
	}
}

// SystemUpdate Perform system update
func (h *SystemUpdate) SystemUpdate() error {
	zf := []zapCore.Field{
		zap.Any("requestID", h.requestID),
		zap.Any("hostname", h.params.Server),
	}
	h.logger.Info("starting update", zf...)
	h.logger.Debug("parameters given",
		zap.Any("requestID", h.requestID),
		zap.Any("hostname", h.params.Server),
		zap.Any("noDryRun", h.params.NoDryRun),
		zap.Any("noReboot", h.params.NoReboot),
		zap.Any("forceReboot", h.params.ForceReboot),
		zap.Any("applyConfig", h.params.ApplyConfig),
		zap.Any("updateScripts", h.params.UpdateScript),
		zap.Any("postScript", h.params.PostScript),
		zap.Any("configFile", h.params.ConfigFile))
	requestID := util.GenerateUniqueID()
	sessionKey, err := h.sumanProxy.SumanLogin(requestID)
	if err != nil {
		h.logger.Error(returnCodes.ErrLoginSuseManager, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return err
	}
	// Fetch auth for further use.
	auth, err := h.suse.GetAuth(sessionKey)
	if err != nil {
		return err
	}
	defer func(sumanProxy _sumanUseCase.IProxy, requestID string, auth _sumanUseCase.AuthParams) {
		err := sumanProxy.SumanLogout(requestID, auth)
		if err != nil {
			h.logger.Error(returnCodes.ErrLogoutSuseManager, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		}
	}(h.sumanProxy, requestID, *auth)
	systemId, err := h.checkSystemExists(h.params.Server, auth, zf...)
	if err != nil {
		return err
	}
	err = h.updateServer(systemId, auth, zf...)
	return nil
}

func (h *SystemUpdate) checkSystemExists(systemName string, auth *_sumanUseCase.AuthParams, zf ...zapCore.Field) (int, error) {
	h.logger.Debug("Function checkSystemExists started", zf...)
	system, err := h.sumanProxy.SystemGetID(h.requestID, *auth, systemName)
	if err != nil {
		h.logger.Debug(fmt.Sprintf("%v, error %v", returnCodes.ErrSystemNotFound, err.Error()), zf...)
		return 0, err
	}
	h.logger.Debug("Function checkSystemExists finished", zf...)
	return system[0].ID, nil
}
