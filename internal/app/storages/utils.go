package storages

// Record хранит данные оригинального URL
type Record struct {
	ID      uint64
	User    string
	URL     string
	Deleted bool
}

// BatchRecord хранит данные оригинального URL, используемого при множественном сокращении
type BatchRecord struct {
	ID            uint64
	User          string
	URL           string
	CorrelationID string
}
