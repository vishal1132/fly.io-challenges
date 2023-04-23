package utils

import (
	"sync"
	"testing"
)

var lock sync.Mutex
var num int

func TestGetNodeIDInt(t *testing.T) {
	tests := []struct {
		nid string
		id  int
	}{
		{"n1", 1},
		{"n2", 2},
		{"n100", 100},
	}
	for _, v := range tests {
		if id := GetNodeIDInt(v.nid); id != v.id {
			t.Errorf("GetNodeIDInt(%s) = %d, want %d", v.nid, id, v.id)
		}
	}
}

type entry struct {
	node     int64
	sequence int64
}

func TestDuplicate(t *testing.T) {
	maps := make(map[int64]entry)
	var wg sync.WaitGroup
	for i := 1; i <= 350000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			num++
			for j := 1; j <= 35; j++ {
				id := GetUniqueID(int64(j), int64(num))
				if v, ok := maps[id]; ok {
					t.Fatalf("Duplicate ID: %d, entry: %v", id, v)
				} else {
					maps[id] = entry{node: int64(j), sequence: int64(num)}
				}
			}
			lock.Unlock()
		}()
	}
	wg.Wait()
}
