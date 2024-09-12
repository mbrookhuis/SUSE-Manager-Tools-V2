package systemUpdate

import (
	"fmt"
	"path/filepath"
	"strings"

	"SUSE-Manager-Tools-V2/internal/config"
	gc "SUSE-Manager-Tools-V2/internal/getConfig"
	_sumanUseCase "SUSE-Manager-Tools-V2/internal/susemanager"
	"SUSE-Manager-Tools-V2/internal/util/contains"
	"SUSE-Manager-Tools-V2/internal/vars"
	"go.uber.org/zap"
	zapCore "go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
)

// updateServer
//
// param: systemId
// param: auth
// param: zf
func (h *SystemUpdate) updateServer(systemId int, auth *_sumanUseCase.AuthParams, zf ...zapCore.Field) error {
	h.logger.Debug("Function updateServer started", zf...)
	// check if server is on the excluded list
	for _, value := range config.GetConfig().Maintenance.ExcludeForPatch {
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
	doSPM, newBaseChannel := h.checkForSpMigration(systemId, auth, zf...)
	h.logger.Debug(fmt.Sprintf("MBMBMB: do spm %t, new basechannel: %s", doSPM, newBaseChannel))

	//
	//     (do_spm, new_basechannel) = check_for_sp_migration()
	//     if do_spm:
	//        smt.log_info("Server {} will get a SupportPack Migration to {} ".format(args.server, new_basechannel))
	//        do_spmigrate(new_basechannel, args.noreboot, args.nodryrun)
	//    else:
	//        smt.log_info("Server {} will be upgraded with latest available patches".format(args.server))
	//        do_upgrade(args.noreboot, args.forcereboot)
	//    highstate_done = False
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

// isSystemInactive
//
// param: auth
// param: zf
func (h *SystemUpdate) isSystemInactive(auth *_sumanUseCase.AuthParams, zf ...zapCore.Field) error {
	h.logger.Debug("Function isSystemInactive started", zf...)
	inActiveSystems, err := h.sumanProxy.SystemListInActiveSystems(*auth)
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

// doUpdateActions
//
// param: phase
// param: eventFile
// param: systemId
// param: auth
// param: zf
func (h *SystemUpdate) doUpdateActions(phase string, eventFile string, systemId int, auth *_sumanUseCase.AuthParams, zf ...zapCore.Field) error {
	h.logger.Debug("Function doUpdateActions started", zf...)
	// first do general, then server based
	// first check if there is an entry for script
	// then execute state.

	// get general file
	data, err := gc.ReadYamlConfigFile(filepath.Join(config.GetConfig().Dirs.UpdateScriptDir, eventFile))
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
			for _, state := range scriptData.BeginScript.State {
				stateToRun = append(stateToRun, state)
			}
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
			for _, state := range scriptData.EndScript.State {
				stateToRun = append(stateToRun, state)
			}
			err = h.executeStates(stateToRun, systemId, auth, zf...)
			if err != nil {
				return err
			}
		} else {
			h.logger.Info(fmt.Sprintf("There are no states to be executed for phase %s", phase), zf...)
		}
	}
	if len(scriptToRun) > 0 {
		var commands string
		for _, command := range scriptToRun {
			commands = commands + "\n" + command
		}
		err := h.sumanProxy.ScheduleScriptRun(*auth, systemId, 240, commands)
		if err != nil {
			return err
		}

	} else {
		h.logger.Info(fmt.Sprintf("There are no scripts to be executed for phase %s", phase), zf...)
	}
	h.logger.Debug("Function doUpdateActions finished", zf...)
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
	err := h.sumanProxy.SystemScheduleApplyStates(*auth, systemId, states, 600)
	if err != nil {
		zf = append(zf, zap.Any("Error", err))
		h.logger.Error("Highstate failed", zf...)
		return err
	}
	h.logger.Debug("Function executeStates finished", zf...)
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
		err := h.sumanProxy.ScheduleScriptRun(*auth, systemId, config.GetConfig().Suman.Timeout, script)
		if err != nil {
			zf = append(zf, zap.Any("Error", err))
			h.logger.Error("Highstate failed", zf...)
			return err
		}
	}
	h.logger.Debug("Function executeScripts finished", zf...)
	return nil
}

