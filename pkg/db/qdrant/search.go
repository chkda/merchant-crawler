package qdrant

import (
	"context"
	"errors"

	pb "github.com/qdrant/go-client/qdrant"
)

func (c *Qdrant) Search(ctx context.Context, vector []float32) (*pb.ScoredPoint, error) {
	pointsClient := pb.NewPointsClient(c.Conn)
	searchResults, err := pointsClient.Search(
		ctx,
		&pb.SearchPoints{
			CollectionName: c.Collection,
			Vector:         vector,
			Limit:          3,
			WithVectors:    &pb.WithVectorsSelector{SelectorOptions: &pb.WithVectorsSelector_Enable{Enable: true}},
			WithPayload:    &pb.WithPayloadSelector{SelectorOptions: &pb.WithPayloadSelector_Enable{Enable: true}},
		},
	)
	if err != nil {
		return nil, err
	}
	items := searchResults.Result
	if len(items) == 0 {
		return nil, errors.New("item not found")
	}
	matchedItem := items[0]
	return matchedItem, nil
}
