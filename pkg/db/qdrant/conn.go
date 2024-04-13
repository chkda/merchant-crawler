package qdrant

import (
	"context"

	pb "github.com/qdrant/go-client/qdrant"
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
	client := pb.NewCollectionsClient(conn)
	_, err = client.Create(ctx, &pb.CreateCollection{
		CollectionName: cfg.Collection,
		VectorsConfig: &pb.VectorsConfig{Config: &pb.VectorsConfig_Params{
			Params: &pb.VectorParams{
				Size:     vectorSize,
				Distance: pb.Distance_Dot,
			},
		}},
	})
	if err != nil {
		return nil, err
	}
	return &Qdrant{
		Conn:       conn,
		Collection: cfg.Collection,
	}, nil
}
