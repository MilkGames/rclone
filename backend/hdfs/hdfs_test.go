// Test HDFS filesystem interface

//go:build !plan9
// +build !plan9

package hdfs_test

import (
	"testing"

	"github.com/MilkGames/rclone/backend/hdfs"
	"github.com/MilkGames/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestHdfs:",
		NilObject:  (*hdfs.Object)(nil),
	})
}
