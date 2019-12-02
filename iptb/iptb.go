package main

import (
	"fmt"
	"os"

	cli "github.com/ipfs/iptb/cli"
	testbed "github.com/ipfs/iptb/testbed"

	browser "github.com/ipfs/iptb-plugins/browser"
	docker "github.com/ipfs/iptb-plugins/docker"
	local "github.com/ipfs/iptb-plugins/local"
	localp2pd "github.com/ipfs/iptb-plugins/localp2pd"
	localbtfs "github.com/ipfs/iptb-plugins/localbtfs"
)

func init() {
	_, err := testbed.RegisterPlugin(testbed.IptbPlugin{
		From:        "<builtin>",
		NewNode:     local.NewNode,
		GetAttrList: local.GetAttrList,
		GetAttrDesc: local.GetAttrDesc,
		PluginName:  local.PluginName,
		BuiltIn:     true,
	}, false)

	if err != nil {
		panic(err)
	}

	_, err = testbed.RegisterPlugin(testbed.IptbPlugin{
		From:        "<builtin>",
		NewNode:     localp2pd.NewNode,
		GetAttrList: localp2pd.GetAttrList,
		GetAttrDesc: localp2pd.GetAttrDesc,
		PluginName:  localp2pd.PluginName,
		BuiltIn:     true,
	}, false)

	if err != nil {
		panic(err)
	}

	_, err = testbed.RegisterPlugin(testbed.IptbPlugin{
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

	_, err = testbed.RegisterPlugin(testbed.IptbPlugin{
		From:        "<builtin>",
		NewNode:     docker.NewNode,
		GetAttrList: docker.GetAttrList,
		GetAttrDesc: docker.GetAttrDesc,
		PluginName:  docker.PluginName,
		BuiltIn:     true,
	}, false)

	if err != nil {
		panic(err)
	}

	_, err = testbed.RegisterPlugin(testbed.IptbPlugin{
		From:       "<builtin>",
		NewNode:    browser.NewNode,
		PluginName: browser.PluginName,
		BuiltIn:    true,
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
