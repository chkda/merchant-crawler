package sql

import (
	"context"
)

type Row struct {
	Id      string `db:"ID"`
	Pattern string `db:"Pattern"`
}

func (s *SQLConnector) GetUnmatchedPattern(ctx context.Context) ([]*Row, error) {
	query := "SELECT ID, Pattern FROM " + TABLE + " WHERE IsFound=0 "
	patterns := make([]*Row, 0, 5)
	row, err := s.DB.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		pattern := &Row{}
		err := row.StructScan(pattern)
		if err != nil {
			continue
		}
		patterns = append(patterns, pattern)

	}
	return patterns, nil
}

func (s *SQLConnector) UpdateFoundQuery(ctx context.Context, id string) error {
	query := "UPDATE " + TABLE + " SET IsFound = ? WHERE id = ? "
	_, err := s.DB.ExecContext(ctx, query, 1, id)
	return err
}

func (s *SQLConnector) InsertPatternQuery(ctx context.Context, pattern string) error {
	query := "INSERT INTO " + TABLE + " (Pattern, IsFound) VALUES (?, ?) "
	_, err := s.DB.ExecContext(ctx, query, pattern, 0)
	return err
}
