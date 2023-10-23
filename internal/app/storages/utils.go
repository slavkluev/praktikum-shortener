package storages

type Record struct {
	ID      uint64
	User    string
	URL     string
	Deleted bool
}

type BatchRecord struct {
	ID            uint64
	User          string
	URL           string
	CorrelationID string
}
