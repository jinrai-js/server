package lru

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

func getKeys(l *LRUCache) string {
	return fmt.Sprintf("%v", slices.Collect(maps.Keys(l.cache)))
}

func TestKey(t *testing.T) {
	l := New(2)
	l.Put("1", "1")
	l.Put("2", "2")

	if val, err := l.Get("1"); err != nil || val != "1" {
		t.Error("Должен вернуть 1")
	}

	l.Put("3", "3")
	if getKeys(&l) != "[1 3]" {
		t.Error("должно быть [1 3] -", getKeys(&l))
	}

	l.Put("4", "4")
	if getKeys(&l) != "[4 3]" {
		t.Error("должно быть [3 4] -", getKeys(&l))
	}
}
