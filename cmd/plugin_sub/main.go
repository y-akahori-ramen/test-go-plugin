package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/y-akahori-ramen/test-go-plugin/common"
)

type SubCalculater struct{}

func (c *SubCalculater) Calc(lhs, rhs float32) (float32, error) {
	return lhs - rhs, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: common.Handshake,
		Plugins: map[string]plugin.Plugin{
			"calc": common.NewCalculaterPlugin(&SubCalculater{}),
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
