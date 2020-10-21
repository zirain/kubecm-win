package parse

import (
	"fmt"
	"io/ioutil"
	"os/user"

	"github.com/spf13/cobra"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func RunParse(cmd *cobra.Command, args []string) {
	configFile, err := cmd.Flags().GetString("file")
	if err != nil {
		fmt.Println(err)
		return
	}

	config, err := LoadClientConfig(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(config)
}

func LoadClientConfig(kubeconfig string) (*api.Config, error) {
	if kubeconfig == "" {
		currentUser, err := user.Current()
		if err != nil {
			return nil, err
		}

		kubeconfig = fmt.Sprintf("%s\\.kube\\config", currentUser.HomeDir)
	}

	b, err := ioutil.ReadFile(kubeconfig)
	if err != nil {
		return nil, err
	}

	if len(b) == 0 {
		return nil, fmt.Errorf("failed to read file %s", kubeconfig)
	}

	config, err := clientcmd.Load(b)
	if err != nil {
		return nil, err
	}
	return config, nil
}
