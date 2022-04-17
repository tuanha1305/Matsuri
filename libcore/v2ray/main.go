package main

import (
	"fmt"
	"libcore/nekoray_rpc"
	"os"
	_ "unsafe"

	"github.com/v2fly/v2ray-core/v5/main/commands"
)

//go:linkname build github.com/v2fly/v2ray-core/v5.build
var build string

var version_v2ray string = "N/A"
var version_standalone string = "N/A"

func main() {
	fmt.Println("V2Ray:", version_v2ray, "Version:", version_standalone)
	fmt.Println()

	if len(os.Args) > 1 && os.Args[1] == "nekoray" {
		nekoray_rpc.Main()
		return
	}

	build = "Matsuridayo/Nekoray"
	commands.CmdRun.Run(commands.CmdRun, os.Args[1:])
}
