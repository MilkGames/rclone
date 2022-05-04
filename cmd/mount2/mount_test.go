//go:build linux || (darwin && amd64)
// +build linux darwin,amd64

package mount2

import (
	"testing"

	"github.com/MilkGames/rclone/fstest/testy"
	"github.com/MilkGames/rclone/vfs/vfstest"
)

func TestMount(t *testing.T) {
	testy.SkipUnreliable(t)
	vfstest.RunTests(t, false, mount)
}
