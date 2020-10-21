package main

import (
	"kubecm-win/pkg/list"
	"kubecm-win/pkg/parse"
	"os"

	"github.com/spf13/cobra"
)

var (
	configFile string
)

func main() {
	rootCmd := &cobra.Command{
		Use:          "kubecm-win",
		Short:        "kubecm for windows",
		Long:         "kubeconfig mgr tool for windows",
		SilenceUsage: true,
		FParseErrWhitelist: cobra.FParseErrWhitelist{
			UnknownFlags: true,
		},
	}

	listCmd := &cobra.Command{
		Use:     "list",
		Short:   "list kubeconfig",
		Long:    "list kubeconfig",
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, args []string) {
			list.RunList(cmd, args)
		},
		Example: "kubecm-win ls",
	}

	parseCmd := &cobra.Command{
		Use:     "parse",
		Short:   "parse kubeconfig",
		Long:    "parse kubeconfig",
		Aliases: []string{"pa"},
		Run: func(cmd *cobra.Command, args []string) {
			parse.RunParse(cmd, args)
		},
		Example: "kubecm-win parse --file C:\\Users\\zirai\\Downloads\\kubeconfig.json",
	}
	parseCmd.PersistentFlags().StringVarP(&configFile, "file", "f", "", "")

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(parseCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}

}
