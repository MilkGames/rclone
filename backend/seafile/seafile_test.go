// Test Seafile filesystem interface
package seafile_test

import (
	"testing"

	"github.com/MilkGames/rclone/backend/seafile"
	"github.com/MilkGames/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestSeafile:",
		NilObject:  (*seafile.Object)(nil),
	})
}
