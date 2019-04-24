package main

import "testing"

// simple cache exercise
func TestGetPut(t *testing.T) {
	tests := []struct {
		k, v string
	}{
		{"foo", "bar"},
		{"boo", "baz"},
	}
	c := &cache{
		store: make(map[string]string),
	}
	for _, test := range tests {
		c.Put(test.k, test.v)
		v, err := c.Get(test.k)
		if err != nil {
			t.Errorf("Wanted %s got %s", test.v, err)
		}
		if test.v != v {
			t.Errorf("Wanted %s got %s", test.v, v)
		}
	}
}

func TestCacheMiss(t *testing.T) {
	c := &cache{
		store: make(map[string]string),
	}
	v, err := c.Get("NotInTheCache")
	if err != errCacheMiss {
		t.Errorf("Expected errCacheMiss, got %s", v)
	}
}
