package client

import (
	"fmt"
	"os"

	"github.com/digital-feather/cryptellation/services/exchanges/pkg/client/proto"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New() (client proto.ExchangesServiceClient, close func() error, err error) {
	grpcAddr := os.Getenv("CRYPTELLATION_EXCHANGES_GRPC_URL")
	if grpcAddr == "" {
		return nil, func() error { return nil }, xerrors.New("no grpc url provided")
	}

	conn, err := grpc.Dial(grpcAddr, grpcDialOpts(grpcAddr)...)
	if err != nil {
		return nil, func() error { return nil }, fmt.Errorf("dialing exchanges grpc server: %w", err)
	}

	return proto.NewExchangesServiceClient(conn), conn.Close, nil
}

func grpcDialOpts(grpcAddr string) []grpc.DialOption {
	return []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
}
