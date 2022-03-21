package common

import (
	"context"

	"github.com/y-akahori-ramen/test-go-plugin/pb"
)

// CalculaterClient gRPCクライアント実装
// Calculaterインターフェースを実装している
type CalculaterClient struct {
	client pb.CalculaterClient
}

// Calc 計算処理 gRPCサーバに左辺と右辺を投げて結果を得る
func (c *CalculaterClient) Calc(lhs, rhs float32) (float32, error) {
	r, err := c.client.Calculate(context.Background(), &pb.Request{Lhs: lhs, Rhs: rhs})
	if err != nil {
		return 0, err
	}
	return r.Result, nil
}
