// Test AmazonCloudDrive filesystem interface

//go:build acd
// +build acd

package amazonclouddrive_test

import (
	"testing"

	"github.com/MilkGames/rclone/backend/amazonclouddrive"
	"github.com/MilkGames/rclone/fs"
	"github.com/MilkGames/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.NilObject = fs.Object((*amazonclouddrive.Object)(nil))
	fstests.RemoteName = "TestAmazonCloudDrive:"
	fstests.Run(t)
}
