// Package susemanager - SUSE Manager api call and support functions
package susemanager

import (
	sumamodels "ecp-golang-cm/pkg/models/susemanager"
	"ecp-golang-cm/pkg/util/rest"

	"go.uber.org/zap"
)

// all public interface here for the model

// ISuseManager - description
type ISuseManager interface {
	GetSystemGroupName(negName string) string
	SetK3sDetails(requestID string, auth AuthParams, systemgroupName string, k3sconfiginput map[string]interface{}) error
	ChangeChannels(requestID string, auth AuthParams, systemID int, targetedVersion string) error
	GetHost(requestID string, negName string, sessionKey string) (*AuthParams, error)
	InstallPackages(requestID string, auth AuthParams, systemID int, pkgs []string, timeout int) error
	GetAuth(sessionkey string) (*AuthParams, error)
}

// ISuseManagerAPI - description
type ISuseManagerAPI interface {
	SuseManagerCall(body []byte, method string, hostname string, path string, sessionKey string) (output *rest.HTTPHelperStruct, er error)
}

type (
	// SumanConfig - SUSE Manager credentials
	SumanConfig struct {
		Host     string
		Login    string
		Password string
		Insecure bool
	}
	// Proxy - general needed information
	Proxy struct {
		cfg               *SumanConfig
		contentTypeHeader map[string]string
		suse              ISuseManagerAPI
		logger            *zap.Logger
		retrycount        int
	}
)

// NewProxy - ope new connection
//
// param: s
// param: suse
// param: logger
// param: retrycount
// return:
func NewProxy(s *SumanConfig, suse ISuseManagerAPI, logger *zap.Logger, retrycount int) IProxy {
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	return &Proxy{
		cfg:               s,
		contentTypeHeader: header,
		suse:              suse,
		logger:            logger,
		retrycount:        retrycount,
	}
}

