package version

import (
	"github.com/hashicorp/packer/packer-plugin-sdk/version"
	packerVersion "github.com/hashicorp/packer/version"
)

var LXCPluginVersion *version.PluginVersion

func init() {
	LXCPluginVersion = version.InitializePluginVersion(
		packerVersion.Version, packerVersion.VersionPrerelease)
}
