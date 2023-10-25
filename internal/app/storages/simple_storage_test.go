package storages

import "testing"

func BenchmarkCreateSimpleStorage(b *testing.B) {
	syncTime := 5
	filename := "test"
	for i := 0; i < b.N; i++ {
		CreateSimpleStorage(filename, syncTime)
	}
}
