package lru

import (
	"fmt"
	"maps"
	"slices"
	"sort"
	"testing"
)

func getKeys(l *LRUCache) string {
	list := slices.Collect(maps.Keys(l.cache))
	sort.Strings(list)
	return fmt.Sprintf("%v", list)
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
	if getKeys(&l) != "[3 4]" {
		t.Error("должно быть [4 3] -", getKeys(&l))
	}

	if val, err := l.Get("1"); err == nil || val != "" {
		t.Error("Должен вернуть ошибку")
	}

}

func TestCapacityValidation(t *testing.T) {
	// Тест с capacity = 0 (должен стать 1)
	l := New(0)
	if l.capacity != 1 {
		t.Error("capacity должен быть минимум 1, получен:", l.capacity)
	}

	// Тест с отрицательным capacity (должен стать 1)
	l2 := New(-5)
	if l2.capacity != 1 {
		t.Error("capacity должен быть минимум 1, получен:", l2.capacity)
	}

	// Тест с capacity = 1
	l3 := New(1)
	l3.Put("1", "1")
	l3.Put("2", "2") // Должен удалить "1"
	if l3.Has("1") {
		t.Error("Ключ '1' должен быть удален")
	}
	if !l3.Has("2") {
		t.Error("Ключ '2' должен существовать")
	}
}

func TestEmptyCache(t *testing.T) {
	l := New(2)
	// Попытка получить элемент из пустого кэша
	if val, err := l.Get("nonexistent"); err == nil || val != "" {
		t.Error("Должна быть ошибка для несуществующего ключа")
	}

	// Проверка popTail на пустом кэше
	tail := l.popTail()
	if tail != nil {
		t.Error("popTail должен вернуть nil для пустого кэша")
	}
}

func TestPutWithZeroCapacity(t *testing.T) {
	l := New(0)
	// После валидации capacity станет 1, но проверим поведение
	l.Put("1", "1")
	if !l.Has("1") {
		t.Error("Ключ должен быть добавлен (capacity стал 1)")
	}
}

func TestUpdateExistingKey(t *testing.T) {
	l := New(2)
	l.Put("1", "value1")
	l.Put("2", "value2")

	// Обновляем существующий ключ
	l.Put("1", "updated_value")

	if val, err := l.Get("1"); err != nil || val != "updated_value" {
		t.Error("Значение должно быть обновлено, получено:", val)
	}

	// Проверяем, что размер кэша не изменился
	if len(l.cache) != 2 {
		t.Error("Размер кэша должен остаться 2, получен:", len(l.cache))
	}
}
