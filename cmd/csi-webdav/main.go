package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zv0n/csi-webdav/pkg/webdav"

	"github.com/spf13/cobra"
)

var (
	endpoint string
	nodeID   string
)

const configPath = "/etc/webdav-proxy.conf"

func init() {
	flag.Set("logtostderr", "true")
}

func main() {

	flag.CommandLine.Parse([]string{})

	cmd := &cobra.Command{
		Use:   "webdav",
		Short: "CSI based webdav driver",
		Run: func(cmd *cobra.Command, args []string) {
			handle()
		},
	}

	cmd.Flags().AddGoFlagSet(flag.CommandLine)

	cmd.PersistentFlags().StringVar(&nodeID, "nodeid", "", "node id")
	cmd.MarkPersistentFlagRequired("nodeid")

	cmd.PersistentFlags().StringVar(&endpoint, "endpoint", "", "CSI endpoint")
	cmd.MarkPersistentFlagRequired("endpoint")

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Prints information about this version of csi webdav plugin",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(`CSI-webdav Plugin
Version:    %s
Build Time: %s
`, webdav.Version, webdav.BuildTime)
		},
	}

	cmd.AddCommand(versionCmd)
	versionCmd.ResetFlags()

	cmd.ParseFlags(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

func handle() {
	d := webdav.NewDriver(nodeID, endpoint, configPath)
	d.Run()
}
