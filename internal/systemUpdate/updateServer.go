package systemUpdate

import (
	"fmt"
	"path/filepath"

	gc "SUSE-Manager-Tools-V2/internal/getConfig"
	_sumanUseCase "SUSE-Manager-Tools-V2/internal/susemanager"
	"SUSE-Manager-Tools-V2/internal/util/contains"
	"SUSE-Manager-Tools-V2/internal/vars"
	"go.uber.org/zap"
	zapCore "go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
)

func (h *SystemUpdate) updateServer(systemId int, auth *_sumanUseCase.AuthParams, zf ...zapCore.Field) error {
	h.logger.Debug("Function updateServer started", zf...)
	// check if server is on the excluded list
	for _, value := range h.configFile.Maintenance.ExcludeForPatch {
		if contains.PartOff(h.params.Server, value) {
			h.logger.Error("Server to update is excluded from patching", zf...)
			return fmt.Errorf("server %s to update is excluded from patching", h.params.Server)
		}
	}
	// check if server is inactive and if this is true, stop processing.
	err := h.isSystemInactive(auth, zf...)
	if err != nil {
		return err
	}
	if h.params.UpdateScript {
		err = h.doUpdateActions("begin", "general", systemId, auth, zf...)
		if err != nil {
			return err
		}
		err = h.doUpdateActions("begin", h.params.Server, systemId, auth, zf...)
		if err != nil {
			return err
		}
	}
	//
	//
	//
	if h.params.UpdateScript {
		err = h.doUpdateActions("end", "general", systemId, auth, zf...)
		if err != nil {
			return err
		}
		err = h.doUpdateActions("end", h.params.Server, systemId, auth, zf...)
		if err != nil {
			return err
		}

	}

	h.logger.Debug("Function updateServer finished", zf...)
	return nil
}

func (h *SystemUpdate) isSystemInactive(auth *_sumanUseCase.AuthParams, zf ...zapCore.Field) error {
	h.logger.Debug("Function isSystemInactive started", zf...)
	inActiveSystems, err := h.sumanProxy.SystemListInActiveSystems(h.requestID, *auth)
	if err != nil {
		zf = append(zf, zap.Any("error", err))
		h.logger.Error("Unable to get lust of inactiveSystems", zf...)
		return err
	}
	for _, system := range inActiveSystems {
		if contains.PartOff(h.params.Server, system.Name) {
			h.logger.Error("Server is inactive", zf...)
			return fmt.Errorf("server %s is inactive", h.params.Server)
		}
	}
	h.logger.Debug("Function isSystemInactive finished", zf...)
	return nil
}

func (h *SystemUpdate) doUpdateActions(phase string, eventFile string, systemId int, auth *_sumanUseCase.AuthParams, zf ...zapCore.Field) error {
	h.logger.Debug("Function performUpdateScript started", zf...)
	// first do general, then server based
	// first check if there is a entry for script
	// then execute state.

	// get general file
	data, err := gc.ReadYamlConfigFile(filepath.Join(h.configFile.Dirs.UpdateScriptDir, eventFile))
	if err != nil {
		h.logger.Warn("Unable to read updateScript, skipping", zap.Any("phase", phase), zap.Any("Type", eventFile), zap.Any("Error", err))
		return nil
	}
	var scriptData vars.UpdateScript
	err = yaml.Unmarshal(data, &scriptData)
	if err != nil {
		h.logger.Error("Unable to convert updateScript", zap.Any("Error", err))
		return fmt.Errorf("unable to convert updatescript %s %s", phase, eventFile)
	}
	var stateToRun, scriptToRun []string
	switch phase {
	case "begin":
		scriptToRun = scriptData.BeginScript.Commands
		if len(scriptData.BeginScript.State) > 0 {
			err = h.executeStates(stateToRun, systemId, auth, zf...)
			if err != nil {
				return err
			}
		} else {
			h.logger.Info(fmt.Sprintf("There are no states to be executed for phase %s", phase), zf...)
		}
	case "end":
		scriptToRun = scriptData.EndScript.Commands
		if len(scriptData.EndScript.State) > 0 {
			err = h.executeStates(stateToRun, systemId, auth, zf...)
			if err != nil {
				return err
			}
		} else {
			h.logger.Info(fmt.Sprintf("There are no states to be executed for phase %s", phase), zf...)
		}
	}

	fmt.Println(scriptToRun)

	if len(scriptToRun) > 0 {
		fmt.Println("er zijn scripts")
	} else {
		h.logger.Info(fmt.Sprintf("There are no scripts to be executed for phase %s", phase), zf...)
	}

	h.logger.Debug("Function performUpdateScript finished", zf...)
	return nil
}

// executeStates
//
// param: states
// param: systemId
// param: auth
// param: zf
func (h SystemUpdate) executeStates(states []string, systemId int, auth *_sumanUseCase.AuthParams, zf ...zapCore.Field) error {
	h.logger.Debug("Function executeStates started", zf...)
	err := h.sumanProxy.SystemScheduleApplyStates(h.requestID, *auth, systemId, states, 600)
	if err != nil {
		zf = append(zf, zap.Any("Error", err))
		h.logger.Error("Highstate failed", zf...)
		return err
	}
	h.logger.Debug("Function executeStates started", zf...)
	return nil
}

// executeScripts
//
// param: scripts
// param: systemId
// param: auth
// param: zf
func (h SystemUpdate) executeScripts(scripts []string, systemId int, auth *_sumanUseCase.AuthParams, zf ...zapCore.Field) error {
	h.logger.Debug("Function executeScripts started", zf...)
	for _, script := range scripts {
		err := h.sumanProxy.ScheduleScriptRun(h.requestID, *auth, systemId, h.configFile.Suman.Timeout, script)
		if err != nil {
			zf = append(zf, zap.Any("Error", err))
			h.logger.Error("Highstate failed", zf...)
			return err
		}
	}
	h.logger.Debug("Function executeScripts started", zf...)
	return nil
}
