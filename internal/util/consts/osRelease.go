package consts

import _osReleaseModels "ecp-golang-cm/pkg/models/createOsRelease"

// ListOsReleaseData - Data needed for OS releases
var ListOsReleaseData = []_osReleaseModels.OsReleaseRecord{
	{
		TreePath:             "/usr/share/tftpboot-installation/SLE-15-SP4-x86_64",
		ParentChannel:        "sle-product-suse-manager-server-4.3-pool-x86_64",
		Label:                "sm43",
		ChildChannelsSpecial: []string{},
		ChildChannelsDefault: []string{"sle-module-basesystem15-sp4-pool-x86_64-sms-4.3", "sle-module-basesystem15-sp4-updates-x86_64-sms-4.3", "sle-module-server-applications15-sp4-pool-x86_64-sms-4.3", "sle-module-server-applications15-sp4-updates-x86_64-sms-4.3", "sle-module-suse-manager-server-4.3-pool-x86_64", "sle-module-suse-manager-server-4.3-updates-x86_64", "sle-module-web-scripting15-sp4-pool-x86_64-sms-4.3", "sle-module-web-scripting15-sp4-updates-x86_64-sms-4.3", "sle-product-suse-manager-server-4.3-updates-x86_64"},
		ChildChannelsExtra:   []string{"general", "sm43-extra"},
	},
	{
		TreePath:             "/usr/share/tftpboot-installation/SLE-Micro-5.2-x86_64",
		ParentChannel:        "suse-microos-5.2-pool-x86_64",
		Label:                "mi52",
		ChildChannelsSpecial: []string{},
		ChildChannelsDefault: []string{"sle-manager-tools-for-micro5-pool-x86_64-5.2", "sle-manager-tools-for-micro5-updates-x86_64-5.2", "suse-microos-5.2-updates-x86_64"},
		ChildChannelsExtra:   []string{"general", "mi52-extra"},
	},
	{
		TreePath:             "/usr/share/tftpboot-installation/SLE-Micro-5.3-x86_64",
		ParentChannel:        "sle-micro-5.3-pool-x86_64",
		Label:                "mi53",
		ChildChannelsSpecial: []string{},
		ChildChannelsDefault: []string{"sle-manager-tools-for-micro5-pool-x86_64-5.3", "sle-manager-tools-for-micro5-updates-x86_64-5.3", "suse-microos-5.3-updates-x86_64"},
		ChildChannelsExtra:   []string{"general", "mi53-extra"},
	},
	{
		TreePath:             "/usr/share/tftpboot-installation/SLE-Micro-5.4-x86_64",
		ParentChannel:        "sle-micro-5.4-pool-x86_64",
		Label:                "mi54",
		ChildChannelsSpecial: []string{},
		ChildChannelsDefault: []string{"sle-manager-tools-for-micro5-pool-x86_64-5.4", "sle-manager-tools-for-micro5-updates-x86_64-5.4", "suse-microos-5.4-updates-x86_64"},
		ChildChannelsExtra:   []string{"general", "mi54-extra"},
	},
	{
		TreePath:             "/usr/share/tftpboot-installation/SLE-Micro-5.5-x86_64",
		ParentChannel:        "sle-micro-5.5-pool-x86_64",
		Label:                "mi55",
		ChildChannelsSpecial: []string{},
		ChildChannelsDefault: []string{"sle-manager-tools-for-micro5-pool-x86_64-5.5", "sle-manager-tools-for-micro5-updates-x86_64-5.5", "sle-micro-5.5-updates-x86_64"},
		ChildChannelsExtra:   []string{"general", "mi55-extra"},
	},
	{
		TreePath:             "/usr/share/tftpboot-installation/SLE-15-SP4-x86_64",
		ParentChannel:        "sle-product-sles15-sp4-pool-x86_64",
		Label:                "s154",
		ChildChannelsSpecial: []string{"sle-module-containers15-sp4-pool-x86_64", "sle-module-containers15-sp4-updates-x86_64", "sle-module-desktop-applications15-sp4-pool-x86_64", "sle-module-desktop-applications15-sp4-updates-x86_64", "sle-module-devtools15-sp4-pool-x86_64", "sle-module-devtools15-sp4-updates-x86_64", "sle-module-packagehub-subpackages15-sp4-pool-x86_64", "sle-module-packagehub-subpackages15-sp4-updates-x86_64", "sle-product-ha15-sp4-pool-x86_64", "sle-product-ha15-sp4-updates-x86_64", "suse-packagehub-15-sp4-backports-pool-x86_64", "suse-packagehub-15-sp4-pool-x86_64"},
		ChildChannelsDefault: []string{"sle-manager-tools15-pool-x86_64-sp4", "sle-manager-tools15-updates-x86_64-sp4", "sle-module-basesystem15-sp4-pool-x86_64", "sle-module-basesystem15-sp4-updates-x86_64", "sle-module-server-applications15-sp4-pool-x86_64", "sle-module-server-applications15-sp4-updates-x86_64", "sle-product-sles15-sp4-updates-x86_64"},
		ChildChannelsExtra:   []string{"general", "s154-extra"},
	},
	{
		TreePath:             "/usr/share/tftpboot-installation/SLE-15-SP5-x86_64",
		ParentChannel:        "sle-product-sles15-sp5-pool-x86_64",
		Label:                "s155",
		ChildChannelsSpecial: []string{"sle-module-containers15-sp5-pool-x86_64", "sle-module-containers15-sp5-updates-x86_64", "sle-module-desktop-applications15-sp5-pool-x86_64", "sle-module-desktop-applications15-sp5-updates-x86_64", "sle-module-devtools15-sp5-pool-x86_64", "sle-module-devtools15-sp5-updates-x86_64", "sle-module-packagehub-subpackages15-sp5-pool-x86_64", "sle-module-packagehub-subpackages15-sp5-updates-x86_64", "sle-product-ha15-sp5-pool-x86_64", "sle-product-ha15-sp5-updates-x86_64", "suse-packagehub-15-sp5-backports-pool-x86_64", "suse-packagehub-15-sp5-pool-x86_64"},
		ChildChannelsDefault: []string{"sle-manager-tools15-pool-x86_64-sp5", "sle-manager-tools15-updates-x86_64-sp5", "sle-module-basesystem15-sp5-pool-x86_64", "sle-module-basesystem15-sp5-updates-x86_64", "sle-module-server-applications15-sp5-pool-x86_64", "sle-module-server-applications15-sp5-updates-x86_64", "sle-product-sles15-sp5-updates-x86_64"},
		ChildChannelsExtra:   []string{"general", "s155-extra"},
	},
}

const BaseChannelExtra string = "dt-repo"
const ExtraRepoDir string = "/srv/repos/"

// CorrectLabels - available OS
var CorrectLabels = []string{"mi52", "mi55", "sm43", "s154", "s155"}

// CorrectEnvironments - available environments
var CorrectEnvironments = []string{"r001"}
