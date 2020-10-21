package list

import (
	"fmt"
	"kubecm-win/pkg/parse"
	"sort"

	"github.com/spf13/cobra"
)

func RunList(cmd *cobra.Command, args []string) {
	config, err := parse.LoadClientConfig("")
	if err != nil {
		fmt.Println(err)
		return
	}

	contexts := make([]string, 0, len(config.Contexts))
	for c := range config.Contexts {
		contexts = append(contexts, c)
	}

	sort.Strings(contexts)

	for _, c := range contexts {
		fmt.Println(c)
	}
}
