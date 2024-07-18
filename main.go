package main

import "github.com/TheMarstonConnell/docute/cmd"

func main() {
	root := cmd.RootCMD()
	root.AddCommand(cmd.GenerateCMD(), cmd.WatchCMD(), cmd.HostCMD(), cmd.GenColorFile())

	cmd.Execute(root)
}
