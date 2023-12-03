package memory

import "testing"

func BenchmarkNewMemoryRecordRepository(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewMemoryRecordRepository()
	}
}