// IProxy - all public interface here for the model
type IProxy interface {
	// activationkey
	ActivationKeyAddChildChannels(requestID string, auth AuthParams, keyName string, childChannels []string) (int, error)
	ActivationKeyAddServerGroups(requestID string, auth AuthParams, keyName string, groups []int) (int, error)
	ActivationKeyCreate(requestID string, auth AuthParams, keyName string, baseChannel string, entitlement []string) (string, error)
	ActivationKeyDelete(requestID string, auth AuthParams, keyName string) (int, error)
	ActivationKeyGetDetails(requestID string, auth AuthParams, keyName string) (sumamodels.ActivationkeyGetDetails, error)
	ActivationKeyListActivationKeys(requestID string, auth AuthParams) ([]sumamodels.ActivationkeyGetDetails, error)
	ActivationKeyRemovePackages(requestID string, auth AuthParams, keyName string, pckgs []sumamodels.ActivationkeyPackages) (int, error)

	// authentication
	GetSessionKey(body []byte, host string, requestID string) (string, error)
	SumanLogin(requestID string) (string, error)
	SumanLogout(requestID string, auth AuthParams) error
	CheckResponseProgress(requestID string, auth AuthParams, response *rest.HTTPHelperStruct, timeOut int, systemID int, funcName string) error

	// configchannel
	ConfigChannelListGlobals(requestID string, auth AuthParams) ([]sumamodels.ConfigChannelListGlobals, error)

	// contentmanagement
	ContentManagementListProjects(requestID string, auth AuthParams) ([]sumamodels.ContentManagementListProjects, error)
	ContentManagementCreate(requestID string, auth AuthParams, projectLabel string, name string, description string) (sumamodels.ContentManagementListProjects, error)
	ContentManagementAttachSource(requestID string, auth AuthParams, projectLabel string, sourceType string, sourceLabel string) (sumamodels.ContentManagementSource, error)
	ContentManagementListFilters(requestID string, auth AuthParams) ([]sumamodels.ContentManagementFilter, error)
	ContentManagementCreateFilter(requestID string, auth AuthParams, name string, rule string, entityType string, criteria sumamodels.FilterCriteria) (sumamodels.ContentManagementFilter, error)
	ContentManagementAttachFilter(requestID string, auth AuthParams, projectLabel string, filterID int) (sumamodels.ContentManagementFilter, error)
	ContentManagementCreateEnvironment(requestID string, auth AuthParams, projectLabel string, predecessorLabel string, envlabel string, name string, description string) (sumamodels.ContentManagementEnvironment, error)
	ContentManagementBuildProject(requestID string, auth AuthParams, projectLabel string) (int, error)

	// system
	SystemGetID(requestID string, auth AuthParams, systemName string) ([]sumamodels.System, error)
	CheckProgress(requestID string, auth AuthParams, actionID int, timeout int, action string, systemID int) (int, error)
	ListCompleteSystem(requestID string, auth AuthParams, actionID int) ([]interface{}, error)
	ListInprogressSystem(requestID string, auth AuthParams, actionID int) ([]interface{}, error)
	ListLatestInstallablePackages(requestID string, auth AuthParams, systemID int) ([]sumamodels.InstallablePackage, error)
	SchedulePackageRefresh(requestID string, auth AuthParams, systemID int) error
	ScheduleScriptRun(requestID string, auth AuthParams, systemID int, timeout int, script string) error
	SystemGetScriptResult(requestID string, auth AuthParams, actionID int, resultCompleted int) (string, error)
	SystemListActiveSystems(requestID string, auth AuthParams) ([]sumamodels.ActiveSystem, error)
	SystemListInstalledPackages(requestID string, auth AuthParams, systemID int) ([]sumamodels.InstalledPackage, error)
	SystemScheduleApplyHighstate(requestID string, auth AuthParams, systemID int, timeout int) error
	SystemScheduleApplyStates(requestID string, auth AuthParams, systemID int, stateNames []string, timeout int) error
	SystemScheduleChangeChannels(requestID string, auth AuthParams, systemID int, basechannel string, childChannel []sumamodels.ChannelSoftwareListChildren) error
	SystemGetSubscribedBaseChannel(requestID string, auth AuthParams, systemID int) (sumamodels.SubscribedBaseChannel, error)
	SystemScheduleReboot(requestID string, auth AuthParams, systemID int, timeout int) error

	// sync
	GetSlaves(requestID string, sessionKey string) ([]sumamodels.Slaves, error)
	SyncSlaveGetSlaveByName(requestID string, auth AuthParams, slaveFQDN string) (sumamodels.Slaves, error)
	SyncSlaveDelete(requestID string, auth AuthParams, slaveID int) (int, error)
	SyncSlaveCreate(requestID string, auth AuthParams, slaveFQDN string, isEnabled bool, allowAllOrgs bool) (sumamodels.Slaves, error)
	SyncMasterGetMasterByLabel(requestID string, auth AuthParams, slaveFQDN string) (sumamodels.SlavesIssMaster, error)
	SyncMasterDelete(requestID string, auth AuthParams, masterID int) (int, error)
	SyncMasterCreate(requestID string, auth AuthParams, masterFQDN string) (sumamodels.SlavesIssMaster, error)
	SyncMasterMakeDefault(requestID string, auth AuthParams, masterID int) (int, error)
	SyncMasterSetCaCert(requestID string, auth AuthParams, masterID int, caCert string) (int, error)

	// Channels
	ChannelListSoftwareChannels(requestID string, auth AuthParams) ([]sumamodels.ChannelListSoftwareChannels, error)
	ChannelSoftwareListChildren(requestID string, auth AuthParams, label string) ([]sumamodels.ChannelSoftwareListChildren, error)
	ChannelSoftwareCreateRepo(requestID string, auth AuthParams, label string, typeRepo string, url string) (sumamodels.ChannelSoftwareCreateRepo, error)
	ChannelSoftwareCreate(requestID string, auth AuthParams, label string, name string, summary string, archLabel string, parentLabel string) (int, error)
	ChannelSoftwareAssociateRepo(requestID string, auth AuthParams, channelLabel string, repoLabel string) (sumamodels.ChannelSoftwareListChildren, error)
	ChannelSoftwareSyncRepo(requestID string, auth AuthParams, channelLabel string) (int, error)
	ChannelSoftwareIsExisting(requestID string, auth AuthParams, label string) (bool, error)
	//	SystemGetSubscribedBaseChannel(auth AuthParams, systemID int) (*sumamodels.SubscribedChannel, error)

	// Add func for formula
	GetFormulasByServerID(requestID string, auth AuthParams, systemID int) ([]string, error)
	GetFormulasByGroupID(requestID string, auth AuthParams, groupID int) ([]string, error)
	GetGroupFormulaData(auth AuthParams, groupID int, formulaName string) (interface{}, error)
	GetSystemFormulaData(auth AuthParams, sid int, formulaname string) (interface{}, error)
	SetGroupFormulaData(requestID string, auth AuthParams, groupID int, formulaName string, formulaData interface{}) (int, error)
	SetSystemFormulaData(requestID string, auth AuthParams, systemID int, formulaName string, formulaData interface{}) (int, error)
	FormulaSetFormulasOfGroup(requestID string, auth AuthParams, systemID int, formulaNames []string) (int, error)
	FormulaSetFormulasOfSystem(requestID string, auth AuthParams, systemID int, formulaNames []string) (int, error)

	// SystemGroup
	SystemGroupCreate(requestID string, auth AuthParams, groupName string, description string) (*sumamodels.SystemGroupGetDetails, error)
	SystemGroupGetDetails(requestID string, auth AuthParams, groupName string) (*sumamodels.SystemGroupGetDetails, error)
	SystemGroupListSystemsMinimal(requestID string, auth AuthParams, groupName string) ([]sumamodels.SystemGroupListSystemsMinimal, error)
	SystemGroupListActiveSystemsInGroup(requestID string, auth AuthParams, groupName string) ([]int, error)

	// KickstartTree
	KickstartTreeGetDetails(requestID string, auth AuthParams, distributionName string) (sumamodels.KickstartTreeGetDetails, error)
	KickstartTreeCreate(requestID string, auth AuthParams, treeLabel string, basePath string, channelLabel string, installType string) (int, error)
	KickstartTreeCreateKernelOptions(requestID string, auth AuthParams, treeLabel string, basePath string, channelLabel string, installType string, kernelOptions string, postKernelOptions string) (int, error)
	KickstartImportRawFile(requestID string, auth AuthParams, profileLabel string, virtType string, channelLabel string, dataXML string) (int, error)
	KickstartListKickstarts(requestID string, auth AuthParams) ([]sumamodels.KickstartListProfiles, error)
	KickstartDeleteProfile(requestID string, auth AuthParams, profileName string) (int, error)
	KickstartProfileSetVariables(requestID string, auth AuthParams, profileLabel string, profileVariables interface{}) (int, error)
}
