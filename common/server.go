package common

import (
	"context"

	"github.com/y-akahori-ramen/test-go-plugin/pb"
)

// CalculaterServer gRPCサーバー実装
type CalculaterServer struct {
	impl Calculater
	pb.UnimplementedCalculaterServer
}

// NewCalculaterServer CalulateServerを作成する
func NewCalculaterServer(calculaterImpl Calculater) *CalculaterServer {
	return &CalculaterServer{impl: calculaterImpl}
}

// Calculate gRPCのレスポンス実装
func (server *CalculaterServer) Calculate(ctx context.Context, r *pb.Request) (*pb.Reply, error) {
	// 受け取った値を実装インスタンスに渡して結果を得る
	result, err := server.impl.Calc(r.Lhs, r.Rhs)
	return &pb.Reply{Result: result}, err
}