func (h SystemUpdate) getSPFromChannel(channelName string) string {
	sp := strings.Split(channelName, "sp")[1]
	sp = strings.Split(sp, "-")[0]
	if len(sp) == 0 {
		sp = "0"
	}
	return fmt.Sprintf("sp%s", sp)
}

// checkForSpMigration
//
// param: systemId
// param: auth
// param: zf
// return:
func (h SystemUpdate) checkForSpMigration(systemId int, auth *_sumanUseCase.AuthParams, zf ...zapCore.Field) (bool, string) {
	h.logger.Debug("Function echeckForSpMigration started", zf...)
	serverInfo, err := h.sumanProxy.SystemGetSubscribedBaseChannel(*auth, systemId)
	if err != nil {
		return false, ""
	}
	if !contains.PartOff(serverInfo.Label, "sle") && !contains.PartOff(serverInfo.Label, "opensuse") {
		h.logger.Warn("System is not running SLE or openSuse. SP Migration not possible", zf...)
		return false, ""
	}
	for projectOld, projectNew := range config.GetConfig().Maintenance.SpMigrationProject {
		if strings.HasPrefix(serverInfo.Label, projectOld) {
			newBaseChannel := strings.Replace(serverInfo.Label, projectOld, projectNew.(string), 1)
			oldSP := h.getSPFromChannel(serverInfo.Label)
			newProjectChannels, err := h.getProjectChannels(projectNew.(string), auth, zf...)
			newSP := h.getSPFromChannel(newProjectChannels[0])
			newBaseChannel = strings.Replace(newBaseChannel, oldSP, newSP, -1)
			if err != nil {
				return false, ""
			}
		}
	}

	h.logger.Debug("Function echeckForSpMigration finished", zf...)
	return false, ""
}

/*
def check_for_sp_migration():
    """
    Check if a sp migration is released for this server
    """
    current_version = None
    current_bc = smt.system_getsubscribedbasechannel().get('label')
    if "sle" not in current_bc or "opensuse" not in current_bc:
        smt.log_info("System is not running SLE. SP Migration not possible")
        return False, ""
    if "sp" not in current_bc:
        current_sp = "sp0"
    else:
        current_sp = "sp" + str(current_bc.split("sp")[1].split("-")[0])
    all_bc = smt.get_labels_all_basechannels()
    if smtools.CONFIGSM['maintenance']['sp_migration_project']:
        for project, new_pr in smtools.CONFIGSM['maintenance']['sp_migration_project'].items():
            if server_is_exception(new_pr):
                return False, ""
            project_environments = smt.contentmanagement_listprojectenvironment(project, True)
            if project_environments:
                for env in smt.contentmanagement_listprojectenvironment(project, True):
                    calc_current_bc = project + "-" + env['label']
                    if calc_current_bc in current_bc:
                        part_new_bc = calc_current_bc.replace(project, new_pr)
                        new_base_channel = None
                        for bc in all_bc:
                            if part_new_bc in bc:
                                new_base_channel = bc
                                remove_ltss()
                                return True, new_base_channel
                        if not new_base_channel:
                            smt.log_info("Given SP Migration path is not available. There are no channels available.")
                            return False, ""
    if smtools.CONFIGSM['maintenance']['sp_migration']:
        #if "11-" in current_bc:
        #    current_version = "sles11-"
        #elif "12-" in current_bc:
        #    current_version = "sles12-"
        #elif "15-" in current_bc:
        #    current_version = "sles15-"
        #current_version += current_sp
        for key, value in smtools.CONFIGSM['maintenance']['sp_migration'].items():
            if key == current_bc and not server_is_exception(value):
                return True, value
    return False, ""
*/

func (h SystemUpdate) getProjectChannels(projectNew string, auth *_sumanUseCase.AuthParams, zf ...zapCore.Field) ([]string, error) {
	h.logger.Debug("Function getProjectChannels started", zf...)
	var projectChannels []string
	projectInfo, err := h.sumanProxy.ContentManagementListProjectSources(*auth, projectNew)
	if err != nil {
		h.logger.Error(fmt.Sprintf("Error getting channels for project %s", projectNew))
		return projectChannels, err
	}
	for _, info := range projectInfo {
		projectChannels = append(projectChannels, info.ChannelLabel)
	}
	return projectChannels, nil
}
