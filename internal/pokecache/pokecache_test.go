package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "http://www.example.com/",
			val: []byte("testdata"),
		},
		{
			key: "http://www.example.com/path",
			val: []byte("moredatatotest"),
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(tc.key, tc.val)

			got, ok := cache.Get(tc.key)
			if !ok {
				t.Errorf("key: %s not found in cache", tc.key)
			}
			if string(got) != string(tc.val) {
				t.Errorf("expect value: %#v doesn't match actual value %#v",
					string(got), string(tc.val))
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	url := "https:://example.com"
	td := []byte("testdata")
	cache.Add(url, td)

	_, ok := cache.Get(url)
	if !ok {
		t.Errorf("expected key: %v to be presented after being added", url)
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(url)
	if ok {
		t.Errorf("expected key: %v to be removed after lapsed time", url)
		return
	}
}
