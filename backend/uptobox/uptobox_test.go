// Test Uptobox filesystem interface
package uptobox_test

import (
	"testing"

	"github.com/MilkGames/rclone/backend/uptobox"
	"github.com/MilkGames/rclone/fstest"
	"github.com/MilkGames/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	if *fstest.RemoteName == "" {
		*fstest.RemoteName = "TestUptobox:"
	}
	fstests.Run(t, &fstests.Opt{
		RemoteName: *fstest.RemoteName,
		NilObject:  (*uptobox.Object)(nil),
	})
}
