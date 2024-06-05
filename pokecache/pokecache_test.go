package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = time.Second * 10
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://exmaple.com",
			value: []byte("hello there"),
		},
		{
			key:   "https://example.com/test",
			value: []byte("test example"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Runing test %d", i), func(t *testing.T) {
			cache := NewCache(interval)

			cache.Add(c.key, c.value)

			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key")
				return
			}

			if string(val) != string(c.value) {
				t.Errorf("Expected to find value")
				return
			}

		})
	}
}

func TestReepLoop(t *testing.T) {
	const baseTime = 2 * time.Second
	const waitTime = baseTime + 2*time.Second
	const url = "https://test.com"
	cache := NewCache(baseTime)

	cache.Add(url, []byte("test data"))

	if _, ok := cache.Get(url); !ok {
		t.Errorf("exptect to find the key")
		return
	}

	time.Sleep(waitTime)

	if _, ok := cache.Get(url); ok {
		t.Errorf("expted to not find the key")
	}

}
