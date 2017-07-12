package cmd

import (
	"log"
	"github.com/spf13/cobra"
)

var pattern, atLeastACopyIn, volume string
var nCopies, ruleNo int
var deprecated bool

var ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "Manage redundancy rules",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("not implemented yet")
	},
}

var ruleAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new redundancy rule",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("not implemented yet")
	},
}

var ruleDelCmd = &cobra.Command{
	Use:   "del",
	Short: "Removes a redundancy rule",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("not implemented yet")
	},
}

var ruleListCmd = &cobra.Command{
	Use:   "list",
	Short: "List rules",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("not implemented yet")
	},
}

func RegisterRuleCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(ruleCmd)
	ruleCmd.AddCommand(ruleAddCmd)
	ruleCmd.AddCommand(ruleDelCmd)
	ruleCmd.AddCommand(ruleListCmd)


	ruleAddCmd.PersistentFlags().StringVarP(&volume, "volume", "V", "", "Volume to apply the rule to")
	ruleAddCmd.PersistentFlags().StringVarP(&pattern, "pattern", "p", "", "file pattern to apply to the rule")
	ruleAddCmd.PersistentFlags().IntVarP(&nCopies, "ncopies", "n", -1, "Number of copies to maintain, '*' indicates to hold a copy on every volumen.")
	ruleAddCmd.PersistentFlags().StringVarP(&atLeastACopyIn, "at-least-a-copy-in", "a", "", "A redundant copy should be held in the indicated Volume")
	ruleAddCmd.PersistentFlags().BoolVarP(&deprecated, "deprecated", "d", false, "Marks a volume as deprecated (it will not add any new file in it and files in it will not be counted as redundant copies)")

	ruleDelCmd.PersistentFlags().IntVarP(&ruleNo, "ruleno", "n", -1, "Rule number to delete")
}
