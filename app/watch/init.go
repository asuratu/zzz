package watch

import (
	"io/ioutil"

	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zzz/app/root"
	"github.com/sohaha/zzz/util"
	"github.com/spf13/cobra"
)

var (
	initCmdUse = "init"
	initCmd    = &cobra.Command{
		Use:   initCmdUse,
		Short: "Generate config file",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, _ := cmd.Flags().GetString("cfg")
			path := zfile.RealPath(cfg)
			if zfile.FileExist(path) && !force {
				util.Log.Fatal("The configuration file already exists. If you need to override it, use --force")
			}
			err := initCfg(path)
			if err != nil {
				util.Log.Fatal(err)
			}
			util.Log.Successf("create %s successful\n", path)
		},
	}
	force bool
)

func InitCmd(watchCmd *cobra.Command) {
	watchCmd.AddCommand(initCmd)
}

func init() {
	initCmd.Flags().BoolVarP(&force, "force", "F", false, "Override config file")
}

func initCfg(path string) error {
	version := util.Version
	config := root.GetExampleWatchConfig(version)

	return ioutil.WriteFile(path, []byte(config), 0644)
}
