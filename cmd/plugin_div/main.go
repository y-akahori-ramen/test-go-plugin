package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/y-akahori-ramen/test-go-plugin/common"
)

type DivCalculater struct{}

func (c *DivCalculater) Calc(lhs, rhs float32) (float32, error) {
	return lhs / rhs, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: common.Handshake,
		Plugins: map[string]plugin.Plugin{
			"calc": common.NewCalculaterPlugin(&DivCalculater{}),
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
