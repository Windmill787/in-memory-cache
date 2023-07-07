package inmemorycache

import "testing"

func TestSet(t *testing.T) {
	cache := NewCache()
	cache.Set("test", 1)
}

func TestGet(t *testing.T) {
	cache := NewCache()
	cache.Get("test")
}

func TestDelete(t *testing.T) {
	cache := NewCache()
	cache.Delete("test")
}

func TestSetAndGet(t *testing.T) {
	cache := NewCache()
	cache.Set("test", 1)
	cache.Get("test")
}

func TestSetAndDelete(t *testing.T) {
	cache := NewCache()
	cache.Set("test", 1)
	cache.Delete("test")
}

func TestSetGetDelete(t *testing.T) {
	cache := NewCache()
	cache.Set("test", 1)
	cache.Get("test")
	cache.Delete("test")
}
