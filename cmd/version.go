package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// Version information set by ldflags at build time
var (
	// Version is the semantic version (e.g., "1.0.0")
	Version = "dev"
	// Commit is the git short hash (e.g., "abc1234")
	Commit = "none"
	// Date is the build date (e.g., "2024-01-01T00:00:00Z")
	Date = "unknown"
)

// GetVersion returns the full version string in X.Y.Z-HASH format
func GetVersion() string {
	if Version == "dev" {
		return fmt.Sprintf("%s-%s", Version, Commit)
	}
	return fmt.Sprintf("%s-%s", Version, Commit)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long:  `Print the version, commit hash, build date, and runtime information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("dvcx version %s\n", GetVersion())
		fmt.Printf("  Commit:     %s\n", Commit)
		fmt.Printf("  Built:      %s\n", Date)
		fmt.Printf("  Go version: %s\n", runtime.Version())
		fmt.Printf("  OS/Arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
