package qdrant

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Qdrant struct {
	Conn       *grpc.ClientConn
	Collection string
}

func New(cfg *QdrantConfig) (*Qdrant, error) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, cfg.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Qdrant{
		Conn:       conn,
		Collection: cfg.Collection,
	}, nil
}
