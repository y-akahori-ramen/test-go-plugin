package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"
	"github.com/y-akahori-ramen/test-go-plugin/common"
)

func main() {
	pluginCmd, ok := os.LookupEnv("CALC_PLUGIN")
	if !ok {
		fmt.Print("CALC_PLUGIN is not set")
		os.Exit(1)
	}

	pluginMap := map[string]plugin.Plugin{
		"calc": &common.CalculaterPlugin{},
	}

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  common.Handshake,
		Plugins:          pluginMap,
		Cmd:              exec.Command("sh", "-c", pluginCmd),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	raw, err := rpcClient.Dispense("calc")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	calculater, ok := raw.(common.Calculater)
	if !ok {
		fmt.Println("Error:", "type assertion failed")
		os.Exit(1)
	}

	var lhs, rhs float32
	lhs = 1.0
	rhs = 2.0
	result, err := calculater.Calc(lhs, rhs)
	fmt.Printf("Lhs:%v Rhs:%v Calc result:%v", lhs, rhs, result)

	// exitで終了しないとエラーになる
	os.Exit(0)
}
