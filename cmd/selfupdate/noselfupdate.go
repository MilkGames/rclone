//go:build noselfupdate
// +build noselfupdate

package selfupdate

import (
	"github.com/MilkGames/rclone/lib/buildinfo"
)

func init() {
	buildinfo.Tags = append(buildinfo.Tags, "noselfupdate")
}
