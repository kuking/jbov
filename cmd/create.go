package cmd

import (
	"fmt"
	"encoding/json"
	"strings"
	"os/user"
	"path/filepath"
	"github.com/kuking/jbov/api"
	"github.com/kuking/jbov/api/md"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create name [vol_alias:/path]...",
	Short: "Creates a new a jbov",
	Run: func(cmd *cobra.Command, args []string) {

		if (len(args) < 2) {
			ErrAndEnd(-1, "you need at least two parameters, the jbov name and at least one initial volume.")
		}

		var jbov = md.JBOV{
			Cname:          args[0],
			Uniqid:         api.GenerateJbovUniqId(),
			LastMountPoint: "",
			Volumes:        make(map[string]md.Volume),
		}

		for i := 1; i < len(args); i++ {
			addVolumeOutOfArg(&jbov, args[i])
		}

		err := api.Create(&jbov)
		if err != nil {
			ErrAndEnd(-1, err.Error())
		}


		json, err := json.MarshalIndent(jbov, "", "    ")
		if err != nil {
			ErrAndEnd(-1, "Failed to serialise Json")
		}
		fmt.Printf("%s\n", json)
	},


}

func addVolumeOutOfArg(jbov *md.JBOV, arg string) {
	var splited = strings.Split(arg, ":")
	if len(splited) != 2 {
		ErrAndEnd(-1, "Volume descriptor should have the format: volumename:/path/to/it")
	}

	var cname = splited[0]
	var mountPoint = splited[1]

	usr, err := user.Current()
	if err != nil {
		ErrAndEnd(-1, "Could not retrieve home: " + err.Error())
	}

	if len(mountPoint) == 0 || mountPoint[0] == '~' {
		// expands ~ as it is inside a string and it will not be expanded by the shell
		mountPoint = filepath.Join(usr.HomeDir, mountPoint[1:])
	}

	mountPoint, err = filepath.Abs(filepath.Clean(mountPoint))
	if err != nil {
		ErrAndEnd(-1, "Path not valid: "+err.Error())
	}

	jbov.Volumes[cname] = md.Volume{
		Uniqid:         api.GenerateVolumeUniqId(),
		LastMountPoint: mountPoint,
		Deprecated:     false,
	}
}

func RegisterCreateCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(createCmd)
}
