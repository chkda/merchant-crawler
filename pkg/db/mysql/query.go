package mysql

import "context"

type Row struct {
	Id      string `db:"ID"`
	Pattern string `db:"Pattern"`
}

func (s *SQLConnector) GetUnmatchedPattern(ctx context.Context) ([]*Row, error) {
	query := "SELECT ID, Pattern FROM " + TABLE + "WHERE IsFound=0 "
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
