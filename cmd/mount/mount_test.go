//go:build linux || freebsd
// +build linux freebsd

package mount

import (
	"runtime"
	"testing"

	"github.com/MilkGames/rclone/vfs/vfstest"
)

func TestMount(t *testing.T) {
	if runtime.NumCPU() <= 2 {
		t.Skip("FIXME skipping mount tests as they lock up on <= 2 CPUs - See: https://github.com/MilkGames/rclone/issues/3154")
	}
	vfstest.RunTests(t, false, mount)
}
