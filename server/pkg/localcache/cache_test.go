package localcache

import (
	"fmt"
	"testing"
)

func TestLocalCache(t *testing.T) {
	c := &Cache{
		Vals: make(map[string][]byte),
	}

	_ = c.Set("key1", "value1")

	var v string
	if err := c.Get("key1", &v); err != nil {
		fmt.Println(err)
	}
	fmt.Println("get key1 value", v)
}
