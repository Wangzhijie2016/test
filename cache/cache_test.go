package cache

import (
	"testing"
)

func TestCache(t *testing.T) {
	if Cache() != "v2.0.7: Cache()" {
		t.Error("Cache() error")
	}
}
