package set

import (
	"log"

	"github.com/hpcng/warewulf/internal/pkg/container"
	"github.com/hpcng/warewulf/internal/pkg/kernel"
	"github.com/hpcng/warewulf/internal/pkg/node"
	"github.com/hpcng/warewulf/internal/pkg/overlay"
	"github.com/spf13/cobra"
)

var (
	baseCmd = &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "set [OPTIONS] PATTERN [PATTERN ...]",
		Short:                 "Configure node properties",
		Long:                  "This command sets configuration properties for nodes matching PATTERN.\n\nNote: use the string 'UNSET' to remove a configuration",
		Args:                  cobra.MinimumNArgs(0),
		RunE:                  CobraRunE,
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) != 0 {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}

			nodeDB, _ := node.New()
			nodes, _ := nodeDB.FindAllNodes()
			var node_names []string
			for _, node := range nodes {
				node_names = append(node_names, node.Id.Get())
			}
			return node_names, cobra.ShellCompDirectiveNoFileComp
		},
	}
	SetNetDevDel string
	SetNodeAll   bool
	SetYes       bool
	SetForce     bool
	OptionStrMap map[string]*string
	KernelStrMap map[string]*string
)

func init() {
	//var emptyNodeConf node.NodeConf
	//var emptyKernelE node.KernelEntry
	myBase := node.CobraCommand{baseCmd}
	var emptyNodeConf node.NodeConf
	emptyNodeConf.Kernel = new(node.KernelConf)
	emptyNodeConf.Ipmi = new(node.IpmiConf)
	//emptyNodeConf.NetDevs = make(map[string]*node.NetDevs)
	OptionStrMap = myBase.CreateFlags(emptyNodeConf)

	baseCmd.PersistentFlags().StringVarP(&SetNetDevDel, "netdel", "D", "", "Delete the node's network device")
	baseCmd.PersistentFlags().BoolVarP(&SetNodeAll, "all", "a", false, "Set all nodes")
	baseCmd.PersistentFlags().BoolVarP(&SetYes, "yes", "y", false, "Set 'yes' to all questions asked")
	baseCmd.PersistentFlags().BoolVarP(&SetForce, "force", "f", false, "Force configuration (even on error)")
	// register the command line completions
	if err := baseCmd.RegisterFlagCompletionFunc("container", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		list, _ := container.ListSources()
		return list, cobra.ShellCompDirectiveNoFileComp
	}); err != nil {
		log.Println(err)
	}
	if err := baseCmd.RegisterFlagCompletionFunc("kerneloverride", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		list, _ := kernel.ListKernels()
		return list, cobra.ShellCompDirectiveNoFileComp
	}); err != nil {
		log.Println(err)
	}
	if err := baseCmd.RegisterFlagCompletionFunc("runtime", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		list, _ := overlay.FindOverlays()
		return list, cobra.ShellCompDirectiveNoFileComp
	}); err != nil {
		log.Println(err)
	}
	if err := baseCmd.RegisterFlagCompletionFunc("wwinit", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		list, _ := overlay.FindOverlays()
		return list, cobra.ShellCompDirectiveNoFileComp
	}); err != nil {
		log.Println(err)
	}
	if err := baseCmd.RegisterFlagCompletionFunc("profile", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var list []string
		nodeDB, _ := node.New()
		profiles, _ := nodeDB.FindAllProfiles()
		for _, profile := range profiles {
			list = append(list, profile.Id.Get())
		}
		return list, cobra.ShellCompDirectiveNoFileComp
	}); err != nil {
		log.Println(err)
	}

}

// GetRootCommand returns the root cobra.Command for the application.
func GetCommand() *cobra.Command {
	return baseCmd
}
