package qdrant

import (
	"context"

	pb "github.com/qdrant/go-client/qdrant"
)

type QdrantItem struct {
	Vector             []float32
	NormalisedMerchant string
	Link               string
}

func (c *Qdrant) Upsert(ctx context.Context, item *QdrantItem) error {
	pointsClient := pb.NewPointsClient(c.Conn)
	point := &pb.PointStruct{
		Vectors: &pb.Vectors{
			VectorsOptions: &pb.Vectors_Vector{
				Vector: &pb.Vector{
					Data: item.Vector,
				},
			},
		},
		Payload: map[string]*pb.Value{
			"link": {
				Kind: &pb.Value_StringValue{StringValue: item.Link},
			},
			"name": {
				Kind: &pb.Value_StringValue{StringValue: item.NormalisedMerchant},
			},
		},
	}
	waitUpsert := true
	_, err := pointsClient.Upsert(ctx, &pb.UpsertPoints{
		CollectionName: c.Collection,
		Wait:           &waitUpsert,
		Points:         []*pb.PointStruct{point},
	})

	return err
}
