// Test Zoho filesystem interface
package zoho_test

import (
	"testing"

	"github.com/MilkGames/rclone/backend/zoho"
	"github.com/MilkGames/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestZoho:",
		NilObject:  (*zoho.Object)(nil),
	})
}
