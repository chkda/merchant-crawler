package sql

type SQLConnConfig struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

const (
	TABLE = "unmatched_patterns"
)
