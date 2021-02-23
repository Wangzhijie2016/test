package cache

import (
	"testing"
)

func TestCache(t *testing.T) {
	if Cache() != "Cache()" {
		t.Error("Cache() error")
	}
}
