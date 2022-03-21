package common

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/y-akahori-ramen/test-go-plugin/pb"
	"google.golang.org/grpc"
)

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "TEST_PLUGIN",
	MagicCookieValue: "test",
}

// Calculater　プラグインとして扱いたいインターフェ-ス
type Calculater interface {
	Calc(lhs, rhs float32) (float32, error)
}

// CalculaterPlugin go-pluginのGRPCPluginインターフェース実装
type CalculaterPlugin struct {
	plugin.Plugin
	impl Calculater
}

func NewCalculaterPlugin(impl Calculater) *CalculaterPlugin {
	return &CalculaterPlugin{impl: impl}
}

// GRPCServer go-pluginのGRPCPluginインターフェース実装 GRPCServerを登録する
func (p *CalculaterPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	pb.RegisterCalculaterServer(s, NewCalculaterServer(p.impl))
	return nil
}

// GRPCServer go-pluginのGRPCPluginインターフェース実装 GRPCClientを作る
func (p *CalculaterPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &CalculaterClient{client: pb.NewCalculaterClient(c)}, nil
}
