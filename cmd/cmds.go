package cmd

import (
	"log"
	"fmt"
	"os"

	"github.com/kuking/jbov"
	"github.com/spf13/cobra"
)

var Verbose bool
var YesMan bool
var DryRun bool

func ErrAndEnd(exitcode int, msg string) {
	fmt.Println("Error:", msg)
	os.Exit(exitcode)
}

func RegisterCommands() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(mountCmd)
	RootCmd.AddCommand(checkCmd)
	RootCmd.AddCommand(syncCmd)
	RootCmd.AddCommand(setCmd)
	RootCmd.AddCommand(statsCmd)
	RootCmd.AddCommand(rebalanceCmd)

	RegisterCreateCommands(RootCmd)
	RegisterRuleCommands(RootCmd)

	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Verbose output")
	RootCmd.PersistentFlags().BoolVarP(&YesMan, "yes", "y", false, "Automatically answers yes (dangerous)")
	RootCmd.PersistentFlags().BoolVarP(&YesMan, "dry-run", "n", false, "Shows what would be done, without applying any change.")
}

// DoitCmd is the base command.
var RootCmd = &cobra.Command{
	Use:   "jbov",
	Short: jbov.Description,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of jbov",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(jbov.FullName, jbov.Version)
	},
}

var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "Mounts a jbov",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("not implemented yet")
	},
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks a jbov",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("not implemented yet")
	},
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync a jbov",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("not implemented yet")
	},
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a flag in a jbov",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("not implemented yet")
	},
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Display statistics about a jbov",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("not implemented yet")
	},
}

var rebalanceCmd = &cobra.Command{
	Use:   "rebalance",
	Short: "Rebalance a jbov",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("not implemented yet")
	},
}
