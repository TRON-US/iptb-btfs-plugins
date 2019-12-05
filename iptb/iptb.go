package main

import (
	"fmt"
	"os"

	cli "github.com/ipfs/iptb/cli"
	testbed "github.com/ipfs/iptb/testbed"

	localbtfs "github.com/TRON-US/iptb-btfs-plugins/localbtfs"
)

func init() {

	_, err := testbed.RegisterPlugin(testbed.IptbPlugin{
		From:        "<builtin>",
		NewNode:     localbtfs.NewNode,
		GetAttrList: localbtfs.GetAttrList,
		GetAttrDesc: localbtfs.GetAttrDesc,
		PluginName:  localbtfs.PluginName,
		BuiltIn:     true,
	}, false)

	if err != nil {
		panic(err)
	}

}

func main() {
	cli := cli.NewCli()
	if err := cli.Run(os.Args); err != nil {
		fmt.Fprintf(cli.ErrWriter, "%s\n", err)
		os.Exit(1)
	}
}
