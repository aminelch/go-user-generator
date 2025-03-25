package models

type Dependencies struct {
	SQLite SQLiteInfo `json:"sqlite"`
}

type SQLiteInfo struct {
	Status string  `json:"status"`
	File   string  `json:"file"`
	SizeMB float64 `json:"size_mb"`
}

type HealthResponse struct {
	Hostname     string       `json:"hostname"`
	Uptime       float64      `json:"uptime"`
	Status       string       `json:"status"`
	Timestamp    string       `json:"timestamp"`
	Version      string       `json:"version"`
	Dependencies Dependencies `json:"dependencies"`
}
