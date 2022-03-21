package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/y-akahori-ramen/test-go-plugin/common"
)

type AddCalculater struct{}

func (c *AddCalculater) Calc(lhs, rhs float32) (float32, error) {
	return lhs + rhs, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: common.Handshake,
		Plugins: map[string]plugin.Plugin{
			"calc": common.NewCalculaterPlugin(&AddCalculater{}),
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
