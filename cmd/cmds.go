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

func RegisterCommands() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(createCmd)
	RootCmd.AddCommand(mountCmd)
	RootCmd.AddCommand(checkCmd)
	RootCmd.AddCommand(syncCmd)
	RootCmd.AddCommand(setCmd)
	RootCmd.AddCommand(statsCmd)
	RootCmd.AddCommand(rebalanceCmd)

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

var createCmd = &cobra.Command{
	Use:   "create name [vol_alias:/path]...",
	Short: "Creates a new a jbov",
	Run: func(cmd *cobra.Command, args []string) {

		if (len(args)<2) {
			fmt.Println("Error: you need at least two parameters, the jbov name and at least one initial volume.")
			os.Exit(-1)
		}





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

