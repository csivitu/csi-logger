package schemas

type LogEntrySchema struct {
	Level     string `json:"level"`
	Message   string `json:"title"`
	Path      string `json:"path"`
	Timestamp string `json:"timestamp"`
}

type LogFetchSchema struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
